package main

/*
#cgo CFLAGS: -I/usr/include/
#cgo LDFLAGS: -Wl,-Bstatic -lndpi -lmaxminddb -Wl,-Bdynamic -lpcap -lm -pthread
#include "ndpi_linux.h"
*/
import "C"

import (
	"fmt"
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

func NdpiPacketProcessing(handle *NdpiHandle, ndpiFlow *NdpiFlowHandle, ipPacket []byte, ipPacketLen uint16, ts int) {
	ipPktPtr := (*C.u_char)(unsafe.Pointer(&ipPacket[0]))
	ipPktLen := C.ushort(ipPacketLen)
	ipPktTs := C.uint64_t(ts)

	proto := C.ndpi_packet_processing(handle.ndpi, ndpiFlow.ndpiFlow, ipPktPtr, ipPktLen, ipPktTs)
	fmt.Printf("%v, %v, %v\n", proto.master_protocol, proto.app_protocol, proto.category)
}
