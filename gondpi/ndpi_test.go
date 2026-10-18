package gondpi_test

import (
	"bufio"
	"io"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/fs714/go-ndpi/gondpi"
	"github.com/fs714/go-ndpi/gondpi/types"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcapgo"
)

// Same limits used by ndpiReader before giving up on a flow
const (
	maxTcpDissectedPkts = 80
	maxUdpDissectedPkts = 24
)

// Pcaps and expected results (*.pcap.out) come from nDPI test suite (tests/cfgs/default).
// Set NDPI_TEST_PCAP_DIR to the "default" directory of a nDPI source tree to check all of them.
const testDataDir = "testdata/ndpi"

func newDetectionModule(t *testing.T) *gondpi.NdpiDetectionModule {
	t.Helper()

	bitmask := gondpi.NdpiProtocolBitmaskSetAll(gondpi.NewNdpiProtocolBitmask())
	dm, err := gondpi.NdpiDetectionModuleInitialize(types.NDPI_LICENSE_NOT_FOR_PROFIT_LGPL, bitmask)
	if err != nil {
		t.Fatalf("failed to initialize ndpi detection module: %s", err.Error())
	}
	t.Cleanup(dm.Close)

	return dm
}

func TestProtocolBitmask(t *testing.T) {
	bitmask := gondpi.NewNdpiProtocolBitmask()
	if len(bitmask) != gondpi.NdpiBitmaskSize {
		t.Fatalf("bitmask size is %d, expected %d", len(bitmask), gondpi.NdpiBitmaskSize)
	}

	bitmask = gondpi.NdpiProtocolBitmaskAdd(bitmask, types.NDPI_PROTOCOL_SSH)
	if !gondpi.NdpiProtocolBitmaskIsSet(bitmask, types.NDPI_PROTOCOL_SSH) {
		t.Errorf("SSH not set after add")
	}
	if gondpi.NdpiProtocolBitmaskIsSet(bitmask, types.NDPI_PROTOCOL_HTTP) {
		t.Errorf("HTTP set without add")
	}

	bitmask = gondpi.NdpiProtocolBitmaskDel(bitmask, types.NDPI_PROTOCOL_SSH)
	if gondpi.NdpiProtocolBitmaskIsSet(bitmask, types.NDPI_PROTOCOL_SSH) {
		t.Errorf("SSH set after del")
	}

	bitmask = gondpi.NdpiProtocolBitmaskSetAll(bitmask)
	for _, w := range bitmask {
		if w != 0xFFFFFFFF {
			t.Fatalf("not all bits set after set all: %v", bitmask)
		}
	}

	bitmask = gondpi.NdpiProtocolBitmaskReset(bitmask)
	for _, w := range bitmask {
		if w != 0 {
			t.Fatalf("bits set after reset: %v", bitmask)
		}
	}
}

func TestDetectionModuleConfig(t *testing.T) {
	bitmask := gondpi.NdpiProtocolBitmaskSetAll(gondpi.NewNdpiProtocolBitmask())
	bitmask = gondpi.NdpiProtocolBitmaskDel(bitmask, types.NDPI_PROTOCOL_SSH)
	dm, err := gondpi.NewNdpiDetectionModule(types.NDPI_LICENSE_NOT_FOR_PROFIT_LGPL, bitmask)
	if err != nil {
		t.Fatalf("failed to create ndpi detection module: %s", err.Error())
	}
	defer dm.Close()

	if err := dm.SetConfig("", "dpi.guess_on_giveup", "0"); err != nil {
		t.Errorf("failed to set valid config: %s", err.Error())
	}
	if err := dm.SetConfig("", "not.existing.param", "1"); err == nil {
		t.Errorf("no error on unknown config param")
	}
	if err := dm.SetLogLevel(types.NDPI_LOG_ERROR); err != nil {
		t.Errorf("failed to set log level: %s", err.Error())
	}

	if dm.IsFinalized() {
		t.Fatalf("module finalized before Finalize")
	}
	if err := dm.Finalize(); err != nil {
		t.Fatalf("failed to finalize: %s", err.Error())
	}
	if !dm.IsFinalized() {
		t.Fatalf("module not finalized after Finalize")
	}

	if err := dm.SetConfig("", "dpi.guess_on_giveup", "1"); err == nil {
		t.Errorf("no error on config set after Finalize")
	}

	checks := []struct {
		proto, param, expected string
	}{
		{"ssh", "enable", "0"},
		{"http", "enable", "1"},
		{"", "dpi.guess_on_giveup", "0"},
	}
	for _, c := range checks {
		value, err := dm.GetConfig(c.proto, c.param)
		if err != nil {
			t.Errorf("failed to get config [%s][%s]: %s", c.proto, c.param, err.Error())
			continue
		}
		if value != c.expected {
			t.Errorf("config [%s][%s] is %q, expected %q", c.proto, c.param, value, c.expected)
		}
	}
}

func TestGetProtoDefaults(t *testing.T) {
	dm := newDetectionModule(t)

	npds := dm.GetProtoDefaults()
	if len(npds) == 0 || uint32(len(npds)) > dm.GetNumProtocols() {
		t.Fatalf("got %d protocol defaults with %d protocols", len(npds), dm.GetNumProtocols())
	}

	for _, npd := range npds {
		if npd.ProtoId != types.NDPI_PROTOCOL_HTTP {
			continue
		}

		if npd.ProtoName != "HTTP" || npd.ProtoCategory != types.NDPI_PROTOCOL_CATEGORY_WEB {
			t.Errorf("unexpected HTTP defaults: %+v", npd)
		}
		if len(npd.TcpDefaultPorts) == 0 || npd.TcpDefaultPorts[0] != (gondpi.NdpiPortRange{Low: 80, High: 80}) {
			t.Errorf("unexpected HTTP tcp default ports: %+v", npd.TcpDefaultPorts)
		}

		return
	}

	t.Errorf("HTTP not found in protocol defaults")
}

// flowKey identifies a bidirectional flow, endpoints are sorted
type flowKey struct {
	l4    string
	addrA string
	addrB string
}

func newFlowKey(l4 string, a string, b string) flowKey {
	if a > b {
		a, b = b, a
	}

	return flowKey{l4: l4, addrA: a, addrB: b}
}

func (k flowKey) String() string {
	return k.l4 + " " + k.addrA + " <-> " + k.addrB
}

// expectedProto is the protocol as printed by ndpiReader, e.g. "91.283" (master.app) or "7" (app only)
type expectedProto struct {
	ids  string
	name string
}

// e.g. "	1	TCP 172.16.238.1:58395 <-> 172.16.238.168:22 [proto: 92/SSH][Stack: SSH]..."
var outFlowRegexp = regexp.MustCompile(`^\t\d+\t(TCP|UDP) (\S+) <-> (\S+) .*?\[proto: ([0-9.]+)/([^\]]+)\]`)

func parseExpectedFlows(t *testing.T, path string) map[flowKey]expectedProto {
	t.Helper()

	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("failed to open expected results: %s", err.Error())
	}
	defer file.Close()

	flows := make(map[flowKey]expectedProto)
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	for scanner.Scan() {
		m := outFlowRegexp.FindStringSubmatch(scanner.Text())
		if m == nil {
			continue
		}

		flows[newFlowKey(m[1], m[2], m[3])] = expectedProto{ids: m[4], name: m[5]}
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("failed to read expected results: %s", err.Error())
	}

	return flows
}

func formatProtoIds(proto gondpi.NdpiProto) string {
	if proto.MasterProtocolId == types.NDPI_PROTOCOL_UNKNOWN || proto.MasterProtocolId == proto.AppProtocolId {
		return strconv.Itoa(int(proto.AppProtocolId))
	}

	return strconv.Itoa(int(proto.MasterProtocolId)) + "." + strconv.Itoa(int(proto.AppProtocolId))
}

type testFlow struct {
	ndpiFlow  *gondpi.NdpiFlow
	packets   int
	completed bool
	proto     gondpi.NdpiProto
}

func openPcap(t *testing.T, path string) (func() ([]byte, gopacket.CaptureInfo, error), gopacket.Decoder) {
	t.Helper()

	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("failed to open pcap: %s", err.Error())
	}
	t.Cleanup(func() { file.Close() })

	// Some pcapng files have a .pcap extension: check the magic number
	magic := make([]byte, 4)
	if _, err := io.ReadFull(file, magic); err != nil {
		t.Fatalf("failed to read pcap magic: %s", err.Error())
	}
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		t.Fatal(err)
	}

	if string(magic) == "\x0a\x0d\x0d\x0a" {
		r, err := pcapgo.NewNgReader(file, pcapgo.DefaultNgReaderOptions)
		if err != nil {
			t.Fatalf("failed to read pcapng: %s", err.Error())
		}
		return r.ReadPacketData, r.LinkType()
	}

	r, err := pcapgo.NewReader(file)
	if err != nil {
		t.Fatalf("failed to read pcap: %s", err.Error())
	}
	return r.ReadPacketData, r.LinkType()
}

// classifyPcap feeds TCP/UDP packets to nDPI the same way ndpiReader does: packets of a flow are
// processed until classification is completed or too many packets are seen, then nDPI gives up
func classifyPcap(t *testing.T, dm *gondpi.NdpiDetectionModule, path string) map[flowKey]gondpi.NdpiProto {
	t.Helper()

	readPacket, decoder := openPcap(t, path)
	flows := make(map[flowKey]*testFlow)
	for {
		data, ci, err := readPacket()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("failed to read packet: %s", err.Error())
		}

		pkt := gopacket.NewPacket(data, decoder, gopacket.NoCopy)
		var ipPacket []byte
		var srcIP, dstIP net.IP
		switch ip := pkt.NetworkLayer().(type) {
		case *layers.IPv4:
			ipPacket = append(append([]byte{}, ip.Contents...), ip.Payload...)
			srcIP, dstIP = ip.SrcIP, ip.DstIP
		case *layers.IPv6:
			ipPacket = append(append([]byte{}, ip.Contents...), ip.Payload...)
			srcIP, dstIP = ip.SrcIP, ip.DstIP
		default:
			continue
		}

		var l4 string
		var srcPort, dstPort int
		maxPkts := 0
		switch tl := pkt.TransportLayer().(type) {
		case *layers.TCP:
			l4, srcPort, dstPort, maxPkts = "TCP", int(tl.SrcPort), int(tl.DstPort), maxTcpDissectedPkts
		case *layers.UDP:
			l4, srcPort, dstPort, maxPkts = "UDP", int(tl.SrcPort), int(tl.DstPort), maxUdpDissectedPkts
		default:
			continue
		}

		src := net.JoinHostPort(srcIP.String(), strconv.Itoa(srcPort))
		dst := net.JoinHostPort(dstIP.String(), strconv.Itoa(dstPort))
		key := newFlowKey(l4, src, dst)
		flow, ok := flows[key]
		if !ok {
			ndpiFlow, err := gondpi.NewNdpiFlow()
			if err != nil {
				t.Fatalf("failed to create ndpi flow: %s", err.Error())
			}
			flow = &testFlow{ndpiFlow: ndpiFlow}
			flows[key] = flow
		}
		if flow.completed {
			continue
		}

		enoughPackets := flow.packets >= maxPkts
		flow.packets++
		flow.proto = dm.PacketProcessing(flow.ndpiFlow, ipPacket, uint16(len(ipPacket)), ci.Timestamp.UnixMilli())
		if enoughPackets || (flow.proto.State == types.NDPI_STATE_CLASSIFIED && !dm.IsExtraDissectionPossible(flow.ndpiFlow)) {
			flow.completed = true
			if flow.proto.State != types.NDPI_STATE_CLASSIFIED {
				flow.proto, _ = dm.DetectionGiveup(flow.ndpiFlow)
			}
		}
	}

	results := make(map[flowKey]gondpi.NdpiProto)
	for key, flow := range flows {
		if !flow.completed {
			flow.proto, _ = dm.DetectionGiveup(flow.ndpiFlow)
		}
		results[key] = flow.proto
		flow.ndpiFlow.Close()
	}

	return results
}

func checkPcap(t *testing.T, pcapPath string, outPath string) {
	expected := parseExpectedFlows(t, outPath)
	if len(expected) == 0 {
		t.Skip("no TCP/UDP flows in expected results")
	}

	dm := newDetectionModule(t)
	results := classifyPcap(t, dm, pcapPath)

	keys := make([]flowKey, 0, len(expected))
	for key := range expected {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i].String() < keys[j].String() })

	for _, key := range keys {
		exp := expected[key]
		proto, ok := results[key]
		if !ok {
			t.Errorf("%s: flow not found, expected %s/%s", key, exp.ids, exp.name)
			continue
		}
		if ids := formatProtoIds(proto); ids != exp.ids {
			t.Errorf("%s: got %s/%s.%s, expected %s/%s", key, ids,
				proto.MasterProtocolId.ToName(), proto.AppProtocolId.ToName(), exp.ids, exp.name)
		}
	}

	if len(results) != len(expected) {
		t.Errorf("got %d TCP/UDP flows, expected %d", len(results), len(expected))
	}
}

func TestPcapClassification(t *testing.T) {
	pcaps, err := filepath.Glob(filepath.Join(testDataDir, "*.pcap*"))
	if err != nil {
		t.Fatal(err)
	}

	for _, pcapPath := range pcaps {
		if strings.HasSuffix(pcapPath, ".out") {
			continue
		}
		t.Run(filepath.Base(pcapPath), func(t *testing.T) {
			checkPcap(t, pcapPath, pcapPath+".out")
		})
	}
}

// TestNdpiTestSuite runs the classification check on the whole nDPI test suite, when available
func TestNdpiTestSuite(t *testing.T) {
	dir := os.Getenv("NDPI_TEST_PCAP_DIR")
	if dir == "" {
		t.Skip("NDPI_TEST_PCAP_DIR not set")
	}

	pcaps, err := filepath.Glob(filepath.Join(dir, "pcap", "*"))
	if err != nil {
		t.Fatal(err)
	}

	for _, pcapPath := range pcaps {
		outPath := filepath.Join(dir, "result", filepath.Base(pcapPath)+".out")
		if _, err := os.Stat(outPath); err != nil {
			continue
		}
		t.Run(filepath.Base(pcapPath), func(t *testing.T) {
			checkPcap(t, pcapPath, outPath)
		})
	}
}
