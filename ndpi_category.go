package gondpi

type Category string

const (
	CATEGORY_UNSPECIFIED        Category = "UNSPECIFIED"
	CATEGORY_MEDIA              Category = "MEDIA"
	CATEGORY_VPN                Category = "VPN"
	CATEGORY_MAIL               Category = "MAIL"
	CATEGORY_DATA_TRANSFER      Category = "DATA_TRANSFER"
	CATEGORY_WEB                Category = "WEB"
	CATEGORY_SOCIAL_NETWOR      Category = "SOCIAL_NETWOR"
	CATEGORY_DOWNLOAD_FT        Category = "DOWNLOAD_FT"
	CATEGORY_GAME               Category = "GAME"
	CATEGORY_CHAT               Category = "CHAT"
	CATEGORY_VOIP               Category = "VOIP"
	CATEGORY_DATABASE           Category = "DATABASE"
	CATEGORY_REMOTE_ACCESS      Category = "REMOTE_ACCESS"
	CATEGORY_CLOUD              Category = "CLOUD"
	CATEGORY_NETWORK            Category = "NETWORK"
	CATEGORY_COLLABORATIVE      Category = "COLLABORATIVE"
	CATEGORY_RPC                Category = "RPC"
	CATEGORY_STREAMING          Category = "STREAMING"
	CATEGORY_SYSTEM_OS          Category = "SYSTEM_OS"
	CATEGORY_SW_UPDATE          Category = "SW_UPDATE"
	CATEGORY_CUSTOM_1           Category = "CUSTOM_1"
	CATEGORY_CUSTOM_2           Category = "CUSTOM_2"
	CATEGORY_CUSTOM_3           Category = "CUSTOM_3"
	CATEGORY_CUSTOM_4           Category = "CUSTOM_4"
	CATEGORY_CUSTOM_5           Category = "CUSTOM_5"
	CATEGORY_MUSIC              Category = "MUSIC"
	CATEGORY_VIDEO              Category = "VIDEO"
	CATEGORY_SHOPPING           Category = "SHOPPING"
	CATEGORY_PRODUCTIVITY       Category = "PRODUCTIVITY"
	CATEGORY_FILE_SHARING       Category = "FILE_SHARING"
	CATEGORY_CONNECTIVITY_CHECK Category = "CONNECTIVITY_CHECK"
	CATEGORY_IOT_SCADA          Category = "IOT_SCADA"
	CATEGORY_VIRTUAL_ASSISTANT  Category = "VIRTUAL_ASSISTANT"
	CATEGORY_CYBERSECURITY      Category = "CYBERSECURITY"
	CATEGORY_MINING             Category = "MINING"
	CATEGORY_MALWARE            Category = "MALWARE"
	CATEGORY_ADVERTISEMENT      Category = "ADVERTISEMENT"
	CATEGORY_BANNED_SITE        Category = "BANNED_SITE"
	CATEGORY_SITE_UNAVAILABLE   Category = "SITE_UNAVAILABLE"
	CATEGORY_ALLOWED_SITE       Category = "ALLOWED_SITE"
	CATEGORY_ANTIMALWARE        Category = "ANTIMALWARE"
)

const (
	NDPI_PROTOCOL_CATEGORY_UNSPECIFIED        = 0
	NDPI_PROTOCOL_CATEGORY_MEDIA              = 1
	NDPI_PROTOCOL_CATEGORY_VPN                = 2
	NDPI_PROTOCOL_CATEGORY_MAIL               = 3
	NDPI_PROTOCOL_CATEGORY_DATA_TRANSFER      = 4
	NDPI_PROTOCOL_CATEGORY_WEB                = 5
	NDPI_PROTOCOL_CATEGORY_SOCIAL_NETWOR      = 6
	NDPI_PROTOCOL_CATEGORY_DOWNLOAD_FT        = 7
	NDPI_PROTOCOL_CATEGORY_GAME               = 8
	NDPI_PROTOCOL_CATEGORY_CHAT               = 9
	NDPI_PROTOCOL_CATEGORY_VOIP               = 10
	NDPI_PROTOCOL_CATEGORY_DATABASE           = 11
	NDPI_PROTOCOL_CATEGORY_REMOTE_ACCESS      = 12
	NDPI_PROTOCOL_CATEGORY_CLOUD              = 13
	NDPI_PROTOCOL_CATEGORY_NETWORK            = 14
	NDPI_PROTOCOL_CATEGORY_COLLABORATIVE      = 15
	NDPI_PROTOCOL_CATEGORY_RPC                = 16
	NDPI_PROTOCOL_CATEGORY_STREAMING          = 17
	NDPI_PROTOCOL_CATEGORY_SYSTEM_OS          = 18
	NDPI_PROTOCOL_CATEGORY_SW_UPDATE          = 19
	NDPI_PROTOCOL_CATEGORY_CUSTOM_1           = 20
	NDPI_PROTOCOL_CATEGORY_CUSTOM_2           = 21
	NDPI_PROTOCOL_CATEGORY_CUSTOM_3           = 22
	NDPI_PROTOCOL_CATEGORY_CUSTOM_4           = 23
	NDPI_PROTOCOL_CATEGORY_CUSTOM_5           = 24
	NDPI_PROTOCOL_CATEGORY_MUSIC              = 25
	NDPI_PROTOCOL_CATEGORY_VIDEO              = 26
	NDPI_PROTOCOL_CATEGORY_SHOPPING           = 27
	NDPI_PROTOCOL_CATEGORY_PRODUCTIVITY       = 28
	NDPI_PROTOCOL_CATEGORY_FILE_SHARING       = 29
	NDPI_PROTOCOL_CATEGORY_CONNECTIVITY_CHECK = 30
	NDPI_PROTOCOL_CATEGORY_IOT_SCADA          = 31
	NDPI_PROTOCOL_CATEGORY_VIRTUAL_ASSISTANT  = 32
	NDPI_PROTOCOL_CATEGORY_CYBERSECURITY      = 33
	CUSTOM_CATEGORY_MINING                    = 99
	CUSTOM_CATEGORY_MALWARE                   = 100
	CUSTOM_CATEGORY_ADVERTISEMENT             = 101
	CUSTOM_CATEGORY_BANNED_SITE               = 102
	CUSTOM_CATEGORY_SITE_UNAVAILABLE          = 103
	CUSTOM_CATEGORY_ALLOWED_SITE              = 104
	CUSTOM_CATEGORY_ANTIMALWARE               = 105
)

var NdpiCategoryIdMap = map[uint16]Category{
	NDPI_PROTOCOL_CATEGORY_UNSPECIFIED:        CATEGORY_UNSPECIFIED,
	NDPI_PROTOCOL_CATEGORY_MEDIA:              CATEGORY_MEDIA,
	NDPI_PROTOCOL_CATEGORY_VPN:                CATEGORY_VPN,
	NDPI_PROTOCOL_CATEGORY_MAIL:               CATEGORY_MAIL,
	NDPI_PROTOCOL_CATEGORY_DATA_TRANSFER:      CATEGORY_DATA_TRANSFER,
	NDPI_PROTOCOL_CATEGORY_WEB:                CATEGORY_WEB,
	NDPI_PROTOCOL_CATEGORY_SOCIAL_NETWOR:      CATEGORY_SOCIAL_NETWOR,
	NDPI_PROTOCOL_CATEGORY_DOWNLOAD_FT:        CATEGORY_DOWNLOAD_FT,
	NDPI_PROTOCOL_CATEGORY_GAME:               CATEGORY_GAME,
	NDPI_PROTOCOL_CATEGORY_CHAT:               CATEGORY_CHAT,
	NDPI_PROTOCOL_CATEGORY_VOIP:               CATEGORY_VOIP,
	NDPI_PROTOCOL_CATEGORY_DATABASE:           CATEGORY_DATABASE,
	NDPI_PROTOCOL_CATEGORY_REMOTE_ACCESS:      CATEGORY_REMOTE_ACCESS,
	NDPI_PROTOCOL_CATEGORY_CLOUD:              CATEGORY_CLOUD,
	NDPI_PROTOCOL_CATEGORY_NETWORK:            CATEGORY_NETWORK,
	NDPI_PROTOCOL_CATEGORY_COLLABORATIVE:      CATEGORY_COLLABORATIVE,
	NDPI_PROTOCOL_CATEGORY_RPC:                CATEGORY_RPC,
	NDPI_PROTOCOL_CATEGORY_STREAMING:          CATEGORY_STREAMING,
	NDPI_PROTOCOL_CATEGORY_SYSTEM_OS:          CATEGORY_SYSTEM_OS,
	NDPI_PROTOCOL_CATEGORY_SW_UPDATE:          CATEGORY_SW_UPDATE,
	NDPI_PROTOCOL_CATEGORY_CUSTOM_1:           CATEGORY_CUSTOM_1,
	NDPI_PROTOCOL_CATEGORY_CUSTOM_2:           CATEGORY_CUSTOM_2,
	NDPI_PROTOCOL_CATEGORY_CUSTOM_3:           CATEGORY_CUSTOM_3,
	NDPI_PROTOCOL_CATEGORY_CUSTOM_4:           CATEGORY_CUSTOM_4,
	NDPI_PROTOCOL_CATEGORY_CUSTOM_5:           CATEGORY_CUSTOM_5,
	NDPI_PROTOCOL_CATEGORY_MUSIC:              CATEGORY_MUSIC,
	NDPI_PROTOCOL_CATEGORY_VIDEO:              CATEGORY_VIDEO,
	NDPI_PROTOCOL_CATEGORY_SHOPPING:           CATEGORY_SHOPPING,
	NDPI_PROTOCOL_CATEGORY_PRODUCTIVITY:       CATEGORY_PRODUCTIVITY,
	NDPI_PROTOCOL_CATEGORY_FILE_SHARING:       CATEGORY_FILE_SHARING,
	NDPI_PROTOCOL_CATEGORY_CONNECTIVITY_CHECK: CATEGORY_CONNECTIVITY_CHECK,
	NDPI_PROTOCOL_CATEGORY_IOT_SCADA:          CATEGORY_IOT_SCADA,
	NDPI_PROTOCOL_CATEGORY_VIRTUAL_ASSISTANT:  CATEGORY_VIRTUAL_ASSISTANT,
	NDPI_PROTOCOL_CATEGORY_CYBERSECURITY:      CATEGORY_CYBERSECURITY,
	CUSTOM_CATEGORY_MINING:                    CATEGORY_MINING,
	CUSTOM_CATEGORY_MALWARE:                   CATEGORY_MALWARE,
	CUSTOM_CATEGORY_ADVERTISEMENT:             CATEGORY_ADVERTISEMENT,
	CUSTOM_CATEGORY_BANNED_SITE:               CATEGORY_BANNED_SITE,
	CUSTOM_CATEGORY_SITE_UNAVAILABLE:          CATEGORY_SITE_UNAVAILABLE,
	CUSTOM_CATEGORY_ALLOWED_SITE:              CATEGORY_ALLOWED_SITE,
	CUSTOM_CATEGORY_ANTIMALWARE:               CATEGORY_ANTIMALWARE,
}
