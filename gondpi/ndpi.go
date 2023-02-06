// Golang wrapper of ndpi 4.4.0
//
// Debian packages needed:
// apt-get install libpcap-dev libmaxminddb-dev
package gondpi

/*
#cgo CFLAGS: -I/usr/include/
#cgo LDFLAGS: -Wl,-Bstatic -lndpi -lmaxminddb -Wl,-Bdynamic -lpcap -lm -pthread
#include "ndpi_linux.h"
*/
import "C"

import (
	"encoding/binary"
	"reflect"
	"sync"
	"unsafe"

	"github.com/fs714/go-ndpi/gondpi/types"
	"github.com/pkg/errors"
)

const (
	NdpiBitmaskSize = 16
)

type NdpiDetectionModuleStructPtr *C.struct_ndpi_detection_module_struct
type NdpiFlowStructPtr *C.struct_ndpi_flow_struct

func NdpiCategoryToId(category C.ndpi_protocol_category_t) types.NdpiCategory {
	return types.NdpiCategory(category)
}

func NewNdpiProtocolBitmask() []uint32 {
	return make([]uint32, NdpiBitmaskSize)
}

func NdpiProtocolBitmaskAdd(bitmask []uint32, proto types.NdpiProtocol) []uint32 {
	ndpiBitmask := &C.NDPI_PROTOCOL_BITMASK{}
	ndpiBitmask.fds_bits = *(*[NdpiBitmaskSize]C.uint32_t)(unsafe.Pointer(&bitmask[0]))

	C.ndpi_protocol_bitmask_add(ndpiBitmask, C.uint16_t(proto))

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bitmask))
	sliceHeader.Len = NdpiBitmaskSize
	sliceHeader.Cap = NdpiBitmaskSize
	sliceHeader.Data = uintptr(unsafe.Pointer(&ndpiBitmask.fds_bits[0]))

	return bitmask
}

func NdpiProtocolBitmaskDel(bitmask []uint32, proto types.NdpiProtocol) []uint32 {
	ndpiBitmask := &C.NDPI_PROTOCOL_BITMASK{}
	ndpiBitmask.fds_bits = *(*[NdpiBitmaskSize]C.uint32_t)(unsafe.Pointer(&bitmask[0]))

	C.ndpi_protocol_bitmask_del(ndpiBitmask, C.uint16_t(proto))

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bitmask))
	sliceHeader.Len = NdpiBitmaskSize
	sliceHeader.Cap = NdpiBitmaskSize
	sliceHeader.Data = uintptr(unsafe.Pointer(&ndpiBitmask.fds_bits[0]))

	return bitmask
}

func NdpiProtocolBitmaskIsSet(bitmask []uint32, proto types.NdpiProtocol) bool {
	ndpiBitmask := &C.NDPI_PROTOCOL_BITMASK{}
	ndpiBitmask.fds_bits = *(*[NdpiBitmaskSize]C.uint32_t)(unsafe.Pointer(&bitmask[0]))

	ndpiIsSet := C.ndpi_protocol_bitmask_is_set(ndpiBitmask, C.uint16_t(proto))

	return bool(ndpiIsSet)
}

func NdpiProtocolBitmaskReset(bitmask []uint32) []uint32 {
	ndpiBitmask := &C.NDPI_PROTOCOL_BITMASK{}
	ndpiBitmask.fds_bits = *(*[NdpiBitmaskSize]C.uint32_t)(unsafe.Pointer(&bitmask[0]))

	C.ndpi_protocol_bitmask_reset(ndpiBitmask)

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bitmask))
	sliceHeader.Len = NdpiBitmaskSize
	sliceHeader.Cap = NdpiBitmaskSize
	sliceHeader.Data = uintptr(unsafe.Pointer(&ndpiBitmask.fds_bits[0]))

	return bitmask
}

func NdpiProtocolBitmaskSetAll(bitmask []uint32) []uint32 {
	ndpiBitmask := &C.NDPI_PROTOCOL_BITMASK{}
	ndpiBitmask.fds_bits = *(*[NdpiBitmaskSize]C.uint32_t)(unsafe.Pointer(&bitmask[0]))

	C.ndpi_protocol_bitmask_set_all(ndpiBitmask)

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bitmask))
	sliceHeader.Len = NdpiBitmaskSize
	sliceHeader.Cap = NdpiBitmaskSize
	sliceHeader.Data = uintptr(unsafe.Pointer(&ndpiBitmask.fds_bits[0]))

	return bitmask
}

type NdpiFlow struct {
	NdpiFlowPtr NdpiFlowStructPtr
	Mu          sync.Mutex
}

func NewNdpiFlow() (*NdpiFlow, error) {
	ndpiFlow := C.ndpi_flow_struct_malloc()
	if ndpiFlow == nil {
		C.ndpi_flow_struct_free(ndpiFlow)
		err := errors.New("null ndpi flow struct")
		return nil, err
	}

	f := &NdpiFlow{}
	f.NdpiFlowPtr = ndpiFlow

	return f, nil
}

func FreeNdpiFlow(f *NdpiFlow) {
	C.ndpi_flow_struct_free(f.NdpiFlowPtr)
}

func (f *NdpiFlow) GetDetectedProtocolStack() [2]types.NdpiProtocol {
	protoStack := [2]types.NdpiProtocol{}
	protoStack[0] = types.NdpiProtocol(f.NdpiFlowPtr.detected_protocol_stack[0])
	protoStack[1] = types.NdpiProtocol(f.NdpiFlowPtr.detected_protocol_stack[1])

	return protoStack
}

func (f *NdpiFlow) GetGuessedProtocolId() types.NdpiProtocol {
	return types.NdpiProtocol(f.NdpiFlowPtr.guessed_protocol_id)
}

func (f *NdpiFlow) GetGuessedHostProtocolId() types.NdpiProtocol {
	return types.NdpiProtocol(f.NdpiFlowPtr.guessed_host_protocol_id)
}

func (f *NdpiFlow) GetGuessedCategoryId() types.NdpiCategory {
	return types.NdpiCategory(f.NdpiFlowPtr.guessed_category)
}

func (f *NdpiFlow) GetGuessedHeaderCategoryId() types.NdpiCategory {
	return types.NdpiCategory(f.NdpiFlowPtr.guessed_header_category)
}

func (f *NdpiFlow) GetL4Protocol() types.IPProto {
	return types.IPProto(f.NdpiFlowPtr.l4_proto)
}

func (f *NdpiFlow) IsProtoIdAlreadyGuessed() bool {
	ret := C.ndpi_flow_get_protocol_id_already_guessed(f.NdpiFlowPtr)

	return (uint8(ret) != 0)
}

func (f *NdpiFlow) IsHostAlreadyGuessed() bool {
	ret := C.ndpi_flow_get_host_already_guessed(f.NdpiFlowPtr)

	return (uint8(ret) != 0)
}

func (f *NdpiFlow) IsFailWithUnknow() bool {
	ret := C.ndpi_flow_get_fail_with_unknown(f.NdpiFlowPtr)

	return (uint8(ret) != 0)
}

func (f *NdpiFlow) IsInitFinished() bool {
	ret := C.ndpi_flow_get_init_finished(f.NdpiFlowPtr)

	return (uint8(ret) != 0)
}

func (f *NdpiFlow) IsSetupPacketDirection() bool {
	ret := C.ndpi_flow_get_setup_packet_direction(f.NdpiFlowPtr)

	return (uint8(ret) != 0)
}

func (f *NdpiFlow) IsPacketDirection() bool {
	ret := C.ndpi_flow_get_packet_direction(f.NdpiFlowPtr)

	return (uint8(ret) != 0)
}

func (f *NdpiFlow) IsCheckExtraPackets() bool {
	ret := C.ndpi_flow_get_check_extra_packets(f.NdpiFlowPtr)

	return (uint8(ret) != 0)
}

func (f *NdpiFlow) IsIpv6() bool {
	ret := C.ndpi_flow_get_is_ipv6(f.NdpiFlowPtr)

	return (uint8(ret) != 0)
}

func (f *NdpiFlow) GetConfidence() types.NdpiConfidence {
	return types.NdpiConfidence(f.NdpiFlowPtr.confidence)
}

func (f *NdpiFlow) GetNextTcpSeqNR() [2]uint16 {
	nextTcpSeqNR := [2]uint16{}
	nextTcpSeqNR[0] = uint16(f.NdpiFlowPtr.next_tcp_seq_nr[0])
	nextTcpSeqNR[1] = uint16(f.NdpiFlowPtr.next_tcp_seq_nr[1])

	return nextTcpSeqNR
}

func (f *NdpiFlow) GetSrcAddr() uint32 {
	srcAddrBigEndian := make([]byte, 4)
	binary.BigEndian.PutUint32(srcAddrBigEndian, uint32(f.NdpiFlowPtr.saddr))
	srcAddr := binary.LittleEndian.Uint32(srcAddrBigEndian)

	return srcAddr
}

func (f *NdpiFlow) GetDstAddr() uint32 {
	dstAddrBigEndian := make([]byte, 4)
	binary.BigEndian.PutUint32(dstAddrBigEndian, uint32(f.NdpiFlowPtr.daddr))
	dstAddr := binary.LittleEndian.Uint32(dstAddrBigEndian)

	return dstAddr
}

func (f *NdpiFlow) GetSrcPort() uint16 {
	srcPortBigEndian := make([]byte, 2)
	binary.BigEndian.PutUint16(srcPortBigEndian, uint16(f.NdpiFlowPtr.sport))
	srcPort := binary.LittleEndian.Uint16(srcPortBigEndian)

	return srcPort
}

func (f *NdpiFlow) GetDstPort() uint16 {
	dstPortBigEndian := make([]byte, 2)
	binary.BigEndian.PutUint16(dstPortBigEndian, uint16(f.NdpiFlowPtr.dport))
	dstPort := binary.LittleEndian.Uint16(dstPortBigEndian)

	return dstPort
}

func (f *NdpiFlow) GetMaxExtraPacketsToCheck() uint8 {
	return uint8(f.NdpiFlowPtr.max_extra_packets_to_check)
}

func (f *NdpiFlow) GetNumExtraPacketsChecked() uint8 {
	return uint8(f.NdpiFlowPtr.num_extra_packets_checked)
}

func (f *NdpiFlow) GetProcessedPktNum() uint16 {
	return uint16(f.NdpiFlowPtr.num_processed_pkts)
}

func (f *NdpiFlow) GetLastPacketTimeMS() uint64 {
	return uint64(f.NdpiFlowPtr.last_packet_time_ms)
}

func (f *NdpiFlow) GetEntropy() float32 {
	return float32(f.NdpiFlowPtr.entropy)
}

func (f *NdpiFlow) GetFlowExtraInfo() string {
	return C.GoString(&f.NdpiFlowPtr.flow_extra_info[0])
}

func (f *NdpiFlow) GetHostServerName() string {
	return C.GoString(&f.NdpiFlowPtr.host_server_name[0])
}

// TODO
// u_int8_t initial_binary_bytes[8], initial_binary_bytes_len;
// u_int8_t risk_checked:1, ip_risk_mask_evaluated:1, host_risk_mask_evaluated:1, tree_risk_checked:1, _notused:4;
// ndpi_risk risk_mask; /* Stores the flow risk mask for flow peers */
// ndpi_risk risk; /* Issues found with this flow [bitmask of ndpi_risk] */
// struct ndpi_risk_information risk_infos[MAX_NUM_RISK_INFOS]; /* String that contains information about the risks found */
// u_int8_t num_risk_infos;

func (f *NdpiFlow) GetHttp() types.NdpiFlowHttp {
	return types.NdpiFlowHttp{
		NdpiHttpMethod:      types.HttpMethod(f.NdpiFlowPtr.http.method),
		RequestVersion:      types.HttpRequestVersion(f.NdpiFlowPtr.http.request_version),
		ResponseStatusCode:  uint16(f.NdpiFlowPtr.http.response_status_code),
		Url:                 C.GoString(f.NdpiFlowPtr.http.url),
		RequestContentType:  C.GoString(f.NdpiFlowPtr.http.request_content_type),
		ResponseContentType: C.GoString(f.NdpiFlowPtr.http.content_type),
		UserAgent:           C.GoString(f.NdpiFlowPtr.http.user_agent),
		DetectedOs:          C.GoString(f.NdpiFlowPtr.http.detected_os),
		XForwardedAddr:      C.GoString(f.NdpiFlowPtr.http.nat_ip),
	}
}

func (f *NdpiFlow) GetKerberosBuf() types.NdpiFlowKerberosBuf {
	return types.NdpiFlowKerberosBuf{
		PktBuf:        C.GoBytes(unsafe.Pointer(f.NdpiFlowPtr.kerberos_buf.pktbuf), C.int(f.NdpiFlowPtr.kerberos_buf.pktbuf_maxlen)),
		PktBufMaxLen:  uint16(f.NdpiFlowPtr.kerberos_buf.pktbuf_maxlen),
		PktBufCurrLen: uint16(f.NdpiFlowPtr.kerberos_buf.pktbuf_currlen),
	}
}

func (f *NdpiFlow) GetStun() types.NdpiFlowStun {
	return types.NdpiFlowStun{
		NumUdpPkts:         uint8(f.NdpiFlowPtr.stun.num_udp_pkts),
		NumBindingRequests: uint8(f.NdpiFlowPtr.stun.num_binding_requests),
		NumProcessedPkts:   uint16(f.NdpiFlowPtr.stun.num_processed_pkts),
	}
}

// TODO
// union {
// ......
// }

func (f *NdpiFlow) GetProtocolCategory() types.NdpiCategory {
	return NdpiCategoryToId(f.NdpiFlowPtr.category)
}

func (f *NdpiFlow) ToNdpiFlowInfo() *types.NdpiFlowInfo {
	return &types.NdpiFlowInfo{
		DetectedProtocolStack:   f.GetDetectedProtocolStack(),
		GuessedProtocolId:       f.GetGuessedProtocolId(),
		GuessedHostProtocolId:   f.GetGuessedHostProtocolId(),
		GuessedCategoryId:       f.GetGuessedCategoryId(),
		GuessedHeaderCategoryId: f.GetGuessedHeaderCategoryId(),
		L4Protocol:              f.GetL4Protocol(),
		IsProtoIdAlreadyGuessed: f.IsProtoIdAlreadyGuessed(),
		IsHostAlreadyGuessed:    f.IsHostAlreadyGuessed(),
		IsFailWithUnknow:        f.IsFailWithUnknow(),
		IsInitFinished:          f.IsInitFinished(),
		IsSetupPacketDirection:  f.IsSetupPacketDirection(),
		IsPacketDirection:       f.IsPacketDirection(),
		IsCheckExtraPackets:     f.IsCheckExtraPackets(),
		IsIpv6:                  f.IsIpv6(),
		Confidence:              f.GetConfidence(),
		NextTcpSeqNR:            f.GetNextTcpSeqNR(),
		SrcAddr:                 f.GetSrcAddr(),
		DstAddr:                 f.GetDstAddr(),
		SrcPort:                 f.GetSrcPort(),
		DstPort:                 f.GetDstPort(),
		MaxExtraPacketsToCheck:  f.GetMaxExtraPacketsToCheck(),
		NumExtraPacketsChecked:  f.GetNumExtraPacketsChecked(),
		ProcessedPktNum:         f.GetProcessedPktNum(),
		LastPacketTimeMS:        f.GetLastPacketTimeMS(),
		Entropy:                 f.GetEntropy(),
		FlowExtraInfo:           f.GetFlowExtraInfo(),
		HostServerName:          f.GetHostServerName(),
		Http:                    f.GetHttp(),
		KerberosBuf:             f.GetKerberosBuf(),
		Stun:                    f.GetStun(),
		ProtocolCategory:        f.GetProtocolCategory(),
	}
}

type NdpiProto struct {
	MasterProtocolId types.NdpiProtocol
	AppProtocolId    types.NdpiProtocol
	CategoryId       types.NdpiCategory
}

type NdpiProtoDefaults struct {
	ProtoName        string
	ProtoCategory    types.NdpiCategory
	IsClearTextProto bool
	IsAppProtocol    bool
	SubProtocols     []types.NdpiProtocol
	ProtoId          types.NdpiProtocol
	ProtoIdx         uint16
	TcpDefaultPorts  []uint16
	UdpDefaultPorts  []uint16
	ProtoBreed       types.NdpiProtocolBreed
}

type NdpiDetectionModule struct {
	NdpiPtr NdpiDetectionModuleStructPtr
	Mu      sync.Mutex
}

func NdpiDetectionModuleInitialize(prefs uint32, detectionBitmask []uint32) (*NdpiDetectionModule, error) {
	ndpiBitmask := &C.NDPI_PROTOCOL_BITMASK{}
	ndpiBitmask.fds_bits = *(*[NdpiBitmaskSize]C.uint32_t)(unsafe.Pointer(&detectionBitmask[0]))
	ndpi := C.ndpi_detection_module_initialize(C.ndpi_init_prefs(prefs), ndpiBitmask)
	if ndpi == nil {
		C.ndpi_detection_module_exit(ndpi)
		err := errors.New("null ndpi detection module struct")
		return nil, err
	}

	dm := &NdpiDetectionModule{}
	dm.NdpiPtr = ndpi

	return dm, nil
}

func NdpiDetectionModuleExit(dm *NdpiDetectionModule) {
	C.ndpi_detection_module_exit(dm.NdpiPtr)
}

func (dm *NdpiDetectionModule) SetLogLevel(level uint16) {
	C.ndpi_set_log_level(dm.NdpiPtr, C.uint(level))
}

func (dm *NdpiDetectionModule) SetDebugBitmask(debugBitmask []uint32) {
	ndpiBitmask := C.NDPI_PROTOCOL_BITMASK{}
	ndpiBitmask.fds_bits = *(*[NdpiBitmaskSize]C.uint32_t)(unsafe.Pointer(&debugBitmask[0]))

	C.ndpi_set_debug_bitmask(dm.NdpiPtr, ndpiBitmask)
}

func (dm *NdpiDetectionModule) LoadProtocolsFile(path string) error {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	ret := C.ndpi_load_protocols_file(dm.NdpiPtr, cPath)
	if int(ret) != 0 {
		err := errors.New("failed to load protocols file")
		return err
	}

	return nil
}

func (dm *NdpiDetectionModule) LoadCategoriesFile(path string, userData string) uint32 {
	cPath := C.CString(path)
	cUserData := C.CString(userData)
	defer C.free(unsafe.Pointer(cPath))

	num := C.ndpi_load_categories_file(dm.NdpiPtr, cPath, unsafe.Pointer(cUserData))

	return uint32(num)
}

func (dm *NdpiDetectionModule) LoadRiskDomainFile(path string) uint32 {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	num := C.ndpi_load_risk_domain_file(dm.NdpiPtr, cPath)

	return uint32(num)
}

func (dm *NdpiDetectionModule) LoadMaliciousJa3File(path string) uint32 {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	num := C.ndpi_load_malicious_ja3_file(dm.NdpiPtr, cPath)

	return uint32(num)
}

func (dm *NdpiDetectionModule) LoadMaliciousSha1File(path string) uint32 {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	num := C.ndpi_load_malicious_sha1_file(dm.NdpiPtr, cPath)

	return uint32(num)
}

func (dm *NdpiDetectionModule) GetProtoDefaults() []NdpiProtoDefaults {
	isClearTextProtoList := make([]bool, C.NDPI_MAX_SUPPORTED_PROTOCOLS+C.NDPI_MAX_NUM_CUSTOM_PROTOCOLS)
	isAppProtocolList := make([]bool, C.NDPI_MAX_SUPPORTED_PROTOCOLS+C.NDPI_MAX_NUM_CUSTOM_PROTOCOLS)

	protoDefaults := C.ndpi_proto_defaults_get(dm.NdpiPtr, (*C.bool)(unsafe.Pointer(&isClearTextProtoList[0])),
		(*C.bool)(unsafe.Pointer(&isAppProtocolList[0])))
	pds := make([]C.ndpi_proto_defaults_t, 0)
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&pds)))
	sliceHeader.Cap = C.NDPI_MAX_SUPPORTED_PROTOCOLS + C.NDPI_MAX_NUM_CUSTOM_PROTOCOLS
	sliceHeader.Len = C.NDPI_MAX_SUPPORTED_PROTOCOLS + C.NDPI_MAX_NUM_CUSTOM_PROTOCOLS
	sliceHeader.Data = uintptr(unsafe.Pointer(protoDefaults))

	npds := make([]NdpiProtoDefaults, 0)
	for i := 0; i < C.NDPI_MAX_SUPPORTED_PROTOCOLS+C.NDPI_MAX_NUM_CUSTOM_PROTOCOLS; i++ {
		if C.GoString(pds[i].protoName) == "" {
			continue
		}

		subProtocols := make([]types.NdpiProtocol, 0)
		if pds[i].subprotocol_count > 0 {
			sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&subProtocols)))
			sliceHeader.Cap = int(pds[i].subprotocol_count)
			sliceHeader.Len = int(pds[i].subprotocol_count)
			sliceHeader.Data = uintptr(unsafe.Pointer(pds[i].subprotocols))
		}

		tcpDefaultPorts := make([]uint16, 0)
		sliceHeader = (*reflect.SliceHeader)((unsafe.Pointer(&tcpDefaultPorts)))
		sliceHeader.Cap = C.MAX_DEFAULT_PORTS
		sliceHeader.Len = C.MAX_DEFAULT_PORTS
		sliceHeader.Data = uintptr(unsafe.Pointer(&pds[i].tcp_default_ports[0]))

		udpDefaultPorts := make([]uint16, 0)
		sliceHeader = (*reflect.SliceHeader)((unsafe.Pointer(&udpDefaultPorts)))
		sliceHeader.Cap = C.MAX_DEFAULT_PORTS
		sliceHeader.Len = C.MAX_DEFAULT_PORTS
		sliceHeader.Data = uintptr(unsafe.Pointer(&pds[i].udp_default_ports[0]))

		npd := NdpiProtoDefaults{
			ProtoName:        C.GoString(pds[i].protoName),
			ProtoCategory:    types.NdpiCategory(pds[i].protoCategory),
			IsClearTextProto: isClearTextProtoList[i],
			IsAppProtocol:    isAppProtocolList[i],
			SubProtocols:     subProtocols,
			ProtoId:          types.NdpiProtocol(pds[i].protoId),
			ProtoIdx:         uint16(pds[i].protoIdx),
			TcpDefaultPorts:  tcpDefaultPorts,
			UdpDefaultPorts:  udpDefaultPorts,
			ProtoBreed:       types.NdpiProtocolBreed(pds[i].protoBreed),
		}

		npds = append(npds, npd)
	}

	return npds
}

func (dm *NdpiDetectionModule) PacketProcessing(flow *NdpiFlow, ipPacket []byte, ipPacketLen uint16, timestamp int64) NdpiProto {
	ipPktPtr := (*C.u_char)(unsafe.Pointer(&ipPacket[0]))
	ipPktLen := C.ushort(ipPacketLen)
	ipPktTs := C.uint64_t(timestamp)

	proto := C.ndpi_packet_processing(dm.NdpiPtr, flow.NdpiFlowPtr, ipPktPtr, ipPktLen, ipPktTs)
	// fmt.Printf("%v, %v, %v\n", proto.master_protocol, proto.app_protocol, proto.category)

	ndpiProto := NdpiProto{
		MasterProtocolId: types.NdpiProtocol(proto.master_protocol),
		AppProtocolId:    types.NdpiProtocol(proto.app_protocol),
		CategoryId:       NdpiCategoryToId(proto.category),
	}

	return ndpiProto
}

func (dm *NdpiDetectionModule) IsExtraDissectionPossible(flow *NdpiFlow) bool {
	ret := C.ndpi_extra_dissection_possible(dm.NdpiPtr, flow.NdpiFlowPtr)
	if uint32(ret) == 0 {
		return false
	} else {
		return true
	}
}

func (dm *NdpiDetectionModule) DetectionGiveup(flow *NdpiFlow, enableGuess bool) (NdpiProto, bool) {
	var cEnableGuess C.uint8_t = 0
	if enableGuess {
		cEnableGuess = 1
	}

	var cIsProtoGuessed *C.uint8_t
	*cIsProtoGuessed = 0

	proto := C.ndpi_detection_giveup(dm.NdpiPtr, flow.NdpiFlowPtr, cEnableGuess, cIsProtoGuessed)

	ndpiProto := NdpiProto{
		MasterProtocolId: types.NdpiProtocol(proto.master_protocol),
		AppProtocolId:    types.NdpiProtocol(proto.app_protocol),
		CategoryId:       NdpiCategoryToId(proto.category),
	}

	isProtoGuessed := false
	if enableGuess && uint8(*cIsProtoGuessed) == 1 {
		isProtoGuessed = true
	}

	return ndpiProto, isProtoGuessed
}
