package types

type NdpiFlowHttp struct {
	NdpiHttpMethod      uint16
	RequestVersion      uint8
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
