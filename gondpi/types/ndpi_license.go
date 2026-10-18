package types

// NdpiLicenseType selects which dissectors are loaded, according to their license
type NdpiLicenseType uint16

func (l *NdpiLicenseType) ToName() string {
	name, ok := NdpiLicenseTypeNameMap[*l]
	if !ok {
		name = ""
	}

	return name
}

const (
	NDPI_LICENSE_NOT_FOR_PROFIT_LGPL     NdpiLicenseType = 0
	NDPI_LICENSE_FOR_PROFIT_LGPL         NdpiLicenseType = 1
	NDPI_LICENSE_FOR_PROFIT_DUAL_LICENSE NdpiLicenseType = 2
)

const (
	LICENSE_NOT_FOR_PROFIT_LGPL     string = "NOT_FOR_PROFIT_LGPL"
	LICENSE_FOR_PROFIT_LGPL         string = "FOR_PROFIT_LGPL"
	LICENSE_FOR_PROFIT_DUAL_LICENSE string = "FOR_PROFIT_DUAL_LICENSE"
)

var NdpiLicenseTypeNameMap = map[NdpiLicenseType]string{
	NDPI_LICENSE_NOT_FOR_PROFIT_LGPL:     LICENSE_NOT_FOR_PROFIT_LGPL,
	NDPI_LICENSE_FOR_PROFIT_LGPL:         LICENSE_FOR_PROFIT_LGPL,
	NDPI_LICENSE_FOR_PROFIT_DUAL_LICENSE: LICENSE_FOR_PROFIT_DUAL_LICENSE,
}
