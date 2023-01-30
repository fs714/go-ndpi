package gondpi

/*
#cgo CFLAGS: -I/usr/include/
#cgo LDFLAGS: -Wl,-Bstatic -lndpi -lmaxminddb -Wl,-Bdynamic -lpcap -lm -pthread
#include "ndpi_linux.h"
*/
import "C"

import (
	"reflect"
	"sync"
	"unsafe"

	"github.com/fs714/go-ndpi/gondpi/types"
	"github.com/pkg/errors"
)

const (
	NdpiBitmaskSize = 16
)

type ndpiDetectionModuleStructPtr *C.struct_ndpi_detection_module_struct
type ndpiFlowStructPtr *C.struct_ndpi_flow_struct

func NewNdpiProtocolBitmask() []uint32 {
	return make([]uint32, NdpiBitmaskSize)
}

func NdpiProtocolBitmaskAdd(bitmask []uint32, proto uint16) []uint32 {
	ndpiBitmask := &C.NDPI_PROTOCOL_BITMASK{}
	ndpiBitmask.fds_bits = *(*[NdpiBitmaskSize]C.uint32_t)(unsafe.Pointer(&bitmask[0]))

	C.ndpi_protocol_bitmask_add(ndpiBitmask, C.uint16_t(proto))

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bitmask))
	sliceHeader.Len = NdpiBitmaskSize
	sliceHeader.Cap = NdpiBitmaskSize
	sliceHeader.Data = uintptr(unsafe.Pointer(&ndpiBitmask.fds_bits[0]))

	return bitmask
}

func NdpiProtocolBitmaskDel(bitmask []uint32, proto uint16) []uint32 {
	ndpiBitmask := &C.NDPI_PROTOCOL_BITMASK{}
	ndpiBitmask.fds_bits = *(*[NdpiBitmaskSize]C.uint32_t)(unsafe.Pointer(&bitmask[0]))

	C.ndpi_protocol_bitmask_del(ndpiBitmask, C.uint16_t(proto))

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bitmask))
	sliceHeader.Len = NdpiBitmaskSize
	sliceHeader.Cap = NdpiBitmaskSize
	sliceHeader.Data = uintptr(unsafe.Pointer(&ndpiBitmask.fds_bits[0]))

	return bitmask
}

func NdpiProtocolBitmaskIsSet(bitmask []uint32, proto uint16) bool {
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

type NdpiProtoDefaults struct {
	ProtoName        string
	ProtoCategory    uint16
	IsClearTextProto bool
	IsAppProtocol    bool
	SubProtocols     []uint16
	ProtoId          uint16
	ProtoIdx         uint16
	TcpDefaultPorts  []uint16
	UdpDefaultPorts  []uint16
	ProtoBreed       uint16
}

type NdpiDetectionModule struct {
	ndpi ndpiDetectionModuleStructPtr
	mu   sync.Mutex
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
	dm.ndpi = ndpi

	return dm, nil
}

func NdpiDetectionModuleExit(dm *NdpiDetectionModule) {
	C.ndpi_detection_module_exit(dm.ndpi)
}

func (dm *NdpiDetectionModule) SetLogLevel(level uint16) {
	C.ndpi_set_log_level(dm.ndpi, C.uint(level))
}

func (dm *NdpiDetectionModule) SetDebugBitmask(debugBitmask []uint32) {
	ndpiBitmask := C.NDPI_PROTOCOL_BITMASK{}
	ndpiBitmask.fds_bits = *(*[NdpiBitmaskSize]C.uint32_t)(unsafe.Pointer(&debugBitmask[0]))

	C.ndpi_set_debug_bitmask(dm.ndpi, ndpiBitmask)
}

func (dm *NdpiDetectionModule) GetProtoDefaults() []NdpiProtoDefaults {
	isClearTextProtoList := make([]bool, C.NDPI_MAX_SUPPORTED_PROTOCOLS+C.NDPI_MAX_NUM_CUSTOM_PROTOCOLS)
	isAppProtocolList := make([]bool, C.NDPI_MAX_SUPPORTED_PROTOCOLS+C.NDPI_MAX_NUM_CUSTOM_PROTOCOLS)

	protoDefaults := C.ndpi_proto_defaults_get(dm.ndpi, (*C.bool)(unsafe.Pointer(&isClearTextProtoList[0])),
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

		subProtocols := make([]uint16, 0)
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
			ProtoCategory:    uint16(pds[i].protoCategory),
			IsClearTextProto: isClearTextProtoList[i],
			IsAppProtocol:    isAppProtocolList[i],
			SubProtocols:     subProtocols,
			ProtoId:          uint16(pds[i].protoId),
			ProtoIdx:         uint16(pds[i].protoIdx),
			TcpDefaultPorts:  tcpDefaultPorts,
			UdpDefaultPorts:  udpDefaultPorts,
			ProtoBreed:       uint16(pds[i].protoBreed),
		}

		npds = append(npds, npd)
	}

	return npds
}

type NdpiFlow struct {
	ndpiFlow ndpiFlowStructPtr
	mu       sync.Mutex
}

func NewNdpiFlow() (*NdpiFlow, error) {
	ndpiFlow := C.ndpi_flow_struct_malloc()
	if ndpiFlow == nil {
		C.ndpi_flow_struct_free(ndpiFlow)
		err := errors.New("null ndpi flow struct")
		return nil, err
	}

	f := &NdpiFlow{}
	f.ndpiFlow = ndpiFlow

	return f, nil
}

func FreeNdpiFlow(f *NdpiFlow) {
	C.ndpi_flow_struct_free(f.ndpiFlow)
}

func (f *NdpiFlow) GetDetectedProtocolStack() [2]uint16 {
	protoStack := [2]uint16{}
	protoStack[0] = uint16(f.ndpiFlow.detected_protocol_stack[0])
	protoStack[1] = uint16(f.ndpiFlow.detected_protocol_stack[1])

	return protoStack
}

func (f *NdpiFlow) GetProcessedPktNum() uint16 {
	return uint16(f.ndpiFlow.num_processed_pkts)
}

func (f *NdpiFlow) GetFlowExtraInfo() string {
	return C.GoString(&f.ndpiFlow.flow_extra_info[0])
}

func (f *NdpiFlow) GetHostServerName() string {
	return C.GoString(&f.ndpiFlow.host_server_name[0])
}

func (f *NdpiFlow) GetHttp() types.NdpiHttp {
	return types.NdpiHttp{
		NdpiHttpMethod:      uint16(f.ndpiFlow.http.method),
		RequestVersion:      uint8(f.ndpiFlow.http.request_version),
		ResponseStatusCode:  uint16(f.ndpiFlow.http.response_status_code),
		Url:                 C.GoString(f.ndpiFlow.http.url),
		RequestContentType:  C.GoString(f.ndpiFlow.http.request_content_type),
		ResponseContentType: C.GoString(f.ndpiFlow.http.content_type),
		UserAgent:           C.GoString(f.ndpiFlow.http.user_agent),
		DetectedOs:          C.GoString(f.ndpiFlow.http.detected_os),
		XForwardedAddr:      C.GoString(f.ndpiFlow.http.nat_ip),
	}
}

func (f *NdpiFlow) GetProtocolCategory() uint16 {
	return NdpiCategoryToId(f.ndpiFlow.category)
}

type NdpiProto struct {
	MasterProtocolId uint16
	AppProtocolId    uint16
	CategoryId       uint16
}

func NdpiPacketProcessing(dm *NdpiDetectionModule, flow *NdpiFlow, ipPacket []byte, ipPacketLen uint16, timestamp int) NdpiProto {
	ipPktPtr := (*C.u_char)(unsafe.Pointer(&ipPacket[0]))
	ipPktLen := C.ushort(ipPacketLen)
	ipPktTs := C.uint64_t(timestamp)

	proto := C.ndpi_packet_processing(dm.ndpi, flow.ndpiFlow, ipPktPtr, ipPktLen, ipPktTs)
	// fmt.Printf("%v, %v, %v\n", proto.master_protocol, proto.app_protocol, proto.category)

	ndpiProto := NdpiProto{
		MasterProtocolId: uint16(proto.master_protocol),
		AppProtocolId:    uint16(proto.app_protocol),
		CategoryId:       NdpiCategoryToId(proto.category),
	}

	return ndpiProto
}

func NdpiCategoryToId(category C.ndpi_protocol_category_t) uint16 {
	return uint16(category)
}
