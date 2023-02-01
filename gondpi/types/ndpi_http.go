package types

type HttpMethod uint16

func (m *HttpMethod) ToName() string {
	name, ok := NdpiHttpMethodNameMap[*m]
	if !ok {
		name = ""
	}

	return name
}

const (
	NDPI_HTTP_METHOD_UNKNOWN      HttpMethod = 0
	NDPI_HTTP_METHOD_OPTIONS      HttpMethod = 1
	NDPI_HTTP_METHOD_GET          HttpMethod = 2
	NDPI_HTTP_METHOD_HEAD         HttpMethod = 3
	NDPI_HTTP_METHOD_PATCH        HttpMethod = 4
	NDPI_HTTP_METHOD_POST         HttpMethod = 5
	NDPI_HTTP_METHOD_PUT          HttpMethod = 6
	NDPI_HTTP_METHOD_DELETE       HttpMethod = 7
	NDPI_HTTP_METHOD_TRACE        HttpMethod = 8
	NDPI_HTTP_METHOD_CONNECT      HttpMethod = 9
	NDPI_HTTP_METHOD_RPC_IN_DATA  HttpMethod = 10
	NDPI_HTTP_METHOD_RPC_OUT_DATA HttpMethod = 11
)

const (
	HTTP_METHOD_UNKNOWN      string = "UNKNOWN"
	HTTP_METHOD_OPTIONS      string = "OPTIONS"
	HTTP_METHOD_GET          string = "GET"
	HTTP_METHOD_HEAD         string = "HEAD"
	HTTP_METHOD_PATCH        string = "PATCH"
	HTTP_METHOD_POST         string = "POST"
	HTTP_METHOD_PUT          string = "PUT"
	HTTP_METHOD_DELETE       string = "DELETE"
	HTTP_METHOD_TRACE        string = "TRACE"
	HTTP_METHOD_CONNECT      string = "CONNECT"
	HTTP_METHOD_RPC_IN_DATA  string = "RPC_IN_DATA"
	HTTP_METHOD_RPC_OUT_DATA string = "RPC_OUT_DATA"
)

var NdpiHttpMethodNameMap = map[HttpMethod]string{
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

type HttpRequestVersion uint8

func (v *HttpRequestVersion) ToName() string {
	name, ok := NdpiHttpRequestVersionNameMap[*v]
	if !ok {
		name = ""
	}

	return name
}

const (
	NDPI_HTTP_Request_Version_1_0 HttpRequestVersion = 0
	NDPI_HTTP_Request_Version_1_1 HttpRequestVersion = 1
)

const (
	HTTP_Request_Version_1_0 string = "1.0"
	HTTP_Request_Version_1_1 string = "1.1"
)

var NdpiHttpRequestVersionNameMap = map[HttpRequestVersion]string{
	NDPI_HTTP_Request_Version_1_0: HTTP_Request_Version_1_0,
	NDPI_HTTP_Request_Version_1_1: HTTP_Request_Version_1_1,
}
