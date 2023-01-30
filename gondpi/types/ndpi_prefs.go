package types

const (
	NDPI_NO_PREFS                            uint32 = 0
	NDPI_DONT_LOAD_TOR_LIST                  uint32 = (1 << 0)
	NDPI_DONT_INIT_LIBGCRYPT                 uint32 = (1 << 1)
	NDPI_ENABLE_JA3_PLUS                     uint32 = (1 << 2)
	NDPI_DONT_LOAD_AZURE_LIST                uint32 = (1 << 3)
	NDPI_DONT_LOAD_WHATSAPP_LIST             uint32 = (1 << 4)
	NDPI_DONT_LOAD_AMAZON_AWS_LIST           uint32 = (1 << 5)
	NDPI_DONT_LOAD_ETHEREUM_LIST             uint32 = (1 << 6)
	NDPI_DONT_LOAD_ZOOM_LIST                 uint32 = (1 << 7)
	NDPI_DONT_LOAD_CLOUDFLARE_LIST           uint32 = (1 << 8)
	NDPI_DONT_LOAD_MICROSOFT_LIST            uint32 = (1 << 9)
	NDPI_DONT_LOAD_GOOGLE_LIST               uint32 = (1 << 10)
	NDPI_DONT_LOAD_GOOGLE_CLOUD_LIST         uint32 = (1 << 11)
	NDPI_DONT_LOAD_ASN_LISTS                 uint32 = (1 << 12)
	NDPI_DONT_LOAD_ICLOUD_PRIVATE_RELAY_LIST uint32 = (1 << 13)
	NDPI_DONT_INIT_RISK_PTREE                uint32 = (1 << 14)
	NDPI_DONT_LOAD_CACHEFLY_LIST                    = (1 << 15)
)
