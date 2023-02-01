package types

type NdpiProtocolBreed uint16

func (b *NdpiProtocolBreed) ToName() string {
	name, ok := NdpiProtocolBreedNameMap[*b]
	if !ok {
		name = ""
	}

	return name
}

const (
	NDPI_PROTOCOL_SAFE                  NdpiProtocolBreed = 0
	NDPI_PROTOCOL_ACCEPTABLE            NdpiProtocolBreed = 1
	NDPI_PROTOCOL_FUN                   NdpiProtocolBreed = 2
	NDPI_PROTOCOL_UNSAFE                NdpiProtocolBreed = 3
	NDPI_PROTOCOL_POTENTIALLY_DANGEROUS NdpiProtocolBreed = 4
	NDPI_PROTOCOL_DANGEROUS             NdpiProtocolBreed = 5
	NDPI_PROTOCOL_TRACKER_ADS           NdpiProtocolBreed = 6
	NDPI_PROTOCOL_UNRATED               NdpiProtocolBreed = 7
)

const (
	PROTOCOL_SAFE                  string = "SAFE"
	PROTOCOL_ACCEPTABLE            string = "ACCEPTABLE"
	PROTOCOL_FUN                   string = "FUN"
	PROTOCOL_UNSAFE                string = "UNSAFE"
	PROTOCOL_POTENTIALLY_DANGEROUS string = "POTENTIALLY_DANGEROUS"
	PROTOCOL_DANGEROUS             string = "DANGEROUS"
	PROTOCOL_TRACKER_ADS           string = "TRACKER_ADS"
	PROTOCOL_UNRATED               string = "UNRATED"
)

var NdpiProtocolBreedNameMap = map[NdpiProtocolBreed]string{
	NDPI_PROTOCOL_SAFE:                  PROTOCOL_SAFE,
	NDPI_PROTOCOL_ACCEPTABLE:            PROTOCOL_ACCEPTABLE,
	NDPI_PROTOCOL_FUN:                   PROTOCOL_FUN,
	NDPI_PROTOCOL_UNSAFE:                PROTOCOL_UNSAFE,
	NDPI_PROTOCOL_POTENTIALLY_DANGEROUS: PROTOCOL_POTENTIALLY_DANGEROUS,
	NDPI_PROTOCOL_DANGEROUS:             PROTOCOL_DANGEROUS,
	NDPI_PROTOCOL_TRACKER_ADS:           PROTOCOL_TRACKER_ADS,
	NDPI_PROTOCOL_UNRATED:               PROTOCOL_UNRATED,
}
