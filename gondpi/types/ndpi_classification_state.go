package types

type NdpiClassificationState uint16

func (s *NdpiClassificationState) ToName() string {
	name, ok := NdpiClassificationStateNameMap[*s]
	if !ok {
		name = ""
	}

	return name
}

const (
	NDPI_STATE_INSPECTING NdpiClassificationState = 0
	NDPI_STATE_PARTIAL    NdpiClassificationState = 1
	NDPI_STATE_MONITORING NdpiClassificationState = 2
	NDPI_STATE_CLASSIFIED NdpiClassificationState = 3
)

const (
	STATE_INSPECTING string = "INSPECTING"
	STATE_PARTIAL    string = "PARTIAL"
	STATE_MONITORING string = "MONITORING"
	STATE_CLASSIFIED string = "CLASSIFIED"
)

var NdpiClassificationStateNameMap = map[NdpiClassificationState]string{
	NDPI_STATE_INSPECTING: STATE_INSPECTING,
	NDPI_STATE_PARTIAL:    STATE_PARTIAL,
	NDPI_STATE_MONITORING: STATE_MONITORING,
	NDPI_STATE_CLASSIFIED: STATE_CLASSIFIED,
}
