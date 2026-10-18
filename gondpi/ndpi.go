// Golang wrapper of ndpi 6.0
//
// Debian packages needed:
// apt-get install libpcap-dev libmaxminddb-dev
package gondpi

/*
#cgo CFLAGS: -I/usr/include/
#cgo LDFLAGS: -Wl,-Bstatic -lndpi -Wl,-Bdynamic -lmaxminddb -lpcap -lm -ldl -pthread
#include "ndpi_linux.h"
*/
import "C"

import (
	"encoding/binary"
	"strconv"
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

// Protocol bitmask is handled on golang side only: since nDPI 4.10 protocols are
// enabled/disabled through ndpi_set_config(), which is done on module creation.

func NewNdpiProtocolBitmask() []uint32 {
	return make([]uint32, NdpiBitmaskSize)
}

func NdpiProtocolBitmaskAdd(bitmask []uint32, proto types.NdpiProtocol) []uint32 {
	bitmask[proto/32] |= 1 << (proto % 32)

	return bitmask
}

func NdpiProtocolBitmaskDel(bitmask []uint32, proto types.NdpiProtocol) []uint32 {
	bitmask[proto/32] &^= 1 << (proto % 32)

	return bitmask
}

func NdpiProtocolBitmaskIsSet(bitmask []uint32, proto types.NdpiProtocol) bool {
	return bitmask[proto/32]&(1<<(proto%32)) != 0
}

func NdpiProtocolBitmaskReset(bitmask []uint32) []uint32 {
	for i := range bitmask {
		bitmask[i] = 0
	}

	return bitmask
}

func NdpiProtocolBitmaskSetAll(bitmask []uint32) []uint32 {
	for i := range bitmask {
		bitmask[i] = 0xFFFFFFFF
	}

	return bitmask
}

type NdpiFlow struct {
	NdpiFlowPtr NdpiFlowStructPtr
	Mu          sync.Mutex
}

func NewNdpiFlow() (*NdpiFlow, error) {
	ndpiFlow := C.ndpi_flow_struct_malloc()
	if ndpiFlow == nil {
		err := errors.New("null ndpi flow struct")
		return nil, err
	}

	f := &NdpiFlow{}
	f.NdpiFlowPtr = ndpiFlow

	return f, nil
}

func (f *NdpiFlow) Close() {
	C.ndpi_flow_struct_free(f.NdpiFlowPtr)
}

func (f *NdpiFlow) GetDetectedProtocolStack() [2]types.NdpiProtocol {
	protoStack := [2]types.NdpiProtocol{}
	protoStack[0] = types.NdpiProtocol(f.NdpiFlowPtr.detected_protocol_stack[0])
	protoStack[1] = types.NdpiProtocol(f.NdpiFlowPtr.detected_protocol_stack[1])

	return protoStack
}

func (f *NdpiFlow) GetClassificationState() types.NdpiClassificationState {
	return types.NdpiClassificationState(f.NdpiFlowPtr.state)
}

func (f *NdpiFlow) GetGuessedProtocolId() types.NdpiProtocol {
	return types.NdpiProtocol(f.NdpiFlowPtr.guessed_protocol_id)
}

func (f *NdpiFlow) GetGuessedProtocolIdByIp() types.NdpiProtocol {
	return types.NdpiProtocol(f.NdpiFlowPtr.guessed_protocol_id_by_ip)
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

func (f *NdpiFlow) IsProtocolWasGuessed() bool {
	ret := C.ndpi_flow_get_protocol_was_guessed(f.NdpiFlowPtr)

	return (uint8(ret) != 0)
}

func (f *NdpiFlow) IsInitFinished() bool {
	ret := C.ndpi_flow_get_init_finished(f.NdpiFlowPtr)

	return (uint8(ret) != 0)
}

func (f *NdpiFlow) IsClientPacketDirection() bool {
	ret := C.ndpi_flow_get_client_packet_direction(f.NdpiFlowPtr)

	return (uint8(ret) != 0)
}

func (f *NdpiFlow) IsPacketDirection() bool {
	ret := C.ndpi_flow_get_packet_direction(f.NdpiFlowPtr)

	return (uint8(ret) != 0)
}

func (f *NdpiFlow) IsIpv6() bool {
	ret := C.ndpi_flow_get_is_ipv6(f.NdpiFlowPtr)

	return (uint8(ret) != 0)
}

func (f *NdpiFlow) IsFirstPktFullyEncrypted() bool {
	ret := C.ndpi_flow_get_first_pkt_fully_encrypted(f.NdpiFlowPtr)

	return (uint8(ret) != 0)
}

func (f *NdpiFlow) IsAlreadyGaveup() bool {
	ret := C.ndpi_flow_get_already_gaveup(f.NdpiFlowPtr)

	return (uint8(ret) != 0)
}

func (f *NdpiFlow) GetConfidence() types.NdpiConfidence {
	return types.NdpiConfidence(f.NdpiFlowPtr.confidence)
}

// Client/Server address (IPv4 only) and ports are converted from network to host byte order
func (f *NdpiFlow) GetClientAddr() uint32 {
	clientAddrBigEndian := make([]byte, 4)
	binary.BigEndian.PutUint32(clientAddrBigEndian, uint32(C.ndpi_flow_get_client_addr_v4(f.NdpiFlowPtr)))
	clientAddr := binary.LittleEndian.Uint32(clientAddrBigEndian)

	return clientAddr
}

func (f *NdpiFlow) GetServerAddr() uint32 {
	serverAddrBigEndian := make([]byte, 4)
	binary.BigEndian.PutUint32(serverAddrBigEndian, uint32(C.ndpi_flow_get_server_addr_v4(f.NdpiFlowPtr)))
	serverAddr := binary.LittleEndian.Uint32(serverAddrBigEndian)

	return serverAddr
}

func (f *NdpiFlow) GetClientPort() uint16 {
	clientPortBigEndian := make([]byte, 2)
	binary.BigEndian.PutUint16(clientPortBigEndian, uint16(f.NdpiFlowPtr.c_port))
	clientPort := binary.LittleEndian.Uint16(clientPortBigEndian)

	return clientPort
}

func (f *NdpiFlow) GetServerPort() uint16 {
	serverPortBigEndian := make([]byte, 2)
	binary.BigEndian.PutUint16(serverPortBigEndian, uint16(f.NdpiFlowPtr.s_port))
	serverPort := binary.LittleEndian.Uint16(serverPortBigEndian)

	return serverPort
}

func (f *NdpiFlow) GetMaxExtraPacketsToCheck() uint16 {
	return uint16(f.NdpiFlowPtr.max_extra_packets_to_check)
}

func (f *NdpiFlow) GetNumExtraPacketsChecked() uint16 {
	return uint16(f.NdpiFlowPtr.num_extra_packets_checked)
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

func (f *NdpiFlow) GetHostServerName() string {
	return C.GoString(&f.NdpiFlowPtr.host_server_name[0])
}

// GetRisk returns the bitmask of the issues found with this flow (bit index is ndpi_risk_enum)
func (f *NdpiFlow) GetRisk() uint64 {
	return uint64(f.NdpiFlowPtr.risk)
}

// TODO
// u_int8_t initial_binary_bytes[8], initial_binary_bytes_len;
// ndpi_risk risk_mask; /* Stores the flow risk mask for flow peers */
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
		Server:              C.GoString(f.NdpiFlowPtr.http.server),
		Referer:             C.GoString(f.NdpiFlowPtr.http.referer),
		Host:                C.GoString(f.NdpiFlowPtr.http.host),
		DetectedOs:          C.GoString(f.NdpiFlowPtr.http.detected_os),
		XForwardedAddr:      C.GoString(f.NdpiFlowPtr.http.nat_ip),
		Filename:            C.GoString(f.NdpiFlowPtr.http.filename),
	}
}

func (f *NdpiFlow) GetKerberosBuf() types.NdpiFlowKerberosBuf {
	var pktBuf []byte
	if f.NdpiFlowPtr.kerberos_buf.pktbuf != nil {
		pktBuf = C.GoBytes(unsafe.Pointer(f.NdpiFlowPtr.kerberos_buf.pktbuf), C.int(f.NdpiFlowPtr.kerberos_buf.pktbuf_currlen))
	}

	return types.NdpiFlowKerberosBuf{
		PktBuf:        pktBuf,
		PktBufMaxLen:  uint16(f.NdpiFlowPtr.kerberos_buf.pktbuf_maxlen),
		PktBufCurrLen: uint16(f.NdpiFlowPtr.kerberos_buf.pktbuf_currlen),
	}
}

// TODO
// union {
// ......
// }

func (f *NdpiFlow) GetProtocolCategory() types.NdpiCategory {
	return NdpiCategoryToId(f.NdpiFlowPtr.category)
}

func (f *NdpiFlow) GetProtocolBreed() types.NdpiProtocolBreed {
	return types.NdpiProtocolBreed(f.NdpiFlowPtr.breed)
}

func (f *NdpiFlow) ToNdpiFlowInfo() *types.NdpiFlowInfo {
	return &types.NdpiFlowInfo{
		DetectedProtocolStack:    f.GetDetectedProtocolStack(),
		ClassificationState:      f.GetClassificationState(),
		GuessedProtocolId:        f.GetGuessedProtocolId(),
		GuessedProtocolIdByIp:    f.GetGuessedProtocolIdByIp(),
		GuessedHeaderCategoryId:  f.GetGuessedHeaderCategoryId(),
		L4Protocol:               f.GetL4Protocol(),
		IsProtoIdAlreadyGuessed:  f.IsProtoIdAlreadyGuessed(),
		IsProtocolWasGuessed:     f.IsProtocolWasGuessed(),
		IsInitFinished:           f.IsInitFinished(),
		IsClientPacketDirection:  f.IsClientPacketDirection(),
		IsPacketDirection:        f.IsPacketDirection(),
		IsIpv6:                   f.IsIpv6(),
		IsFirstPktFullyEncrypted: f.IsFirstPktFullyEncrypted(),
		IsAlreadyGaveup:          f.IsAlreadyGaveup(),
		Confidence:               f.GetConfidence(),
		ClientAddr:               f.GetClientAddr(),
		ServerAddr:               f.GetServerAddr(),
		ClientPort:               f.GetClientPort(),
		ServerPort:               f.GetServerPort(),
		MaxExtraPacketsToCheck:   f.GetMaxExtraPacketsToCheck(),
		NumExtraPacketsChecked:   f.GetNumExtraPacketsChecked(),
		ProcessedPktNum:          f.GetProcessedPktNum(),
		LastPacketTimeMS:         f.GetLastPacketTimeMS(),
		Entropy:                  f.GetEntropy(),
		HostServerName:           f.GetHostServerName(),
		Risk:                     f.GetRisk(),
		Http:                     f.GetHttp(),
		KerberosBuf:              f.GetKerberosBuf(),
		ProtocolCategory:         f.GetProtocolCategory(),
		ProtocolBreed:            f.GetProtocolBreed(),
	}
}

type NdpiProto struct {
	MasterProtocolId types.NdpiProtocol
	AppProtocolId    types.NdpiProtocol
	CategoryId       types.NdpiCategory
	Breed            types.NdpiProtocolBreed
	State            types.NdpiClassificationState
}

func toNdpiProto(proto C.ndpi_protocol) NdpiProto {
	return NdpiProto{
		MasterProtocolId: types.NdpiProtocol(proto.proto.master_protocol),
		AppProtocolId:    types.NdpiProtocol(proto.proto.app_protocol),
		CategoryId:       NdpiCategoryToId(proto.category),
		Breed:            types.NdpiProtocolBreed(proto.breed),
		State:            types.NdpiClassificationState(proto.state),
	}
}

type NdpiPortRange struct {
	Low  uint16
	High uint16
}

type NdpiProtoDefaults struct {
	ProtoName        string
	ProtoCategory    types.NdpiCategory
	IsClearTextProto bool
	IsAppProtocol    bool
	IsCustomProto    bool
	SubProtocols     []types.NdpiProtocol
	ProtoId          types.NdpiProtocol
	DissectorIdx     uint16
	TcpDefaultPorts  []NdpiPortRange
	UdpDefaultPorts  []NdpiPortRange
	ProtoBreed       types.NdpiProtocolBreed
}

type NdpiDetectionModule struct {
	NdpiPtr   NdpiDetectionModuleStructPtr
	Mu        sync.Mutex
	finalized bool
}

// NewNdpiDetectionModule creates a detection module which is not finalized yet:
// configuration (SetConfig, SetLogLevel, SetDebugBitmask) and custom files (Load*)
// must be applied before calling Finalize. Protocols not set in detectionBitmask are disabled.
func NewNdpiDetectionModule(license types.NdpiLicenseType, detectionBitmask []uint32) (*NdpiDetectionModule, error) {
	ndpi := C.ndpi_detection_module_create(C.enum_ndpi_license_type(license),
		(*C.uint32_t)(unsafe.Pointer(&detectionBitmask[0])), C.uint32_t(len(detectionBitmask)))
	if ndpi == nil {
		err := errors.New("null ndpi detection module struct")
		return nil, err
	}

	dm := &NdpiDetectionModule{}
	dm.NdpiPtr = ndpi

	return dm, nil
}

// NdpiDetectionModuleInitialize creates and finalizes a detection module with default configuration
func NdpiDetectionModuleInitialize(license types.NdpiLicenseType, detectionBitmask []uint32) (*NdpiDetectionModule, error) {
	dm, err := NewNdpiDetectionModule(license, detectionBitmask)
	if err != nil {
		return nil, err
	}

	err = dm.Finalize()
	if err != nil {
		dm.Close()
		return nil, err
	}

	return dm, nil
}

func (dm *NdpiDetectionModule) Finalize() error {
	ret := C.ndpi_finalize_initialization(dm.NdpiPtr)
	if int(ret) != 0 {
		err := errors.Errorf("failed to finalize ndpi detection module with code %d", int(ret))
		return err
	}

	dm.finalized = true

	return nil
}

func (dm *NdpiDetectionModule) IsFinalized() bool {
	return dm.finalized
}

func (dm *NdpiDetectionModule) Close() {
	C.ndpi_detection_module_exit(dm.NdpiPtr)
}

func ndpiCfgErrorToString(cfgErr C.ndpi_cfg_error) string {
	switch cfgErr {
	case C.NDPI_CFG_INVALID_CONTEXT:
		return "invalid context"
	case C.NDPI_CFG_NOT_FOUND:
		return "configuration parameter not found"
	case C.NDPI_CFG_INVALID_PARAM:
		return "invalid configuration parameter"
	case C.NDPI_CFG_CONTEXT_ALREADY_INITIALIZED:
		return "configuration not allowed on an already finalized context"
	case C.NDPI_CFG_CALLBACK_ERROR:
		return "configuration callback error"
	default:
		return "error " + strconv.Itoa(int(cfgErr))
	}
}

// SetConfig sets a nDPI configuration parameter, see doc/configuration.md in nDPI.
// Empty proto means a global parameter. It must be called before Finalize.
func (dm *NdpiDetectionModule) SetConfig(proto string, param string, value string) error {
	var cProto *C.char
	if proto != "" {
		cProto = C.CString(proto)
		defer C.free(unsafe.Pointer(cProto))
	}
	cParam := C.CString(param)
	defer C.free(unsafe.Pointer(cParam))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	ret := C.ndpi_set_config(dm.NdpiPtr, cProto, cParam, cValue)
	if ret != C.NDPI_CFG_OK {
		err := errors.Errorf("failed to set config [%s][%s][%s]: %s", proto, param, value, ndpiCfgErrorToString(ret))
		return err
	}

	return nil
}

func (dm *NdpiDetectionModule) GetConfig(proto string, param string) (string, error) {
	var cProto *C.char
	if proto != "" {
		cProto = C.CString(proto)
		defer C.free(unsafe.Pointer(cProto))
	}
	cParam := C.CString(param)
	defer C.free(unsafe.Pointer(cParam))

	buf := make([]byte, 128)
	ret := C.ndpi_get_config(dm.NdpiPtr, cProto, cParam, (*C.char)(unsafe.Pointer(&buf[0])), C.int(len(buf)))
	if ret == nil {
		err := errors.Errorf("failed to get config [%s][%s]", proto, param)
		return "", err
	}

	return C.GoString(ret), nil
}

// SetLogLevel must be called before Finalize, level is one of types.NDPI_LOG_*
func (dm *NdpiDetectionModule) SetLogLevel(level uint16) error {
	return dm.SetConfig("", "log.level", strconv.Itoa(int(level)))
}

// SetDebugBitmask enables logs for protocols set in debugBitmask. It must be called before Finalize
func (dm *NdpiDetectionModule) SetDebugBitmask(debugBitmask []uint32) {
	for id := 1; id < len(debugBitmask)*32; id++ {
		value := "disable"
		if NdpiProtocolBitmaskIsSet(debugBitmask, types.NdpiProtocol(id)) {
			value = "enable"
		}

		// Ids not matching any internal protocol are refused by nDPI: ignore them
		_ = dm.SetConfig(strconv.Itoa(id), "log", value)
	}
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

func (dm *NdpiDetectionModule) LoadCategoriesFile(path string, userData string) (uint32, error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	// userData is kept by nDPI as custom category user data, so it is not freed
	cUserData := C.CString(userData)

	num := C.ndpi_load_categories_file(dm.NdpiPtr, cPath, unsafe.Pointer(cUserData))
	if int(num) < 0 {
		C.free(unsafe.Pointer(cUserData))
		err := errors.New("failed to load categories file")
		return 0, err
	}

	return uint32(num), nil
}

func (dm *NdpiDetectionModule) LoadRiskDomainFile(path string) (uint32, error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	num := C.ndpi_load_risk_domain_file(dm.NdpiPtr, cPath)
	if int(num) < 0 {
		err := errors.New("failed to load risk domain file")
		return 0, err
	}

	return uint32(num), nil
}

func (dm *NdpiDetectionModule) LoadMaliciousJa4File(path string) (uint32, error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	num := C.ndpi_load_malicious_ja4_file(dm.NdpiPtr, cPath)
	if int(num) < 0 {
		err := errors.New("failed to load malicious ja4 file")
		return 0, err
	}

	return uint32(num), nil
}

func (dm *NdpiDetectionModule) LoadMaliciousSha1File(path string) (uint32, error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	num := C.ndpi_load_malicious_sha1_file(dm.NdpiPtr, cPath)
	if int(num) < 0 {
		err := errors.New("failed to load malicious sha1 file")
		return 0, err
	}

	return uint32(num), nil
}

// GetNumProtocols returns the number of supported protocols (internal and custom), 0 if not finalized
func (dm *NdpiDetectionModule) GetNumProtocols() uint32 {
	return uint32(C.ndpi_get_num_protocols(dm.NdpiPtr))
}

func toNdpiPortRanges(ports *C.ndpi_port_range) []NdpiPortRange {
	cPorts := unsafe.Slice(ports, C.MAX_DEFAULT_PORTS)
	portRanges := make([]NdpiPortRange, 0, len(cPorts))
	for _, p := range cPorts {
		portRanges = append(portRanges, NdpiPortRange{Low: uint16(p.port_low), High: uint16(p.port_high)})
	}

	return portRanges
}

// GetProtoDefaults must be called after Finalize
func (dm *NdpiDetectionModule) GetProtoDefaults() []NdpiProtoDefaults {
	numProtocols := dm.GetNumProtocols()
	protoDefaults := C.ndpi_get_proto_defaults(dm.NdpiPtr)
	if protoDefaults == nil || numProtocols == 0 {
		return []NdpiProtoDefaults{}
	}
	pds := unsafe.Slice(protoDefaults, numProtocols)

	npds := make([]NdpiProtoDefaults, 0, numProtocols)
	for i := range pds {
		pd := &pds[i]
		protoName := C.GoString(&pd.protoName[0])
		if protoName == "" {
			continue
		}

		subProtocols := make([]types.NdpiProtocol, 0, int(pd.subprotocol_count))
		if pd.subprotocol_count > 0 && pd.subprotocols != nil {
			for _, sp := range unsafe.Slice(pd.subprotocols, pd.subprotocol_count) {
				subProtocols = append(subProtocols, types.NdpiProtocol(sp))
			}
		}

		npd := NdpiProtoDefaults{
			ProtoName:        protoName,
			ProtoCategory:    types.NdpiCategory(pd.protoCategory),
			IsClearTextProto: uint8(C.ndpi_proto_defaults_get_is_clear_text_proto(pd)) != 0,
			IsAppProtocol:    uint8(C.ndpi_proto_defaults_get_is_app_protocol(pd)) != 0,
			IsCustomProto:    uint8(C.ndpi_proto_defaults_get_is_custom_proto(pd)) != 0,
			SubProtocols:     subProtocols,
			ProtoId:          types.NdpiProtocol(pd.protoId),
			DissectorIdx:     uint16(pd.dissector_idx),
			TcpDefaultPorts:  toNdpiPortRanges(&pd.tcp_default_ports[0]),
			UdpDefaultPorts:  toNdpiPortRanges(&pd.udp_default_ports[0]),
			ProtoBreed:       types.NdpiProtocolBreed(pd.protoBreed),
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

	return toNdpiProto(proto)
}

// IsExtraDissectionPossible reports whether nDPI wants more packets of this flow to extract metadata
func (dm *NdpiDetectionModule) IsExtraDissectionPossible(flow *NdpiFlow) bool {
	ret := C.ndpi_flow_get_extra_packets_func_set(flow.NdpiFlowPtr)

	return (uint8(ret) != 0)
}

// DetectionGiveup returns the detected protocol and whether it was guessed.
// Guessing is controlled by nDPI configuration (e.g. "dpi.guess_on_giveup")
func (dm *NdpiDetectionModule) DetectionGiveup(flow *NdpiFlow) (NdpiProto, bool) {
	proto := C.ndpi_detection_giveup(dm.NdpiPtr, flow.NdpiFlowPtr)

	return toNdpiProto(proto), flow.IsProtocolWasGuessed()
}
