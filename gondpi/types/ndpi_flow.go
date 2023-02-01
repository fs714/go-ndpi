package types

import (
	"encoding/binary"
	"encoding/json"
	"net"
	"time"
)

type NdpiFlowHttp struct {
	NdpiHttpMethod      HttpMethod
	RequestVersion      HttpRequestVersion
	ResponseStatusCode  uint16
	Url                 string
	RequestContentType  string
	ResponseContentType string
	UserAgent           string
	DetectedOs          string
	XForwardedAddr      string
}

type NdpiFlowKerberosBuf struct {
	PktBuf        []byte
	PktBufMaxLen  uint16
	PktBufCurrLen uint16
}

type NdpiFlowStun struct {
	NumUdpPkts         uint8
	NumBindingRequests uint8
	NumProcessedPkts   uint16
}

type NdpiFlowDns struct {
	NumQueries uint8
	NumAnswers uint8
	ReplyCode  uint8
	IsQuery    uint8
	QueryType  uint16
	QueryClass uint16
	RspType    uint16
	RspAddr    uint32
}

type NdpiFlowNtp struct {
	RequestCode uint8
	Version     uint8
}

type NdpiFlowKerberos struct {
	Hostname string
	Domain   string
	Username string
}

type NdpiFlowQuic struct {
	ServerNames                string
	ALPN                       string
	TlsSupportedVersions       string
	IssuerDN                   string
	SubjectDN                  string
	NotBefore                  uint32
	NotAfter                   uint32
	Ja3Client                  string
	Ja3Server                  string
	ServerCipher               uint16
	Sha1CertificateFingerprint []uint8
	HelloProcessed             bool
	SupProtocolDetected        bool
	SslVersion                 uint16
	ServerNamesLen             uint16
}

type NdpiFlowSsh struct {
	ClientSignature []byte
	ServerSignature []byte
	HasshClient     []byte
	HasshServer     []byte
}

type NdpiFlowTftp struct {
	Filename string
}

type NdpiFlowTelnet struct {
	UsernameDetected bool
	UsernameFound    bool
	PasswordDetected bool
	PasswordFound    bool
	CharacterId      uint8
	Username         string
	Password         string
}

type NdpiFlowRsh struct {
	ClientUsername string
	ServerUsername string
	Command        string
}

type NdpiFlowCollectd struct {
	ClientUsername string
}

type NdpiFlowUbntac2 struct {
	Version string
}

type NdpiFlowBittorrent struct {
	Hash []byte
}

type NdpiFlowDhcp struct {
	Fingerprint []byte
	ClassIdent  []byte
}

type NdpiFlowSnmp struct {
	Version     uint8
	Primitive   uint8
	ErrorStatus uint8
}

type NdpiFlowInfo struct {
	DetectedProtocolStack   [2]NdpiProtocol
	GuessedProtocolId       NdpiProtocol
	GuessedHostProtocolId   NdpiProtocol
	GuessedCategoryId       NdpiCategory
	GuessedHeaderCategoryId NdpiCategory
	L4Protocol              IPProto
	IsProtoIdAlreadyGuessed bool
	IsHostAlreadyGuessed    bool
	IsFailWithUnknow        bool
	IsInitFinished          bool
	IsSetupPacketDirection  bool
	IsPacketDirection       bool
	IsCheckExtraPackets     bool
	IsIpv6                  bool
	Confidence              NdpiConfidence
	NextTcpSeqNR            [2]uint16
	SrcAddr                 uint32
	DstAddr                 uint32
	SrcPort                 uint16
	DstPort                 uint16
	MaxExtraPacketsToCheck  uint8
	NumExtraPacketsChecked  uint8
	ProcessedPktNum         uint16
	LastPacketTimeMS        uint64
	Entropy                 float32
	FlowExtraInfo           string
	HostServerName          string
	Http                    NdpiFlowHttp
	KerberosBuf             NdpiFlowKerberosBuf
	Stun                    NdpiFlowStun
	ProtocolCategory        NdpiCategory
}

func (f *NdpiFlowInfo) ToString() (string, error) {
	flowInfoExtend := struct {
		NdpiFlowInfo
		DetectedProtocolStackString [2]string
		GuessedProtocolString       string
		GuessedHostProtocolString   string
		GuessedCategoryString       string
		GuessedHeaderCategoryString string
		L4ProtocolString            string
		SrcAddrString               string
		DstAddrString               string
		LastPacketTimeMSString      string
		ProtocolCategoryString      string
	}{
		NdpiFlowInfo:                *f,
		DetectedProtocolStackString: [2]string{f.DetectedProtocolStack[0].ToName(), f.DetectedProtocolStack[1].ToName()},
		GuessedProtocolString:       f.GuessedProtocolId.ToName(),
		GuessedHostProtocolString:   f.GuessedHostProtocolId.ToName(),
		GuessedCategoryString:       f.GuessedCategoryId.ToName(),
		GuessedHeaderCategoryString: f.GuessedHeaderCategoryId.ToName(),
		L4ProtocolString:            f.L4Protocol.ToName(),
		SrcAddrString:               IntToIPv4(f.SrcAddr).String(),
		DstAddrString:               IntToIPv4(f.DstAddr).String(),
		LastPacketTimeMSString:      time.Unix(0, int64(f.LastPacketTimeMS)*1000000).Format(time.RFC3339Nano),
		ProtocolCategoryString:      f.ProtocolCategory.ToName(),
	}

	flowInfoExtendJson, err := json.MarshalIndent(flowInfoExtend, "", "  ")
	if err != nil {
		return "", err
	}

	return string(flowInfoExtendJson), nil
}

func IPv4ToInt(ip net.IP) uint32 {
	return binary.BigEndian.Uint32(ip.To4())
}

func IntToIPv4(n uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, n)
	return ip
}
