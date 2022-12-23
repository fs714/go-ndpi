package main

/*
#cgo CFLAGS: -I/usr/include/
#cgo LDFLAGS: -Wl,-Bstatic -lndpi -lmaxminddb -Wl,-Bdynamic -lpcap -lm -pthread
#include "ndpi_linux.h"
*/
import "C"

import (
	"sync"
	"unsafe"

	"github.com/pkg/errors"
)

const (
	NdpiFlowStructSize uint32 = uint32(C.SIZEOF_FLOW_STRUCT)
)

type ndpiStructPtr *C.struct_ndpi_detection_module_struct
type ndpiFlowStructPtr *C.struct_ndpi_flow_struct

type NdpiHandle struct {
	ndpi ndpiStructPtr
	mu   sync.Mutex
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

func (f *NdpiFlowHandle) GetHttp() NdpiHttp {
	return NdpiHttp{
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

func NdpiHandleInitialize() (*NdpiHandle, error) {
	ndpi := C.ndpi_struct_initialize()
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
