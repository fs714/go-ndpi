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

	"github.com/fs714/go-ndpi/types"
	"github.com/pkg/errors"
)

const (
	NdpiFlowStructSize uint32 = uint32(C.SIZEOF_FLOW_STRUCT)
)

type ndpiStructPtr *C.struct_ndpi_detection_module_struct
type ndpiFlowStructPtr *C.struct_ndpi_flow_struct

func NewNdpiProtocolBitmask() []uint32 {
	return make([]uint32, 16)
}

func NdpiProtocolBitmaskAdd(bitmask []uint32, proto uint16) []uint32 {
	ndpiBitmask := &C.NDPI_PROTOCOL_BITMASK{}
	ndpiBitmask.fds_bits = *(*[16]C.uint32_t)(unsafe.Pointer(&bitmask[0]))

	C.ndpi_protocol_bitmask_add(ndpiBitmask, C.uint16_t(proto))

	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&bitmask)))
	sliceHeader.Data = uintptr(unsafe.Pointer(&ndpiBitmask.fds_bits))

	return bitmask
}

func NdpiProtocolBitmaskDel(bitmask []uint32, proto uint16) []uint32 {
	ndpiBitmask := &C.NDPI_PROTOCOL_BITMASK{}
	ndpiBitmask.fds_bits = *(*[16]C.uint32_t)(unsafe.Pointer(&bitmask[0]))

	C.ndpi_protocol_bitmask_del(ndpiBitmask, C.uint16_t(proto))

	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&bitmask)))
	sliceHeader.Data = uintptr(unsafe.Pointer(&ndpiBitmask.fds_bits))

	return bitmask
}

func NdpiProtocolBitmaskIsSet(bitmask []uint32, proto uint16) bool {
	ndpiBitmask := &C.NDPI_PROTOCOL_BITMASK{}
	ndpiBitmask.fds_bits = *(*[16]C.uint32_t)(unsafe.Pointer(&bitmask[0]))

	ndpiIsSet := C.ndpi_protocol_bitmask_is_set(ndpiBitmask, C.uint16_t(proto))

	return bool(ndpiIsSet)
}

func NdpiProtocolBitmaskReset(bitmask []uint32) []uint32 {
	ndpiBitmask := &C.NDPI_PROTOCOL_BITMASK{}
	ndpiBitmask.fds_bits = *(*[16]C.uint32_t)(unsafe.Pointer(&bitmask[0]))

	C.ndpi_protocol_bitmask_reset(ndpiBitmask)

	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&bitmask)))
	sliceHeader.Data = uintptr(unsafe.Pointer(&ndpiBitmask.fds_bits))

	return bitmask
}

func NdpiProtocolBitmaskSetAll(bitmask []uint32) []uint32 {
	ndpiBitmask := &C.NDPI_PROTOCOL_BITMASK{}
	ndpiBitmask.fds_bits = *(*[16]C.uint32_t)(unsafe.Pointer(&bitmask[0]))

	C.ndpi_protocol_bitmask_set_all(ndpiBitmask)

	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&bitmask)))
	sliceHeader.Data = uintptr(unsafe.Pointer(&ndpiBitmask.fds_bits))

	return bitmask
}

type NdpiHandle struct {
	ndpi ndpiStructPtr
	mu   sync.Mutex
}

type NdpiProtoDefaults struct {
	ProtoName        string
	ProtoCategory    string
	IsClearTextProto bool
	IsAppProtocol    bool
	SubProtocols     []uint16
	ProtoId          uint16
	ProtoIdx         uint16
	TcpDefaultPorts  []uint16
	UdpDefaultPorts  []uint16
	ProtoBreed       string
}

func (h *NdpiHandle) GetProtoDefaults() []NdpiProtoDefaults {
	isClearTextProtoList := make([]bool, C.NDPI_MAX_SUPPORTED_PROTOCOLS+C.NDPI_MAX_NUM_CUSTOM_PROTOCOLS)
	isAppProtocolList := make([]bool, C.NDPI_MAX_SUPPORTED_PROTOCOLS+C.NDPI_MAX_NUM_CUSTOM_PROTOCOLS)

	protoDefaults := C.ndpi_get_proto_defaults_wrapper(h.ndpi, (*C.bool)(unsafe.Pointer(&isClearTextProtoList[0])),
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
		sliceHeader.Data = uintptr(unsafe.Pointer(&pds[i].tcp_default_ports))

		udpDefaultPorts := make([]uint16, 0)
		sliceHeader = (*reflect.SliceHeader)((unsafe.Pointer(&udpDefaultPorts)))
		sliceHeader.Cap = C.MAX_DEFAULT_PORTS
		sliceHeader.Len = C.MAX_DEFAULT_PORTS
		sliceHeader.Data = uintptr(unsafe.Pointer(&pds[i].udp_default_ports))

		npd := NdpiProtoDefaults{
			ProtoName:        C.GoString(pds[i].protoName),
			ProtoCategory:    string(types.NdpiCategoryIdMap[uint16(pds[i].protoCategory)]),
			IsClearTextProto: isClearTextProtoList[i],
			IsAppProtocol:    isAppProtocolList[i],
			SubProtocols:     subProtocols,
			ProtoId:          uint16(pds[i].protoId),
			ProtoIdx:         uint16(pds[i].protoIdx),
			TcpDefaultPorts:  tcpDefaultPorts,
			UdpDefaultPorts:  udpDefaultPorts,
			ProtoBreed:       string(types.NdpiProtocolBreedIdMap[uint16(pds[i].protoBreed)]),
		}

		npds = append(npds, npd)
	}

	return npds
}

type NdpiFlowHandle struct {
	ndpiFlow ndpiFlowStructPtr
	mu       sync.Mutex
}

func (f *NdpiFlowHandle) GetDetectedProtocolStack() [2]uint16 {
	protoStack := [2]uint16{}
	protoStack[0] = uint16(f.ndpiFlow.detected_protocol_stack[0])
	protoStack[1] = uint16(f.ndpiFlow.detected_protocol_stack[1])

	return protoStack
}

func (f *NdpiFlowHandle) GetProcessedPktNum() uint16 {
	return uint16(f.ndpiFlow.num_processed_pkts)
}

func (f *NdpiFlowHandle) GetFlowExtraInfo() string {
	return C.GoString(&f.ndpiFlow.flow_extra_info[0])
}

func (f *NdpiFlowHandle) GetHostServerName() string {
	return C.GoString(&f.ndpiFlow.host_server_name[0])
}

func (f *NdpiFlowHandle) GetHttp() types.NdpiHttp {
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

func (f *NdpiFlowHandle) GetProtocolCategory() uint16 {
	return NdpiCategoryToId(f.ndpiFlow.category)
}

type NdpiProto struct {
	MasterProtocolId uint16
	AppProtocolId    uint16
	CategoryId       uint16
}

func NdpiHandleInitialize(detectionBitmask []uint32) (*NdpiHandle, error) {
	ndpiBitmask := &C.NDPI_PROTOCOL_BITMASK{}
	ndpiBitmask.fds_bits = *(*[16]C.uint32_t)(unsafe.Pointer(&detectionBitmask[0]))
	ndpi := C.ndpi_struct_initialize(ndpiBitmask)
	if ndpi == nil {
		C.ndpi_struct_exit(ndpi)
		err := errors.New("null ndpi struct")
		return nil, err
	}

	handle := &NdpiHandle{}
	handle.ndpi = ndpi

	return handle, nil
}

func NdpiHandleExit(h *NdpiHandle) {
	C.ndpi_struct_exit(h.ndpi)
}

func GetNdpiFlowHandle() (*NdpiFlowHandle, error) {
	ndpiFlow := C.get_ndpi_flow()
	if ndpiFlow == nil {
		C.free_ndpi_flow(ndpiFlow)
		err := errors.New("null ndpi flow struct")
		return nil, err
	}

	handle := &NdpiFlowHandle{}
	handle.ndpiFlow = ndpiFlow

	return handle, nil
}

func FreeNdpiFlowHandle(h *NdpiFlowHandle) {
	C.free_ndpi_flow(h.ndpiFlow)
}

func NdpiPacketProcessing(handle *NdpiHandle, ndpiFlow *NdpiFlowHandle, ipPacket []byte, ipPacketLen uint16, ts int) NdpiProto {
	ipPktPtr := (*C.u_char)(unsafe.Pointer(&ipPacket[0]))
	ipPktLen := C.ushort(ipPacketLen)
	ipPktTs := C.uint64_t(ts)

	proto := C.ndpi_packet_processing(handle.ndpi, ndpiFlow.ndpiFlow, ipPktPtr, ipPktLen, ipPktTs)
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
