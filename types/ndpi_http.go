package types

type HttpMethod string

const (
	HTTP_METHOD_UNKNOWN      HttpMethod = "UNKNOWN"
	HTTP_METHOD_OPTIONS      HttpMethod = "OPTIONS"
	HTTP_METHOD_GET          HttpMethod = "GET"
	HTTP_METHOD_HEAD         HttpMethod = "HEAD"
	HTTP_METHOD_PATCH        HttpMethod = "PATCH"
	HTTP_METHOD_POST         HttpMethod = "POST"
	HTTP_METHOD_PUT          HttpMethod = "PUT"
	HTTP_METHOD_DELETE       HttpMethod = "DELETE"
	HTTP_METHOD_TRACE        HttpMethod = "TRACE"
	HTTP_METHOD_CONNECT      HttpMethod = "CONNECT"
	HTTP_METHOD_RPC_IN_DATA  HttpMethod = "RPC_IN_DATA"
	HTTP_METHOD_RPC_OUT_DATA HttpMethod = "RPC_OUT_DATA"
)

const (
	NDPI_HTTP_METHOD_UNKNOWN      = 0
	NDPI_HTTP_METHOD_OPTIONS      = 1
	NDPI_HTTP_METHOD_GET          = 2
	NDPI_HTTP_METHOD_HEAD         = 3
	NDPI_HTTP_METHOD_PATCH        = 4
	NDPI_HTTP_METHOD_POST         = 5
	NDPI_HTTP_METHOD_PUT          = 6
	NDPI_HTTP_METHOD_DELETE       = 7
	NDPI_HTTP_METHOD_TRACE        = 8
	NDPI_HTTP_METHOD_CONNECT      = 9
	NDPI_HTTP_METHOD_RPC_IN_DATA  = 10
	NDPI_HTTP_METHOD_RPC_OUT_DATA = 11
)

var NdpiHttpMethodIdMap = map[uint16]HttpMethod{
	NDPI_HTTP_METHOD_UNKNOWN:      HTTP_METHOD_UNKNOWN,
	NDPI_HTTP_METHOD_OPTIONS:      HTTP_METHOD_OPTIONS,
	NDPI_HTTP_METHOD_GET:          HTTP_METHOD_GET,
	NDPI_HTTP_METHOD_HEAD:         HTTP_METHOD_HEAD,
	NDPI_HTTP_METHOD_PATCH:        HTTP_METHOD_PATCH,
	NDPI_HTTP_METHOD_POST:         HTTP_METHOD_POST,
	NDPI_HTTP_METHOD_PUT:          HTTP_METHOD_PUT,
	NDPI_HTTP_METHOD_DELETE:       HTTP_METHOD_DELETE,
	NDPI_HTTP_METHOD_TRACE:        HTTP_METHOD_TRACE,
	NDPI_HTTP_METHOD_CONNECT:      HTTP_METHOD_CONNECT,
	NDPI_HTTP_METHOD_RPC_IN_DATA:  HTTP_METHOD_RPC_IN_DATA,
	NDPI_HTTP_METHOD_RPC_OUT_DATA: HTTP_METHOD_RPC_OUT_DATA,
}

type HttpRequestVersion string

const (
	HTTP_Request_Version_1_0 HttpRequestVersion = "1.0"
	HTTP_Request_Version_1_1 HttpRequestVersion = "1.1"
)

const (
	NDPI_HTTP_Request_Version_1_0 = 0
	NDPI_HTTP_Request_Version_1_1 = 1
)

var NdpiHttpRequestVersionIdMap = map[uint8]HttpRequestVersion{
	NDPI_HTTP_Request_Version_1_0: HTTP_Request_Version_1_0,
	NDPI_HTTP_Request_Version_1_1: HTTP_Request_Version_1_1,
}

type NdpiHttp struct {
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
