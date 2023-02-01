package types

type NdpiConfidence uint16

func (c *NdpiConfidence) ToName() string {
	name, ok := NdpiConfidenceNameMap[*c]
	if !ok {
		name = ""
	}

	return name
}

const (
	NDPI_CONFIDENCE_UNKNOWN           NdpiConfidence = 0
	NDPI_CONFIDENCE_MATCH_BY_PORT     NdpiConfidence = 1
	NDPI_CONFIDENCE_MATCH_BY_IP       NdpiConfidence = 2
	NDPI_CONFIDENCE_DPI_PARTIAL       NdpiConfidence = 3
	NDPI_CONFIDENCE_DPI_PARTIAL_CACHE NdpiConfidence = 4
	NDPI_CONFIDENCE_DPI_CACHE         NdpiConfidence = 5
	NDPI_CONFIDENCE_DPI               NdpiConfidence = 6
)

const (
	CONFIDENCE_UNKNOWN           string = "UNKNOWN"
	CONFIDENCE_MATCH_BY_PORT     string = "MATCH_BY_PORT"
	CONFIDENCE_MATCH_BY_IP       string = "MATCH_BY_IP"
	CONFIDENCE_DPI_PARTIAL       string = "DPI_PARTIAL"
	CONFIDENCE_DPI_PARTIAL_CACHE string = "DPI_PARTIAL_CACHE"
	CONFIDENCE_DPI_CACHE         string = "DPI_CACHE"
	CONFIDENCE_DPI               string = "DPI"
)

var NdpiConfidenceNameMap = map[NdpiConfidence]string{
	NDPI_CONFIDENCE_UNKNOWN:           CONFIDENCE_UNKNOWN,
	NDPI_CONFIDENCE_MATCH_BY_PORT:     CONFIDENCE_MATCH_BY_PORT,
	NDPI_CONFIDENCE_MATCH_BY_IP:       CONFIDENCE_MATCH_BY_IP,
	NDPI_CONFIDENCE_DPI_PARTIAL:       CONFIDENCE_DPI_PARTIAL,
	NDPI_CONFIDENCE_DPI_PARTIAL_CACHE: CONFIDENCE_DPI_PARTIAL_CACHE,
	NDPI_CONFIDENCE_DPI_CACHE:         CONFIDENCE_DPI_CACHE,
	NDPI_CONFIDENCE_DPI:               CONFIDENCE_DPI,
}
