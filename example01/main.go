/*
Example of simple uage:
- Packet Capture by libpcap
- Packet L2 - L4 inspection by gopacket
- Packet L7 inspection by nDPI
- Golang map for flow recording
*/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"sync"

	"github.com/fs714/go-ndpi/gondpi"
	"github.com/fs714/go-ndpi/gondpi/types"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

const (
	MaxPacketsUdpDissection = 24
	MaxPacketsTcpDissection = 80
)

type FlowFingerprint struct {
	SrcAddr  uint32
	DstAddr  uint32
	SrcPort  uint16
	DstPort  uint16
	Protocol uint8
}

func (f *FlowFingerprint) ToString() string {
	ff := struct {
		SrcAddr  string
		DstAddr  string
		SrcPort  uint16
		DstPort  uint16
		Protocol string
	}{
		SrcAddr:  types.IntToIPv4(f.SrcAddr).String(),
		DstAddr:  types.IntToIPv4(f.DstAddr).String(),
		SrcPort:  f.SrcPort,
		DstPort:  f.DstPort,
		Protocol: (*types.IPProto)(&f.Protocol).ToName(),
	}

	ffJson, _ := json.MarshalIndent(ff, "", "  ")

	return string(ffJson)
}

type FlowInfo struct {
	TotalL2Packets         int64
	TotalL3PayloadBytes    int64
	StartTsMilli           int64
	EndTsMilli             int64
	NdpiProcessedPackets   int64
	NdpiProcessedBytes     int64
	NdpiDetectionCompleted bool
	NdpiIsGuessed          bool
	NdpiFlow               *gondpi.NdpiFlow
	Mu                     sync.Mutex
}

func NewFlowInfo() (*FlowInfo, error) {
	flowInfo := FlowInfo{}

	ndpiFlow, err := gondpi.NewNdpiFlow()
	if err != nil {
		return nil, err
	}

	flowInfo.NdpiFlow = ndpiFlow

	return &flowInfo, nil
}

type IfaceStats struct {
	FlowMap map[FlowFingerprint]*FlowInfo
	Mu      sync.Mutex
}

func (s *IfaceStats) AddFlow(fp FlowFingerprint, flowInfo *FlowInfo) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	s.FlowMap[fp] = flowInfo
}

func (s *IfaceStats) DeleteFlow(fp FlowFingerprint) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	_, ok := s.FlowMap[fp]
	if !ok {
		return
	}

	delete(s.FlowMap, fp)
}

func (s *IfaceStats) GetFlow(fp FlowFingerprint) *FlowInfo {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	flowInfo, ok := s.FlowMap[fp]
	if ok {
		return flowInfo
	}

	fpr := FlowFingerprint{
		SrcAddr:  fp.DstAddr,
		DstAddr:  fp.SrcAddr,
		SrcPort:  fp.DstPort,
		DstPort:  fp.SrcPort,
		Protocol: fp.Protocol,
	}

	flowInfo, ok = s.FlowMap[fpr]
	if ok {
		return flowInfo
	}

	return nil
}

func main() {
	var ifaceName string
	var enableGuess bool
	flag.StringVar(&ifaceName, "i", "", "interface name")
	flag.BoolVar(&enableGuess, "g", true, "enable protocol guess or not")
	flag.Parse()

	Stats := IfaceStats{
		FlowMap: map[FlowFingerprint]*FlowInfo{},
	}

	handle, err := pcap.OpenLive(ifaceName, 65535, true, pcap.BlockForever)
	if err != nil {
		fmt.Printf("failed to open live interface %s with err: %s\n", ifaceName, err.Error())
		return
	}

	err = handle.SetBPFFilter("")
	if err != nil {
		fmt.Printf("failed to set BPF filter with err: %s\n", err.Error())
		return
	}

	defer handle.Close()

	detectionBitmask := gondpi.NewNdpiProtocolBitmask()
	detectionBitmask = gondpi.NdpiProtocolBitmaskSetAll(detectionBitmask)
	ndpiDM, err := gondpi.NdpiDetectionModuleInitialize(types.NDPI_NO_PREFS, detectionBitmask)
	if err != nil {
		fmt.Printf("failed to initialize ndpi detection module with err: %s\n", err.Error())
		return
	}

	defer ndpiDM.Close()

	eth := &layers.Ethernet{}
	ip4 := &layers.IPv4{}
	tcp := &layers.TCP{}
	udp := &layers.UDP{}
	payload := &gopacket.Payload{}
	parser := gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, eth, ip4, tcp, udp, payload)
	decoded := make([]gopacket.LayerType, 0, 4)

	var data []byte
	var ci gopacket.CaptureInfo
	var fp FlowFingerprint
	var flowInfo *FlowInfo
	var ndpiProto gondpi.NdpiProto
	var ts int64
	var ipData []byte
	var ipLength uint16
	var isGuessed bool
	for {
		data, ci, err = handle.ZeroCopyReadPacketData()
		if err != nil {
			fmt.Printf("error getting packet: %s\n", err.Error())
			continue
		}

		err = parser.DecodeLayers(data, &decoded)
		if err != nil {
			fmt.Printf("error decoding packet: %s\n", err.Error())
			continue
		}

		for _, typ := range decoded {
			switch typ {
			case layers.LayerTypeTCP:
				ipLength = ip4.Length
				fp = FlowFingerprint{
					SrcAddr:  types.IPv4ToInt(ip4.SrcIP),
					DstAddr:  types.IPv4ToInt(ip4.DstIP),
					SrcPort:  uint16(tcp.SrcPort),
					DstPort:  uint16(tcp.DstPort),
					Protocol: uint8(types.IPProtocolTCP),
				}

				flowInfo = Stats.GetFlow(fp)
				if flowInfo == nil {
					flowInfo, err = NewFlowInfo()
					if err != nil {
						fmt.Printf("failed to new FlowInfo with err: %s\n", err.Error())
						continue
					}

					flowInfo.Mu.Lock()
					flowInfo.StartTsMilli = ci.Timestamp.UnixMilli()
					flowInfo.Mu.Unlock()

					Stats.AddFlow(fp, flowInfo)
				}

				flowInfo.Mu.Lock()
				flowInfo.TotalL2Packets++
				flowInfo.TotalL3PayloadBytes += int64(ipLength)
				flowInfo.EndTsMilli = ci.Timestamp.UnixMilli()

				if !flowInfo.NdpiDetectionCompleted {
					flowInfo.NdpiProcessedPackets++
					flowInfo.NdpiProcessedBytes += int64(ipLength)

					// ipData := make([]byte, len(eth.Payload))
					// copy(ipData, eth.Payload)
					ipData = eth.Payload
					ts = ci.Timestamp.UnixMilli()
					ndpiProto = ndpiDM.PacketProcessing(flowInfo.NdpiFlow, ipData, ipLength, ts)

					enoughPackets := false
					if flowInfo.NdpiProcessedPackets > MaxPacketsTcpDissection {
						enoughPackets = true
					}

					extraDissectionPossible := ndpiDM.IsExtraDissectionPossible(flowInfo.NdpiFlow)

					if enoughPackets || ndpiProto.AppProtocolId != types.NDPI_PROTOCOL_UNKNOWN {
						if !enoughPackets && extraDissectionPossible {

						} else {
							if enableGuess && ndpiProto.AppProtocolId == types.NDPI_PROTOCOL_UNKNOWN {
								ndpiProto, isGuessed = ndpiDM.DetectionGiveup(flowInfo.NdpiFlow, enableGuess)
								flowInfo.NdpiIsGuessed = isGuessed
							}

							flowInfo.NdpiDetectionCompleted = true

							fmt.Println("------")
							masterProto := ndpiProto.MasterProtocolId.ToName()
							appProto := ndpiProto.AppProtocolId.ToName()
							category := ndpiProto.CategoryId.ToName()

							fmt.Printf("Master Protocol: %s, App Protocol: %s, Category: %s\n", masterProto, appProto, category)

							fpString := fp.ToString()
							flowInfoJson, _ := json.MarshalIndent(flowInfo, "", "  ")
							ndpiFlowInfoString := flowInfo.NdpiFlow.ToNdpiFlowInfo().ToString()
							fmt.Println(fpString)
							fmt.Println(string(flowInfoJson))
							fmt.Println(ndpiFlowInfoString)

							flowInfo.NdpiFlow.Close()
						}
					}
				}

				flowInfo.Mu.Unlock()
			}
		}
	}
}
