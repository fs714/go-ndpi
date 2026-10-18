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
	Server              string
	Referer             string
	Host                string
	DetectedOs          string
	XForwardedAddr      string
	Filename            string
}

type NdpiFlowKerberosBuf struct {
	PktBuf        []byte
	PktBufMaxLen  uint16
	PktBufCurrLen uint16
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
	DetectedProtocolStack    [2]NdpiProtocol
	ClassificationState      NdpiClassificationState
	GuessedProtocolId        NdpiProtocol
	GuessedProtocolIdByIp    NdpiProtocol
	GuessedHeaderCategoryId  NdpiCategory
	L4Protocol               IPProto
	IsProtoIdAlreadyGuessed  bool
	IsProtocolWasGuessed     bool
	IsInitFinished           bool
	IsClientPacketDirection  bool
	IsPacketDirection        bool
	IsIpv6                   bool
	IsFirstPktFullyEncrypted bool
	IsAlreadyGaveup          bool
	Confidence               NdpiConfidence
	ClientAddr               uint32
	ServerAddr               uint32
	ClientPort               uint16
	ServerPort               uint16
	MaxExtraPacketsToCheck   uint16
	NumExtraPacketsChecked   uint16
	ProcessedPktNum          uint16
	LastPacketTimeMS         uint64
	Entropy                  float32
	HostServerName           string
	Risk                     uint64
	Http                     NdpiFlowHttp
	KerberosBuf              NdpiFlowKerberosBuf
	ProtocolCategory         NdpiCategory
	ProtocolBreed            NdpiProtocolBreed
}

func (f *NdpiFlowInfo) ToString() string {
	flowInfoExtend := struct {
		NdpiFlowInfo
		DetectedProtocolStackString [2]string
		ClassificationStateString   string
		GuessedProtocolString       string
		GuessedProtocolByIpString   string
		GuessedHeaderCategoryString string
		L4ProtocolString            string
		ConfidenceString            string
		ClientAddrString            string
		ServerAddrString            string
		LastPacketTimeMSString      string
		ProtocolCategoryString      string
		ProtocolBreedString         string
	}{
		NdpiFlowInfo:                *f,
		DetectedProtocolStackString: [2]string{f.DetectedProtocolStack[0].ToName(), f.DetectedProtocolStack[1].ToName()},
		ClassificationStateString:   f.ClassificationState.ToName(),
		GuessedProtocolString:       f.GuessedProtocolId.ToName(),
		GuessedProtocolByIpString:   f.GuessedProtocolIdByIp.ToName(),
		GuessedHeaderCategoryString: f.GuessedHeaderCategoryId.ToName(),
		L4ProtocolString:            f.L4Protocol.ToName(),
		ConfidenceString:            f.Confidence.ToName(),
		ClientAddrString:            IntToIPv4(f.ClientAddr).String(),
		ServerAddrString:            IntToIPv4(f.ServerAddr).String(),
		LastPacketTimeMSString:      time.Unix(0, int64(f.LastPacketTimeMS)*1000000).Format(time.RFC3339Nano),
		ProtocolCategoryString:      f.ProtocolCategory.ToName(),
		ProtocolBreedString:         f.ProtocolBreed.ToName(),
	}

	flowInfoExtendJson, _ := json.MarshalIndent(flowInfoExtend, "", "  ")

	return string(flowInfoExtendJson)
}

func IPv4ToInt(ip net.IP) uint32 {
	return binary.BigEndian.Uint32(ip.To4())
}

func IntToIPv4(n uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, n)
	return ip
}
