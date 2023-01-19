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
	"time"

	gondpi "github.com/fs714/go-ndpi"
	"github.com/fs714/go-ndpi/types"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type FlowFingerprint struct {
	SrcAddr  string
	DstAddr  string
	SrcPort  string
	DstPort  string
	Protocol string
}

type FlowInfo struct {
	TotalL2Packets         int64
	TotalL3PayloadBytes    int64
	StartTs                time.Time
	EndTs                  time.Time
	NdpiProcessedPackets   int64
	NdpiProcessedBytes     int64
	NdpiDetectionCompleted bool
	NdpiFlow               *gondpi.NdpiFlowHandle
	Mu                     sync.Mutex
}

func NewFlowInfo() (*FlowInfo, error) {
	flowInfo := FlowInfo{}

	ndpiFlow, err := gondpi.GetNdpiFlowHandle()
	if err != nil {
		return nil, err
	}

	flowInfo.NdpiFlow = ndpiFlow

	return &flowInfo, nil
}

type FlowStats struct {
	FlowMap map[FlowFingerprint]*FlowInfo
	Mu      sync.Mutex
}

func (fs *FlowStats) AddFlow(fp FlowFingerprint, flowInfo *FlowInfo) {
	fs.Mu.Lock()
	defer fs.Mu.Unlock()

	fs.FlowMap[fp] = flowInfo
}

func (fs *FlowStats) GetFlow(fp FlowFingerprint) *FlowInfo {
	fs.Mu.Lock()
	defer fs.Mu.Unlock()

	flowInfo, ok := fs.FlowMap[fp]
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

	flowInfo, ok = fs.FlowMap[fpr]
	if ok {
		return flowInfo
	}

	return nil
}

func main() {
	var ifaceName string
	flag.StringVar(&ifaceName, "i", "", "interface name")
	flag.Parse()

	Stats := FlowStats{
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
	ndpiHandle, err := gondpi.NdpiHandleInitialize(detectionBitmask)
	if err != nil {
		fmt.Printf("failed to initialize NdpiHandle with err: %s\n", err.Error())
		return
	}

	defer gondpi.NdpiHandleExit(ndpiHandle)

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
	var ts int
	var ipData []byte
	var ipLength uint16
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
					SrcAddr:  ip4.SrcIP.String(),
					DstAddr:  ip4.DstIP.String(),
					SrcPort:  tcp.SrcPort.String(),
					DstPort:  tcp.DstPort.String(),
					Protocol: "tcp",
				}

				flowInfo = Stats.GetFlow(fp)
				if flowInfo == nil {
					flowInfo, err = NewFlowInfo()
					if err != nil {
						fmt.Printf("failed to new FlowInfo with err: %s\n", err.Error())
						continue
					}

					flowInfo.Mu.Lock()
					flowInfo.StartTs = time.Now()
					flowInfo.Mu.Unlock()

					Stats.AddFlow(fp, flowInfo)
				}

				flowInfo.Mu.Lock()
				flowInfo.TotalL2Packets++
				flowInfo.TotalL3PayloadBytes += int64(ipLength)
				flowInfo.EndTs = time.Now()

				if !flowInfo.NdpiDetectionCompleted && len(eth.Payload) > 0 {
					flowInfo.NdpiProcessedPackets++
					flowInfo.NdpiProcessedBytes += int64(ipLength)

					ts = ci.Timestamp.Second()*1000 + ci.Timestamp.Nanosecond()/1000000
					ipData = eth.Payload

					// ipData := make([]byte, len(eth.Payload))
					// copy(ipData, eth.Payload)

					ndpiProto = gondpi.NdpiPacketProcessing(ndpiHandle, flowInfo.NdpiFlow, ipData, ipLength, ts)
					if ndpiProto.MasterProtocolId != types.NDPI_PROTOCOL_UNKNOWN || ndpiProto.AppProtocolId != types.NDPI_PROTOCOL_UNKNOWN {
						flowInfo.NdpiDetectionCompleted = true
						gondpi.FreeNdpiFlowHandle(flowInfo.NdpiFlow)

						masterProto := types.NdpiProtocolIdMap[ndpiProto.MasterProtocolId]
						appProt := types.NdpiProtocolIdMap[ndpiProto.AppProtocolId]
						category := types.NdpiCategoryIdMap[ndpiProto.CategoryId]

						detectedProtocolStack := flowInfo.NdpiFlow.GetDetectedProtocolStack()
						processedPktNum := flowInfo.NdpiFlow.GetProcessedPktNum()
						flowExtraInfo := flowInfo.NdpiFlow.GetFlowExtraInfo()
						hostServerName := flowInfo.NdpiFlow.GetHostServerName()
						http := flowInfo.NdpiFlow.GetHttp()
						protocolCategory := flowInfo.NdpiFlow.GetProtocolCategory()

						fmt.Println("------")
						fmt.Printf("Master Protocol: %s, App Protocol: %s, Category: %s\n", masterProto, appProt, category)

						fmt.Printf("Detected Protocol Stack: %s, %s\n", types.NdpiProtocolIdMap[detectedProtocolStack[0]], types.NdpiProtocolIdMap[detectedProtocolStack[1]])
						fmt.Printf("Processed Pkt Num: %d\n", processedPktNum)
						fmt.Printf("Flow Extra Info: %s\n", flowExtraInfo)
						fmt.Printf("Host Server Name: %s\n", hostServerName)
						httpJson, _ := json.MarshalIndent(http, "", "  ")
						fmt.Printf("http struct:\n%s\n", string(httpJson))
						fmt.Printf("Http Method: %s, Http Request Version: %s\n", types.NdpiHttpMethodIdMap[http.NdpiHttpMethod], types.NdpiHttpRequestVersionIdMap[http.RequestVersion])
						fmt.Printf("Protocol Category: %s\n", types.NdpiCategoryIdMap[protocolCategory])

						fpJson, _ := json.MarshalIndent(fp, "", "  ")
						flowInfoJson, _ := json.MarshalIndent(flowInfo, "", "  ")
						fmt.Println(string(fpJson))
						fmt.Println(string(flowInfoJson))
					}
				}

				flowInfo.Mu.Unlock()
			}
		}
	}
}
