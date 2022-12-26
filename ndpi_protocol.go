package gondpi

type Protocol string

const (
	PROTO_UNKNOWN              Protocol = "UNKNOWN"
	PROTO_FTP_CONTROL          Protocol = "FTP_CONTROL"
	PROTO_MAIL_POP             Protocol = "MAIL_POP"
	PROTO_MAIL_SMTP            Protocol = "MAIL_SMTP"
	PROTO_MAIL_IMAP            Protocol = "MAIL_IMAP"
	PROTO_DNS                  Protocol = "DNS"
	PROTO_IPP                  Protocol = "IPP"
	PROTO_HTTP                 Protocol = "HTTP"
	PROTO_MDNS                 Protocol = "MDNS"
	PROTO_NTP                  Protocol = "NTP"
	PROTO_NETBIOS              Protocol = "NETBIOS"
	PROTO_NFS                  Protocol = "NFS"
	PROTO_SSDP                 Protocol = "SSDP"
	PROTO_BGP                  Protocol = "BGP"
	PROTO_SNMP                 Protocol = "SNMP"
	PROTO_XDMCP                Protocol = "XDMCP"
	PROTO_SMBV1                Protocol = "SMBV1"
	PROTO_SYSLOG               Protocol = "SYSLOG"
	PROTO_DHCP                 Protocol = "DHCP"
	PROTO_POSTGRES             Protocol = "POSTGRES"
	PROTO_MYSQL                Protocol = "MYSQL"
	PROTO_MS_OUTLOOK           Protocol = "MS_OUTLOOK"
	PROTO_DIRECT_DOWNLOAD_LINK Protocol = "DIRECT_DOWNLOAD_LINK"
	PROTO_MAIL_POPS            Protocol = "MAIL_POPS"
	PROTO_APPLEJUICE           Protocol = "APPLEJUICE"
	PROTO_DIRECTCONNECT        Protocol = "DIRECTCONNECT"
	PROTO_NTOP                 Protocol = "NTOP"
	PROTO_COAP                 Protocol = "COAP"
	PROTO_VMWARE               Protocol = "VMWARE"
	PROTO_MAIL_SMTPS           Protocol = "MAIL_SMTPS"
	PROTO_DTLS                 Protocol = "DTLS"
	PROTO_UBNTAC2              Protocol = "UBNTAC2"
	PROTO_KONTIKI              Protocol = "KONTIKI"
	PROTO_OPENFT               Protocol = "OPENFT"
	PROTO_FASTTRACK            Protocol = "FASTTRACK"
	PROTO_GNUTELLA             Protocol = "GNUTELLA"
	PROTO_EDONKEY              Protocol = "EDONKEY"
	PROTO_BITTORRENT           Protocol = "BITTORRENT"
	PROTO_SKYPE_TEAMS_CALL     Protocol = "SKYPE_TEAMS_CALL"
	PROTO_SIGNAL               Protocol = "SIGNAL"
	PROTO_MEMCACHED            Protocol = "MEMCACHED"
	PROTO_SMBV23               Protocol = "SMBV23"
	PROTO_MINING               Protocol = "MINING"
	PROTO_NEST_LOG_SINK        Protocol = "NEST_LOG_SINK"
	PROTO_MODBUS               Protocol = "MODBUS"
	PROTO_WHATSAPP_CALL        Protocol = "WHATSAPP_CALL"
	PROTO_DATASAVER            Protocol = "DATASAVER"
	PROTO_XBOX                 Protocol = "XBOX"
	PROTO_QQ                   Protocol = "QQ"
	PROTO_TIKTOK               Protocol = "TIKTOK"
	PROTO_RTSP                 Protocol = "RTSP"
	PROTO_MAIL_IMAPS           Protocol = "MAIL_IMAPS"
	PROTO_ICECAST              Protocol = "ICECAST"
	PROTO_CPHA                 Protocol = "CPHA"
	PROTO_PPSTREAM             Protocol = "PPSTREAM"
	PROTO_ZATTOO               Protocol = "ZATTOO"
	PROTO_SHOUTCAST            Protocol = "SHOUTCAST"
	PROTO_SOPCAST              Protocol = "SOPCAST"
	PROTO_DISCORD              Protocol = "DISCORD"
	PROTO_TVUPLAYER            Protocol = "TVUPLAYER"
	PROTO_MONGODB              Protocol = "MONGODB"
	PROTO_PLURALSIGHT          Protocol = "PLURALSIGHT"
	PROTO_THUNDER              Protocol = "THUNDER"
	PROTO_OCSP                 Protocol = "OCSP"
	PROTO_VXLAN                Protocol = "VXLAN"
	PROTO_IRC                  Protocol = "IRC"
	PROTO_AYIYA                Protocol = "AYIYA"
	PROTO_JABBER               Protocol = "JABBER"
	PROTO_NATS                 Protocol = "NATS"
	PROTO_AMONG_US             Protocol = "AMONG_US"
	PROTO_YAHOO                Protocol = "YAHOO"
	PROTO_DISNEYPLUS           Protocol = "DISNEYPLUS"
	PROTO_GOOGLE_PLUS          Protocol = "GOOGLE_PLUS"
	PROTO_IP_VRRP              Protocol = "IP_VRRP"
	PROTO_STEAM                Protocol = "STEAM"
	PROTO_HALFLIFE2            Protocol = "HALFLIFE2"
	PROTO_WORLDOFWARCRAFT      Protocol = "WORLDOFWARCRAFT"
	PROTO_TELNET               Protocol = "TELNET"
	PROTO_STUN                 Protocol = "STUN"
	PROTO_IPSEC                Protocol = "IPSEC"
	PROTO_IP_GRE               Protocol = "IP_GRE"
	PROTO_IP_ICMP              Protocol = "IP_ICMP"
	PROTO_IP_IGMP              Protocol = "IP_IGMP"
	PROTO_IP_EGP               Protocol = "IP_EGP"
	PROTO_IP_SCTP              Protocol = "IP_SCTP"
	PROTO_IP_OSPF              Protocol = "IP_OSPF"
	PROTO_IP_IP_IN_IP          Protocol = "IP_IP_IN_IP"
	PROTO_RTP                  Protocol = "RTP"
	PROTO_RDP                  Protocol = "RDP"
	PROTO_VNC                  Protocol = "VNC"
	PROTO_TUMBLR               Protocol = "TUMBLR"
	PROTO_TLS                  Protocol = "TLS"
	PROTO_SSH                  Protocol = "SSH"
	PROTO_USENET               Protocol = "USENET"
	PROTO_MGCP                 Protocol = "MGCP"
	PROTO_IAX                  Protocol = "IAX"
	PROTO_TFTP                 Protocol = "TFTP"
	PROTO_AFP                  Protocol = "AFP"
	PROTO_STEALTHNET           Protocol = "STEALTHNET"
	PROTO_AIMINI               Protocol = "AIMINI"
	PROTO_SIP                  Protocol = "SIP"
	PROTO_TRUPHONE             Protocol = "TRUPHONE"
	PROTO_IP_ICMPV6            Protocol = "IP_ICMPV6"
	PROTO_DHCPV6               Protocol = "DHCPV6"
	PROTO_ARMAGETRON           Protocol = "ARMAGETRON"
	PROTO_CROSSFIRE            Protocol = "CROSSFIRE"
	PROTO_DOFUS                Protocol = "DOFUS"
	PROTO_FIESTA               Protocol = "FIESTA"
	PROTO_FLORENSIA            Protocol = "FLORENSIA"
	PROTO_GUILDWARS            Protocol = "GUILDWARS"
	PROTO_AMAZON_ALEXA         Protocol = "AMAZON_ALEXA"
	PROTO_KERBEROS             Protocol = "KERBEROS"
	PROTO_LDAP                 Protocol = "LDAP"
	PROTO_MAPLESTORY           Protocol = "MAPLESTORY"
	PROTO_MSSQL_TDS            Protocol = "MSSQL_TDS"
	PROTO_PPTP                 Protocol = "PPTP"
	PROTO_WARCRAFT3            Protocol = "WARCRAFT3"
	PROTO_WORLD_OF_KUNG_FU     Protocol = "WORLD_OF_KUNG_FU"
	PROTO_SLACK                Protocol = "SLACK"
	PROTO_FACEBOOK             Protocol = "FACEBOOK"
	PROTO_TWITTER              Protocol = "TWITTER"
	PROTO_DROPBOX              Protocol = "DROPBOX"
	PROTO_GMAIL                Protocol = "GMAIL"
	PROTO_GOOGLE_MAPS          Protocol = "GOOGLE_MAPS"
	PROTO_YOUTUBE              Protocol = "YOUTUBE"
	PROTO_SKYPE_TEAMS          Protocol = "SKYPE_TEAMS"
	PROTO_GOOGLE               Protocol = "GOOGLE"
	PROTO_RPC                  Protocol = "RPC"
	PROTO_NETFLOW              Protocol = "NETFLOW"
	PROTO_SFLOW                Protocol = "SFLOW"
	PROTO_HTTP_CONNECT         Protocol = "HTTP_CONNECT"
	PROTO_HTTP_PROXY           Protocol = "HTTP_PROXY"
	PROTO_CITRIX               Protocol = "CITRIX"
	PROTO_NETFLIX              Protocol = "NETFLIX"
	PROTO_LASTFM               Protocol = "LASTFM"
	PROTO_WAZE                 Protocol = "WAZE"
	PROTO_YOUTUBE_UPLOAD       Protocol = "YOUTUBE_UPLOAD"
	PROTO_HULU                 Protocol = "HULU"
	PROTO_CHECKMK              Protocol = "CHECKMK"
	PROTO_AJP                  Protocol = "AJP"
	PROTO_APPLE                Protocol = "APPLE"
	PROTO_WEBEX                Protocol = "WEBEX"
	PROTO_WHATSAPP             Protocol = "WHATSAPP"
	PROTO_APPLE_ICLOUD         Protocol = "APPLE_ICLOUD"
	PROTO_VIBER                Protocol = "VIBER"
	PROTO_APPLE_ITUNES         Protocol = "APPLE_ITUNES"
	PROTO_RADIUS               Protocol = "RADIUS"
	PROTO_WINDOWS_UPDATE       Protocol = "WINDOWS_UPDATE"
	PROTO_TEAMVIEWER           Protocol = "TEAMVIEWER"
	PROTO_TUENTI               Protocol = "TUENTI"
	PROTO_LOTUS_NOTES          Protocol = "LOTUS_NOTES"
	PROTO_SAP                  Protocol = "SAP"
	PROTO_GTP                  Protocol = "GTP"
	PROTO_WSD                  Protocol = "WSD"
	PROTO_LLMNR                Protocol = "LLMNR"
	PROTO_TOCA_BOCA            Protocol = "TOCA_BOCA"
	PROTO_SPOTIFY              Protocol = "SPOTIFY"
	PROTO_MESSENGER            Protocol = "MESSENGER"
	PROTO_H323                 Protocol = "H323"
	PROTO_OPENVPN              Protocol = "OPENVPN"
	PROTO_NOE                  Protocol = "NOE"
	PROTO_CISCOVPN             Protocol = "CISCOVPN"
	PROTO_TEAMSPEAK            Protocol = "TEAMSPEAK"
	PROTO_TOR                  Protocol = "TOR"
	PROTO_SKINNY               Protocol = "SKINNY"
	PROTO_RTCP                 Protocol = "RTCP"
	PROTO_RSYNC                Protocol = "RSYNC"
	PROTO_ORACLE               Protocol = "ORACLE"
	PROTO_CORBA                Protocol = "CORBA"
	PROTO_UBUNTUONE            Protocol = "UBUNTUONE"
	PROTO_WHOIS_DAS            Protocol = "WHOIS_DAS"
	PROTO_SD_RTN               Protocol = "SD_RTN"
	PROTO_SOCKS                Protocol = "SOCKS"
	PROTO_NINTENDO             Protocol = "NINTENDO"
	PROTO_RTMP                 Protocol = "RTMP"
	PROTO_FTP_DATA             Protocol = "FTP_DATA"
	PROTO_WIKIPEDIA            Protocol = "WIKIPEDIA"
	PROTO_ZMQ                  Protocol = "ZMQ"
	PROTO_AMAZON               Protocol = "AMAZON"
	PROTO_EBAY                 Protocol = "EBAY"
	PROTO_CNN                  Protocol = "CNN"
	PROTO_MEGACO               Protocol = "MEGACO"
	PROTO_REDIS                Protocol = "REDIS"
	PROTO_PINTEREST            Protocol = "PINTEREST"
	PROTO_VHUA                 Protocol = "VHUA"
	PROTO_TELEGRAM             Protocol = "TELEGRAM"
	PROTO_VEVO                 Protocol = "VEVO"
	PROTO_PANDORA              Protocol = "PANDORA"
	PROTO_QUIC                 Protocol = "QUIC"
	PROTO_ZOOM                 Protocol = "ZOOM"
	PROTO_EAQ                  Protocol = "EAQ"
	PROTO_OOKLA                Protocol = "OOKLA"
	PROTO_AMQP                 Protocol = "AMQP"
	PROTO_KAKAOTALK            Protocol = "KAKAOTALK"
	PROTO_KAKAOTALK_VOICE      Protocol = "KAKAOTALK_VOICE"
	PROTO_TWITCH               Protocol = "TWITCH"
	PROTO_DOH_DOT              Protocol = "DOH_DOT"
	PROTO_WECHAT               Protocol = "WECHAT"
	PROTO_MPEGTS               Protocol = "MPEGTS"
	PROTO_SNAPCHAT             Protocol = "SNAPCHAT"
	PROTO_SINA                 Protocol = "SINA"
	PROTO_HANGOUT_DUO          Protocol = "HANGOUT_DUO"
	PROTO_IFLIX                Protocol = "IFLIX"
	PROTO_GITHUB               Protocol = "GITHUB"
	PROTO_BJNP                 Protocol = "BJNP"
	PROTO_REDDIT               Protocol = "REDDIT"
	PROTO_WIREGUARD            Protocol = "WIREGUARD"
	PROTO_SMPP                 Protocol = "SMPP"
	PROTO_DNSCRYPT             Protocol = "DNSCRYPT"
	PROTO_TINC                 Protocol = "TINC"
	PROTO_DEEZER               Protocol = "DEEZER"
	PROTO_INSTAGRAM            Protocol = "INSTAGRAM"
	PROTO_MICROSOFT            Protocol = "MICROSOFT"
	PROTO_STARCRAFT            Protocol = "STARCRAFT"
	PROTO_TEREDO               Protocol = "TEREDO"
	PROTO_HOTSPOT_SHIELD       Protocol = "HOTSPOT_SHIELD"
	PROTO_IMO                  Protocol = "IMO"
	PROTO_GOOGLE_DRIVE         Protocol = "GOOGLE_DRIVE"
	PROTO_OCS                  Protocol = "OCS"
	PROTO_MICROSOFT_365        Protocol = "MICROSOFT_365"
	PROTO_CLOUDFLARE           Protocol = "CLOUDFLARE"
	PROTO_MS_ONE_DRIVE         Protocol = "MS_ONE_DRIVE"
	PROTO_MQTT                 Protocol = "MQTT"
	PROTO_RX                   Protocol = "RX"
	PROTO_APPLESTORE           Protocol = "APPLESTORE"
	PROTO_OPENDNS              Protocol = "OPENDNS"
	PROTO_GIT                  Protocol = "GIT"
	PROTO_DRDA                 Protocol = "DRDA"
	PROTO_PLAYSTORE            Protocol = "PLAYSTORE"
	PROTO_SOMEIP               Protocol = "SOMEIP"
	PROTO_FIX                  Protocol = "FIX"
	PROTO_PLAYSTATION          Protocol = "PLAYSTATION"
	PROTO_PASTEBIN             Protocol = "PASTEBIN"
	PROTO_LINKEDIN             Protocol = "LINKEDIN"
	PROTO_SOUNDCLOUD           Protocol = "SOUNDCLOUD"
	PROTO_CSGO                 Protocol = "CSGO"
	PROTO_LISP                 Protocol = "LISP"
	PROTO_DIAMETER             Protocol = "DIAMETER"
	PROTO_APPLE_PUSH           Protocol = "APPLE_PUSH"
	PROTO_GOOGLE_SERVICES      Protocol = "GOOGLE_SERVICES"
	PROTO_AMAZON_VIDEO         Protocol = "AMAZON_VIDEO"
	PROTO_GOOGLE_DOCS          Protocol = "GOOGLE_DOCS"
	PROTO_WHATSAPP_FILES       Protocol = "WHATSAPP_FILES"
	PROTO_TARGUS_GETDATA       Protocol = "TARGUS_GETDATA"
	PROTO_DNP3                 Protocol = "DNP3"
	PROTO_IEC60870             Protocol = "IEC60870"
	PROTO_BLOOMBERG            Protocol = "BLOOMBERG"
	PROTO_CAPWAP               Protocol = "CAPWAP"
	PROTO_ZABBIX               Protocol = "ZABBIX"
	PROTO_S7COMM               Protocol = "S7COMM"
	PROTO_MSTEAMS              Protocol = "MSTEAMS"
	PROTO_WEBSOCKET            Protocol = "WEBSOCKET"
	PROTO_ANYDESK              Protocol = "ANYDESK"
	PROTO_SOAP                 Protocol = "SOAP"
	PROTO_APPLE_SIRI           Protocol = "APPLE_SIRI"
	PROTO_SNAPCHAT_CALL        Protocol = "SNAPCHAT_CALL"
	PROTO_HPVIRTGRP            Protocol = "HPVIRTGRP"
	PROTO_GENSHIN_IMPACT       Protocol = "GENSHIN_IMPACT"
	PROTO_ACTIVISION           Protocol = "ACTIVISION"
	PROTO_FORTICLIENT          Protocol = "FORTICLIENT"
	PROTO_Z3950                Protocol = "Z3950"
	PROTO_LIKEE                Protocol = "LIKEE"
	PROTO_GITLAB               Protocol = "GITLAB"
	PROTO_AVAST_SECUREDNS      Protocol = "AVAST_SECUREDNS"
	PROTO_CASSANDRA            Protocol = "CASSANDRA"
	PROTO_AMAZON_AWS           Protocol = "AMAZON_AWS"
	PROTO_SALESFORCE           Protocol = "SALESFORCE"
	PROTO_VIMEO                Protocol = "VIMEO"
	PROTO_FACEBOOK_VOIP        Protocol = "FACEBOOK_VOIP"
	PROTO_SIGNAL_VOIP          Protocol = "SIGNAL_VOIP"
	PROTO_FUZE                 Protocol = "FUZE"
	PROTO_GTP_U                Protocol = "GTP_U"
	PROTO_GTP_C                Protocol = "GTP_C"
	PROTO_GTP_PRIME            Protocol = "GTP_PRIME"
	PROTO_ALIBABA              Protocol = "ALIBABA"
	PROTO_CRASHLYSTICS         Protocol = "CRASHLYSTICS"
	PROTO_MICROSOFT_AZURE      Protocol = "MICROSOFT_AZURE"
	PROTO_ICLOUD_PRIVATE_RELAY Protocol = "ICLOUD_PRIVATE_RELAY"
	PROTO_ETHERNET_IP          Protocol = "ETHERNET_IP"
	PROTO_BADOO                Protocol = "BADOO"
	PROTO_ACCUWEATHER          Protocol = "ACCUWEATHER"
	PROTO_GOOGLE_CLASSROOM     Protocol = "GOOGLE_CLASSROOM"
	PROTO_HSRP                 Protocol = "HSRP"
	PROTO_CYBERSECURITY        Protocol = "CYBERSECURITY"
	PROTO_GOOGLE_CLOUD         Protocol = "GOOGLE_CLOUD"
	PROTO_TENCENT              Protocol = "TENCENT"
	PROTO_RAKNET               Protocol = "RAKNET"
	PROTO_XIAOMI               Protocol = "XIAOMI"
	PROTO_EDGECAST             Protocol = "EDGECAST"
	PROTO_CACHEFLY             Protocol = "CACHEFLY"
	PROTO_SOFTETHER            Protocol = "SOFTETHER"
	PROTO_MPEGDASH             Protocol = "MPEGDASH"
	PROTO_DAZN                 Protocol = "DAZN"
	PROTO_GOTO                 Protocol = "GOTO"
	PROTO_RSH                  Protocol = "RSH"
	PROTO_1KXUN                Protocol = "1KXUN"
	PROTO_IP_PGM               Protocol = "IP_PGM"
	PROTO_IP_PIM               Protocol = "IP_PIM"
	PROTO_COLLECTD             Protocol = "COLLECTD"
	PROTO_TUNNELBEAR           Protocol = "TUNNELBEAR"
	PROTO_CLOUDFLARE_WARP      Protocol = "CLOUDFLARE_WARP"
	PROTO_I3D                  Protocol = "I3D"
	PROTO_RIOTGAMES            Protocol = "RIOTGAMES"
	PROTO_PSIPHON              Protocol = "PSIPHON"
	PROTO_ULTRASURF            Protocol = "ULTRASURF"
)

const (
	NDPI_PROTOCOL_UNKNOWN              = 0
	NDPI_PROTOCOL_FTP_CONTROL          = 1
	NDPI_PROTOCOL_MAIL_POP             = 2
	NDPI_PROTOCOL_MAIL_SMTP            = 3
	NDPI_PROTOCOL_MAIL_IMAP            = 4
	NDPI_PROTOCOL_DNS                  = 5
	NDPI_PROTOCOL_IPP                  = 6
	NDPI_PROTOCOL_HTTP                 = 7
	NDPI_PROTOCOL_MDNS                 = 8
	NDPI_PROTOCOL_NTP                  = 9
	NDPI_PROTOCOL_NETBIOS              = 10
	NDPI_PROTOCOL_NFS                  = 11
	NDPI_PROTOCOL_SSDP                 = 12
	NDPI_PROTOCOL_BGP                  = 13
	NDPI_PROTOCOL_SNMP                 = 14
	NDPI_PROTOCOL_XDMCP                = 15
	NDPI_PROTOCOL_SMBV1                = 16
	NDPI_PROTOCOL_SYSLOG               = 17
	NDPI_PROTOCOL_DHCP                 = 18
	NDPI_PROTOCOL_POSTGRES             = 19
	NDPI_PROTOCOL_MYSQL                = 20
	NDPI_PROTOCOL_MS_OUTLOOK           = 21
	NDPI_PROTOCOL_DIRECT_DOWNLOAD_LINK = 22
	NDPI_PROTOCOL_MAIL_POPS            = 23
	NDPI_PROTOCOL_APPLEJUICE           = 24
	NDPI_PROTOCOL_DIRECTCONNECT        = 25
	NDPI_PROTOCOL_NTOP                 = 26
	NDPI_PROTOCOL_COAP                 = 27
	NDPI_PROTOCOL_VMWARE               = 28
	NDPI_PROTOCOL_MAIL_SMTPS           = 29
	NDPI_PROTOCOL_DTLS                 = 30
	NDPI_PROTOCOL_UBNTAC2              = 31
	NDPI_PROTOCOL_KONTIKI              = 32
	NDPI_PROTOCOL_OPENFT               = 33
	NDPI_PROTOCOL_FASTTRACK            = 34
	NDPI_PROTOCOL_GNUTELLA             = 35
	NDPI_PROTOCOL_EDONKEY              = 36
	NDPI_PROTOCOL_BITTORRENT           = 37
	NDPI_PROTOCOL_SKYPE_TEAMS_CALL     = 38
	NDPI_PROTOCOL_SIGNAL               = 39
	NDPI_PROTOCOL_MEMCACHED            = 40
	NDPI_PROTOCOL_SMBV23               = 41
	NDPI_PROTOCOL_MINING               = 42
	NDPI_PROTOCOL_NEST_LOG_SINK        = 43
	NDPI_PROTOCOL_MODBUS               = 44
	NDPI_PROTOCOL_WHATSAPP_CALL        = 45
	NDPI_PROTOCOL_DATASAVER            = 46
	NDPI_PROTOCOL_XBOX                 = 47
	NDPI_PROTOCOL_QQ                   = 48
	NDPI_PROTOCOL_TIKTOK               = 49
	NDPI_PROTOCOL_RTSP                 = 50
	NDPI_PROTOCOL_MAIL_IMAPS           = 51
	NDPI_PROTOCOL_ICECAST              = 52
	NDPI_PROTOCOL_CPHA                 = 53
	NDPI_PROTOCOL_PPSTREAM             = 54
	NDPI_PROTOCOL_ZATTOO               = 55
	NDPI_PROTOCOL_SHOUTCAST            = 56
	NDPI_PROTOCOL_SOPCAST              = 57
	NDPI_PROTOCOL_DISCORD              = 58
	NDPI_PROTOCOL_TVUPLAYER            = 59
	NDPI_PROTOCOL_MONGODB              = 60
	NDPI_PROTOCOL_PLURALSIGHT          = 61
	NDPI_PROTOCOL_THUNDER              = 62
	NDPI_PROTOCOL_OCSP                 = 63
	NDPI_PROTOCOL_VXLAN                = 64
	NDPI_PROTOCOL_IRC                  = 65
	NDPI_PROTOCOL_AYIYA                = 66
	NDPI_PROTOCOL_JABBER               = 67
	NDPI_PROTOCOL_NATS                 = 68
	NDPI_PROTOCOL_AMONG_US             = 69
	NDPI_PROTOCOL_YAHOO                = 70
	NDPI_PROTOCOL_DISNEYPLUS           = 71
	NDPI_PROTOCOL_GOOGLE_PLUS          = 72
	NDPI_PROTOCOL_IP_VRRP              = 73
	NDPI_PROTOCOL_STEAM                = 74
	NDPI_PROTOCOL_HALFLIFE2            = 75
	NDPI_PROTOCOL_WORLDOFWARCRAFT      = 76
	NDPI_PROTOCOL_TELNET               = 77
	NDPI_PROTOCOL_STUN                 = 78
	NDPI_PROTOCOL_IPSEC                = 79
	NDPI_PROTOCOL_IP_GRE               = 80
	NDPI_PROTOCOL_IP_ICMP              = 81
	NDPI_PROTOCOL_IP_IGMP              = 82
	NDPI_PROTOCOL_IP_EGP               = 83
	NDPI_PROTOCOL_IP_SCTP              = 84
	NDPI_PROTOCOL_IP_OSPF              = 85
	NDPI_PROTOCOL_IP_IP_IN_IP          = 86
	NDPI_PROTOCOL_RTP                  = 87
	NDPI_PROTOCOL_RDP                  = 88
	NDPI_PROTOCOL_VNC                  = 89
	NDPI_PROTOCOL_TUMBLR               = 90
	NDPI_PROTOCOL_TLS                  = 91
	NDPI_PROTOCOL_SSH                  = 92
	NDPI_PROTOCOL_USENET               = 93
	NDPI_PROTOCOL_MGCP                 = 94
	NDPI_PROTOCOL_IAX                  = 95
	NDPI_PROTOCOL_TFTP                 = 96
	NDPI_PROTOCOL_AFP                  = 97
	NDPI_PROTOCOL_STEALTHNET           = 98
	NDPI_PROTOCOL_AIMINI               = 99
	NDPI_PROTOCOL_SIP                  = 100
	NDPI_PROTOCOL_TRUPHONE             = 101
	NDPI_PROTOCOL_IP_ICMPV6            = 102
	NDPI_PROTOCOL_DHCPV6               = 103
	NDPI_PROTOCOL_ARMAGETRON           = 104
	NDPI_PROTOCOL_CROSSFIRE            = 105
	NDPI_PROTOCOL_DOFUS                = 106
	NDPI_PROTOCOL_FIESTA               = 107
	NDPI_PROTOCOL_FLORENSIA            = 108
	NDPI_PROTOCOL_GUILDWARS            = 109
	NDPI_PROTOCOL_AMAZON_ALEXA         = 110
	NDPI_PROTOCOL_KERBEROS             = 111
	NDPI_PROTOCOL_LDAP                 = 112
	NDPI_PROTOCOL_MAPLESTORY           = 113
	NDPI_PROTOCOL_MSSQL_TDS            = 114
	NDPI_PROTOCOL_PPTP                 = 115
	NDPI_PROTOCOL_WARCRAFT3            = 116
	NDPI_PROTOCOL_WORLD_OF_KUNG_FU     = 117
	NDPI_PROTOCOL_SLACK                = 118
	NDPI_PROTOCOL_FACEBOOK             = 119
	NDPI_PROTOCOL_TWITTER              = 120
	NDPI_PROTOCOL_DROPBOX              = 121
	NDPI_PROTOCOL_GMAIL                = 122
	NDPI_PROTOCOL_GOOGLE_MAPS          = 123
	NDPI_PROTOCOL_YOUTUBE              = 124
	NDPI_PROTOCOL_SKYPE_TEAMS          = 125
	NDPI_PROTOCOL_GOOGLE               = 126
	NDPI_PROTOCOL_RPC                  = 127
	NDPI_PROTOCOL_NETFLOW              = 128
	NDPI_PROTOCOL_SFLOW                = 129
	NDPI_PROTOCOL_HTTP_CONNECT         = 130
	NDPI_PROTOCOL_HTTP_PROXY           = 131
	NDPI_PROTOCOL_CITRIX               = 132
	NDPI_PROTOCOL_NETFLIX              = 133
	NDPI_PROTOCOL_LASTFM               = 134
	NDPI_PROTOCOL_WAZE                 = 135
	NDPI_PROTOCOL_YOUTUBE_UPLOAD       = 136
	NDPI_PROTOCOL_HULU                 = 137
	NDPI_PROTOCOL_CHECKMK              = 138
	NDPI_PROTOCOL_AJP                  = 139
	NDPI_PROTOCOL_APPLE                = 140
	NDPI_PROTOCOL_WEBEX                = 141
	NDPI_PROTOCOL_WHATSAPP             = 142
	NDPI_PROTOCOL_APPLE_ICLOUD         = 143
	NDPI_PROTOCOL_VIBER                = 144
	NDPI_PROTOCOL_APPLE_ITUNES         = 145
	NDPI_PROTOCOL_RADIUS               = 146
	NDPI_PROTOCOL_WINDOWS_UPDATE       = 147
	NDPI_PROTOCOL_TEAMVIEWER           = 148
	NDPI_PROTOCOL_TUENTI               = 149
	NDPI_PROTOCOL_LOTUS_NOTES          = 150
	NDPI_PROTOCOL_SAP                  = 151
	NDPI_PROTOCOL_GTP                  = 152
	NDPI_PROTOCOL_WSD                  = 153
	NDPI_PROTOCOL_LLMNR                = 154
	NDPI_PROTOCOL_TOCA_BOCA            = 155
	NDPI_PROTOCOL_SPOTIFY              = 156
	NDPI_PROTOCOL_MESSENGER            = 157
	NDPI_PROTOCOL_H323                 = 158
	NDPI_PROTOCOL_OPENVPN              = 159
	NDPI_PROTOCOL_NOE                  = 160
	NDPI_PROTOCOL_CISCOVPN             = 161
	NDPI_PROTOCOL_TEAMSPEAK            = 162
	NDPI_PROTOCOL_TOR                  = 163
	NDPI_PROTOCOL_SKINNY               = 164
	NDPI_PROTOCOL_RTCP                 = 165
	NDPI_PROTOCOL_RSYNC                = 166
	NDPI_PROTOCOL_ORACLE               = 167
	NDPI_PROTOCOL_CORBA                = 168
	NDPI_PROTOCOL_UBUNTUONE            = 169
	NDPI_PROTOCOL_WHOIS_DAS            = 170
	NDPI_PROTOCOL_SD_RTN               = 171
	NDPI_PROTOCOL_SOCKS                = 172
	NDPI_PROTOCOL_NINTENDO             = 173
	NDPI_PROTOCOL_RTMP                 = 174
	NDPI_PROTOCOL_FTP_DATA             = 175
	NDPI_PROTOCOL_WIKIPEDIA            = 176
	NDPI_PROTOCOL_ZMQ                  = 177
	NDPI_PROTOCOL_AMAZON               = 178
	NDPI_PROTOCOL_EBAY                 = 179
	NDPI_PROTOCOL_CNN                  = 180
	NDPI_PROTOCOL_MEGACO               = 181
	NDPI_PROTOCOL_REDIS                = 182
	NDPI_PROTOCOL_PINTEREST            = 183
	NDPI_PROTOCOL_VHUA                 = 184
	NDPI_PROTOCOL_TELEGRAM             = 185
	NDPI_PROTOCOL_VEVO                 = 186
	NDPI_PROTOCOL_PANDORA              = 187
	NDPI_PROTOCOL_QUIC                 = 188
	NDPI_PROTOCOL_ZOOM                 = 189
	NDPI_PROTOCOL_EAQ                  = 190
	NDPI_PROTOCOL_OOKLA                = 191
	NDPI_PROTOCOL_AMQP                 = 192
	NDPI_PROTOCOL_KAKAOTALK            = 193
	NDPI_PROTOCOL_KAKAOTALK_VOICE      = 194
	NDPI_PROTOCOL_TWITCH               = 195
	NDPI_PROTOCOL_DOH_DOT              = 196
	NDPI_PROTOCOL_WECHAT               = 197
	NDPI_PROTOCOL_MPEGTS               = 198
	NDPI_PROTOCOL_SNAPCHAT             = 199
	NDPI_PROTOCOL_SINA                 = 200
	NDPI_PROTOCOL_HANGOUT_DUO          = 201
	NDPI_PROTOCOL_IFLIX                = 202
	NDPI_PROTOCOL_GITHUB               = 203
	NDPI_PROTOCOL_BJNP                 = 204
	NDPI_PROTOCOL_REDDIT               = 205
	NDPI_PROTOCOL_WIREGUARD            = 206
	NDPI_PROTOCOL_SMPP                 = 207
	NDPI_PROTOCOL_DNSCRYPT             = 208
	NDPI_PROTOCOL_TINC                 = 209
	NDPI_PROTOCOL_DEEZER               = 210
	NDPI_PROTOCOL_INSTAGRAM            = 211
	NDPI_PROTOCOL_MICROSOFT            = 212
	NDPI_PROTOCOL_STARCRAFT            = 213
	NDPI_PROTOCOL_TEREDO               = 214
	NDPI_PROTOCOL_HOTSPOT_SHIELD       = 215
	NDPI_PROTOCOL_IMO                  = 216
	NDPI_PROTOCOL_GOOGLE_DRIVE         = 217
	NDPI_PROTOCOL_OCS                  = 218
	NDPI_PROTOCOL_MICROSOFT_365        = 219
	NDPI_PROTOCOL_CLOUDFLARE           = 220
	NDPI_PROTOCOL_MS_ONE_DRIVE         = 221
	NDPI_PROTOCOL_MQTT                 = 222
	NDPI_PROTOCOL_RX                   = 223
	NDPI_PROTOCOL_APPLESTORE           = 224
	NDPI_PROTOCOL_OPENDNS              = 225
	NDPI_PROTOCOL_GIT                  = 226
	NDPI_PROTOCOL_DRDA                 = 227
	NDPI_PROTOCOL_PLAYSTORE            = 228
	NDPI_PROTOCOL_SOMEIP               = 229
	NDPI_PROTOCOL_FIX                  = 230
	NDPI_PROTOCOL_PLAYSTATION          = 231
	NDPI_PROTOCOL_PASTEBIN             = 232
	NDPI_PROTOCOL_LINKEDIN             = 233
	NDPI_PROTOCOL_SOUNDCLOUD           = 234
	NDPI_PROTOCOL_CSGO                 = 235
	NDPI_PROTOCOL_LISP                 = 236
	NDPI_PROTOCOL_DIAMETER             = 237
	NDPI_PROTOCOL_APPLE_PUSH           = 238
	NDPI_PROTOCOL_GOOGLE_SERVICES      = 239
	NDPI_PROTOCOL_AMAZON_VIDEO         = 240
	NDPI_PROTOCOL_GOOGLE_DOCS          = 241
	NDPI_PROTOCOL_WHATSAPP_FILES       = 242
	NDPI_PROTOCOL_TARGUS_GETDATA       = 243
	NDPI_PROTOCOL_DNP3                 = 244
	NDPI_PROTOCOL_IEC60870             = 245
	NDPI_PROTOCOL_BLOOMBERG            = 246
	NDPI_PROTOCOL_CAPWAP               = 247
	NDPI_PROTOCOL_ZABBIX               = 248
	NDPI_PROTOCOL_S7COMM               = 249
	NDPI_PROTOCOL_MSTEAMS              = 250
	NDPI_PROTOCOL_WEBSOCKET            = 251
	NDPI_PROTOCOL_ANYDESK              = 252
	NDPI_PROTOCOL_SOAP                 = 253
	NDPI_PROTOCOL_APPLE_SIRI           = 254
	NDPI_PROTOCOL_SNAPCHAT_CALL        = 255
	NDPI_PROTOCOL_HPVIRTGRP            = 256
	NDPI_PROTOCOL_GENSHIN_IMPACT       = 257
	NDPI_PROTOCOL_ACTIVISION           = 258
	NDPI_PROTOCOL_FORTICLIENT          = 259
	NDPI_PROTOCOL_Z3950                = 260
	NDPI_PROTOCOL_LIKEE                = 261
	NDPI_PROTOCOL_GITLAB               = 262
	NDPI_PROTOCOL_AVAST_SECUREDNS      = 263
	NDPI_PROTOCOL_CASSANDRA            = 264
	NDPI_PROTOCOL_AMAZON_AWS           = 265
	NDPI_PROTOCOL_SALESFORCE           = 266
	NDPI_PROTOCOL_VIMEO                = 267
	NDPI_PROTOCOL_FACEBOOK_VOIP        = 268
	NDPI_PROTOCOL_SIGNAL_VOIP          = 269
	NDPI_PROTOCOL_FUZE                 = 270
	NDPI_PROTOCOL_GTP_U                = 271
	NDPI_PROTOCOL_GTP_C                = 272
	NDPI_PROTOCOL_GTP_PRIME            = 273
	NDPI_PROTOCOL_ALIBABA              = 274
	NDPI_PROTOCOL_CRASHLYSTICS         = 275
	NDPI_PROTOCOL_MICROSOFT_AZURE      = 276
	NDPI_PROTOCOL_ICLOUD_PRIVATE_RELAY = 277
	NDPI_PROTOCOL_ETHERNET_IP          = 278
	NDPI_PROTOCOL_BADOO                = 279
	NDPI_PROTOCOL_ACCUWEATHER          = 280
	NDPI_PROTOCOL_GOOGLE_CLASSROOM     = 281
	NDPI_PROTOCOL_HSRP                 = 282
	NDPI_PROTOCOL_CYBERSECURITY        = 283
	NDPI_PROTOCOL_GOOGLE_CLOUD         = 284
	NDPI_PROTOCOL_TENCENT              = 285
	NDPI_PROTOCOL_RAKNET               = 286
	NDPI_PROTOCOL_XIAOMI               = 287
	NDPI_PROTOCOL_EDGECAST             = 288
	NDPI_PROTOCOL_CACHEFLY             = 289
	NDPI_PROTOCOL_SOFTETHER            = 290
	NDPI_PROTOCOL_MPEGDASH             = 291
	NDPI_PROTOCOL_DAZN                 = 292
	NDPI_PROTOCOL_GOTO                 = 293
	NDPI_PROTOCOL_RSH                  = 294
	NDPI_PROTOCOL_1KXUN                = 295
	NDPI_PROTOCOL_IP_PGM               = 296
	NDPI_PROTOCOL_IP_PIM               = 297
	NDPI_PROTOCOL_COLLECTD             = 298
	NDPI_PROTOCOL_TUNNELBEAR           = 299
	NDPI_PROTOCOL_CLOUDFLARE_WARP      = 300
	NDPI_PROTOCOL_I3D                  = 301
	NDPI_PROTOCOL_RIOTGAMES            = 302
	NDPI_PROTOCOL_PSIPHON              = 303
	NDPI_PROTOCOL_ULTRASURF            = 304
)

var NdpiProtocolIdMap = map[uint16]Protocol{
	NDPI_PROTOCOL_UNKNOWN:              PROTO_UNKNOWN,
	NDPI_PROTOCOL_FTP_CONTROL:          PROTO_FTP_CONTROL,
	NDPI_PROTOCOL_MAIL_POP:             PROTO_MAIL_POP,
	NDPI_PROTOCOL_MAIL_SMTP:            PROTO_MAIL_SMTP,
	NDPI_PROTOCOL_MAIL_IMAP:            PROTO_MAIL_IMAP,
	NDPI_PROTOCOL_DNS:                  PROTO_DNS,
	NDPI_PROTOCOL_IPP:                  PROTO_IPP,
	NDPI_PROTOCOL_HTTP:                 PROTO_HTTP,
	NDPI_PROTOCOL_MDNS:                 PROTO_MDNS,
	NDPI_PROTOCOL_NTP:                  PROTO_NTP,
	NDPI_PROTOCOL_NETBIOS:              PROTO_NETBIOS,
	NDPI_PROTOCOL_NFS:                  PROTO_NFS,
	NDPI_PROTOCOL_SSDP:                 PROTO_SSDP,
	NDPI_PROTOCOL_BGP:                  PROTO_BGP,
	NDPI_PROTOCOL_SNMP:                 PROTO_SNMP,
	NDPI_PROTOCOL_XDMCP:                PROTO_XDMCP,
	NDPI_PROTOCOL_SMBV1:                PROTO_SMBV1,
	NDPI_PROTOCOL_SYSLOG:               PROTO_SYSLOG,
	NDPI_PROTOCOL_DHCP:                 PROTO_DHCP,
	NDPI_PROTOCOL_POSTGRES:             PROTO_POSTGRES,
	NDPI_PROTOCOL_MYSQL:                PROTO_MYSQL,
	NDPI_PROTOCOL_MS_OUTLOOK:           PROTO_MS_OUTLOOK,
	NDPI_PROTOCOL_DIRECT_DOWNLOAD_LINK: PROTO_DIRECT_DOWNLOAD_LINK,
	NDPI_PROTOCOL_MAIL_POPS:            PROTO_MAIL_POPS,
	NDPI_PROTOCOL_APPLEJUICE:           PROTO_APPLEJUICE,
	NDPI_PROTOCOL_DIRECTCONNECT:        PROTO_DIRECTCONNECT,
	NDPI_PROTOCOL_NTOP:                 PROTO_NTOP,
	NDPI_PROTOCOL_COAP:                 PROTO_COAP,
	NDPI_PROTOCOL_VMWARE:               PROTO_VMWARE,
	NDPI_PROTOCOL_MAIL_SMTPS:           PROTO_MAIL_SMTPS,
	NDPI_PROTOCOL_DTLS:                 PROTO_DTLS,
	NDPI_PROTOCOL_UBNTAC2:              PROTO_UBNTAC2,
	NDPI_PROTOCOL_KONTIKI:              PROTO_KONTIKI,
	NDPI_PROTOCOL_OPENFT:               PROTO_OPENFT,
	NDPI_PROTOCOL_FASTTRACK:            PROTO_FASTTRACK,
	NDPI_PROTOCOL_GNUTELLA:             PROTO_GNUTELLA,
	NDPI_PROTOCOL_EDONKEY:              PROTO_EDONKEY,
	NDPI_PROTOCOL_BITTORRENT:           PROTO_BITTORRENT,
	NDPI_PROTOCOL_SKYPE_TEAMS_CALL:     PROTO_SKYPE_TEAMS_CALL,
	NDPI_PROTOCOL_SIGNAL:               PROTO_SIGNAL,
	NDPI_PROTOCOL_MEMCACHED:            PROTO_MEMCACHED,
	NDPI_PROTOCOL_SMBV23:               PROTO_SMBV23,
	NDPI_PROTOCOL_MINING:               PROTO_MINING,
	NDPI_PROTOCOL_NEST_LOG_SINK:        PROTO_NEST_LOG_SINK,
	NDPI_PROTOCOL_MODBUS:               PROTO_MODBUS,
	NDPI_PROTOCOL_WHATSAPP_CALL:        PROTO_WHATSAPP_CALL,
	NDPI_PROTOCOL_DATASAVER:            PROTO_DATASAVER,
	NDPI_PROTOCOL_XBOX:                 PROTO_XBOX,
	NDPI_PROTOCOL_QQ:                   PROTO_QQ,
	NDPI_PROTOCOL_TIKTOK:               PROTO_TIKTOK,
	NDPI_PROTOCOL_RTSP:                 PROTO_RTSP,
	NDPI_PROTOCOL_MAIL_IMAPS:           PROTO_MAIL_IMAPS,
	NDPI_PROTOCOL_ICECAST:              PROTO_ICECAST,
	NDPI_PROTOCOL_CPHA:                 PROTO_CPHA,
	NDPI_PROTOCOL_PPSTREAM:             PROTO_PPSTREAM,
	NDPI_PROTOCOL_ZATTOO:               PROTO_ZATTOO,
	NDPI_PROTOCOL_SHOUTCAST:            PROTO_SHOUTCAST,
	NDPI_PROTOCOL_SOPCAST:              PROTO_SOPCAST,
	NDPI_PROTOCOL_DISCORD:              PROTO_DISCORD,
	NDPI_PROTOCOL_TVUPLAYER:            PROTO_TVUPLAYER,
	NDPI_PROTOCOL_MONGODB:              PROTO_MONGODB,
	NDPI_PROTOCOL_PLURALSIGHT:          PROTO_PLURALSIGHT,
	NDPI_PROTOCOL_THUNDER:              PROTO_THUNDER,
	NDPI_PROTOCOL_OCSP:                 PROTO_OCSP,
	NDPI_PROTOCOL_VXLAN:                PROTO_VXLAN,
	NDPI_PROTOCOL_IRC:                  PROTO_IRC,
	NDPI_PROTOCOL_AYIYA:                PROTO_AYIYA,
	NDPI_PROTOCOL_JABBER:               PROTO_JABBER,
	NDPI_PROTOCOL_NATS:                 PROTO_NATS,
	NDPI_PROTOCOL_AMONG_US:             PROTO_AMONG_US,
	NDPI_PROTOCOL_YAHOO:                PROTO_YAHOO,
	NDPI_PROTOCOL_DISNEYPLUS:           PROTO_DISNEYPLUS,
	NDPI_PROTOCOL_GOOGLE_PLUS:          PROTO_GOOGLE_PLUS,
	NDPI_PROTOCOL_IP_VRRP:              PROTO_IP_VRRP,
	NDPI_PROTOCOL_STEAM:                PROTO_STEAM,
	NDPI_PROTOCOL_HALFLIFE2:            PROTO_HALFLIFE2,
	NDPI_PROTOCOL_WORLDOFWARCRAFT:      PROTO_WORLDOFWARCRAFT,
	NDPI_PROTOCOL_TELNET:               PROTO_TELNET,
	NDPI_PROTOCOL_STUN:                 PROTO_STUN,
	NDPI_PROTOCOL_IPSEC:                PROTO_IPSEC,
	NDPI_PROTOCOL_IP_GRE:               PROTO_IP_GRE,
	NDPI_PROTOCOL_IP_ICMP:              PROTO_IP_ICMP,
	NDPI_PROTOCOL_IP_IGMP:              PROTO_IP_IGMP,
	NDPI_PROTOCOL_IP_EGP:               PROTO_IP_EGP,
	NDPI_PROTOCOL_IP_SCTP:              PROTO_IP_SCTP,
	NDPI_PROTOCOL_IP_OSPF:              PROTO_IP_OSPF,
	NDPI_PROTOCOL_IP_IP_IN_IP:          PROTO_IP_IP_IN_IP,
	NDPI_PROTOCOL_RTP:                  PROTO_RTP,
	NDPI_PROTOCOL_RDP:                  PROTO_RDP,
	NDPI_PROTOCOL_VNC:                  PROTO_VNC,
	NDPI_PROTOCOL_TUMBLR:               PROTO_TUMBLR,
	NDPI_PROTOCOL_TLS:                  PROTO_TLS,
	NDPI_PROTOCOL_SSH:                  PROTO_SSH,
	NDPI_PROTOCOL_USENET:               PROTO_USENET,
	NDPI_PROTOCOL_MGCP:                 PROTO_MGCP,
	NDPI_PROTOCOL_IAX:                  PROTO_IAX,
	NDPI_PROTOCOL_TFTP:                 PROTO_TFTP,
	NDPI_PROTOCOL_AFP:                  PROTO_AFP,
	NDPI_PROTOCOL_STEALTHNET:           PROTO_STEALTHNET,
	NDPI_PROTOCOL_AIMINI:               PROTO_AIMINI,
	NDPI_PROTOCOL_SIP:                  PROTO_SIP,
	NDPI_PROTOCOL_TRUPHONE:             PROTO_TRUPHONE,
	NDPI_PROTOCOL_IP_ICMPV6:            PROTO_IP_ICMPV6,
	NDPI_PROTOCOL_DHCPV6:               PROTO_DHCPV6,
	NDPI_PROTOCOL_ARMAGETRON:           PROTO_ARMAGETRON,
	NDPI_PROTOCOL_CROSSFIRE:            PROTO_CROSSFIRE,
	NDPI_PROTOCOL_DOFUS:                PROTO_DOFUS,
	NDPI_PROTOCOL_FIESTA:               PROTO_FIESTA,
	NDPI_PROTOCOL_FLORENSIA:            PROTO_FLORENSIA,
	NDPI_PROTOCOL_GUILDWARS:            PROTO_GUILDWARS,
	NDPI_PROTOCOL_AMAZON_ALEXA:         PROTO_AMAZON_ALEXA,
	NDPI_PROTOCOL_KERBEROS:             PROTO_KERBEROS,
	NDPI_PROTOCOL_LDAP:                 PROTO_LDAP,
	NDPI_PROTOCOL_MAPLESTORY:           PROTO_MAPLESTORY,
	NDPI_PROTOCOL_MSSQL_TDS:            PROTO_MSSQL_TDS,
	NDPI_PROTOCOL_PPTP:                 PROTO_PPTP,
	NDPI_PROTOCOL_WARCRAFT3:            PROTO_WARCRAFT3,
	NDPI_PROTOCOL_WORLD_OF_KUNG_FU:     PROTO_WORLD_OF_KUNG_FU,
	NDPI_PROTOCOL_SLACK:                PROTO_SLACK,
	NDPI_PROTOCOL_FACEBOOK:             PROTO_FACEBOOK,
	NDPI_PROTOCOL_TWITTER:              PROTO_TWITTER,
	NDPI_PROTOCOL_DROPBOX:              PROTO_DROPBOX,
	NDPI_PROTOCOL_GMAIL:                PROTO_GMAIL,
	NDPI_PROTOCOL_GOOGLE_MAPS:          PROTO_GOOGLE_MAPS,
	NDPI_PROTOCOL_YOUTUBE:              PROTO_YOUTUBE,
	NDPI_PROTOCOL_SKYPE_TEAMS:          PROTO_SKYPE_TEAMS,
	NDPI_PROTOCOL_GOOGLE:               PROTO_GOOGLE,
	NDPI_PROTOCOL_RPC:                  PROTO_RPC,
	NDPI_PROTOCOL_NETFLOW:              PROTO_NETFLOW,
	NDPI_PROTOCOL_SFLOW:                PROTO_SFLOW,
	NDPI_PROTOCOL_HTTP_CONNECT:         PROTO_HTTP_CONNECT,
	NDPI_PROTOCOL_HTTP_PROXY:           PROTO_HTTP_PROXY,
	NDPI_PROTOCOL_CITRIX:               PROTO_CITRIX,
	NDPI_PROTOCOL_NETFLIX:              PROTO_NETFLIX,
	NDPI_PROTOCOL_LASTFM:               PROTO_LASTFM,
	NDPI_PROTOCOL_WAZE:                 PROTO_WAZE,
	NDPI_PROTOCOL_YOUTUBE_UPLOAD:       PROTO_YOUTUBE_UPLOAD,
	NDPI_PROTOCOL_HULU:                 PROTO_HULU,
	NDPI_PROTOCOL_CHECKMK:              PROTO_CHECKMK,
	NDPI_PROTOCOL_AJP:                  PROTO_AJP,
	NDPI_PROTOCOL_APPLE:                PROTO_APPLE,
	NDPI_PROTOCOL_WEBEX:                PROTO_WEBEX,
	NDPI_PROTOCOL_WHATSAPP:             PROTO_WHATSAPP,
	NDPI_PROTOCOL_APPLE_ICLOUD:         PROTO_APPLE_ICLOUD,
	NDPI_PROTOCOL_VIBER:                PROTO_VIBER,
	NDPI_PROTOCOL_APPLE_ITUNES:         PROTO_APPLE_ITUNES,
	NDPI_PROTOCOL_RADIUS:               PROTO_RADIUS,
	NDPI_PROTOCOL_WINDOWS_UPDATE:       PROTO_WINDOWS_UPDATE,
	NDPI_PROTOCOL_TEAMVIEWER:           PROTO_TEAMVIEWER,
	NDPI_PROTOCOL_TUENTI:               PROTO_TUENTI,
	NDPI_PROTOCOL_LOTUS_NOTES:          PROTO_LOTUS_NOTES,
	NDPI_PROTOCOL_SAP:                  PROTO_SAP,
	NDPI_PROTOCOL_GTP:                  PROTO_GTP,
	NDPI_PROTOCOL_WSD:                  PROTO_WSD,
	NDPI_PROTOCOL_LLMNR:                PROTO_LLMNR,
	NDPI_PROTOCOL_TOCA_BOCA:            PROTO_TOCA_BOCA,
	NDPI_PROTOCOL_SPOTIFY:              PROTO_SPOTIFY,
	NDPI_PROTOCOL_MESSENGER:            PROTO_MESSENGER,
	NDPI_PROTOCOL_H323:                 PROTO_H323,
	NDPI_PROTOCOL_OPENVPN:              PROTO_OPENVPN,
	NDPI_PROTOCOL_NOE:                  PROTO_NOE,
	NDPI_PROTOCOL_CISCOVPN:             PROTO_CISCOVPN,
	NDPI_PROTOCOL_TEAMSPEAK:            PROTO_TEAMSPEAK,
	NDPI_PROTOCOL_TOR:                  PROTO_TOR,
	NDPI_PROTOCOL_SKINNY:               PROTO_SKINNY,
	NDPI_PROTOCOL_RTCP:                 PROTO_RTCP,
	NDPI_PROTOCOL_RSYNC:                PROTO_RSYNC,
	NDPI_PROTOCOL_ORACLE:               PROTO_ORACLE,
	NDPI_PROTOCOL_CORBA:                PROTO_CORBA,
	NDPI_PROTOCOL_UBUNTUONE:            PROTO_UBUNTUONE,
	NDPI_PROTOCOL_WHOIS_DAS:            PROTO_WHOIS_DAS,
	NDPI_PROTOCOL_SD_RTN:               PROTO_SD_RTN,
	NDPI_PROTOCOL_SOCKS:                PROTO_SOCKS,
	NDPI_PROTOCOL_NINTENDO:             PROTO_NINTENDO,
	NDPI_PROTOCOL_RTMP:                 PROTO_RTMP,
	NDPI_PROTOCOL_FTP_DATA:             PROTO_FTP_DATA,
	NDPI_PROTOCOL_WIKIPEDIA:            PROTO_WIKIPEDIA,
	NDPI_PROTOCOL_ZMQ:                  PROTO_ZMQ,
	NDPI_PROTOCOL_AMAZON:               PROTO_AMAZON,
	NDPI_PROTOCOL_EBAY:                 PROTO_EBAY,
	NDPI_PROTOCOL_CNN:                  PROTO_CNN,
	NDPI_PROTOCOL_MEGACO:               PROTO_MEGACO,
	NDPI_PROTOCOL_REDIS:                PROTO_REDIS,
	NDPI_PROTOCOL_PINTEREST:            PROTO_PINTEREST,
	NDPI_PROTOCOL_VHUA:                 PROTO_VHUA,
	NDPI_PROTOCOL_TELEGRAM:             PROTO_TELEGRAM,
	NDPI_PROTOCOL_VEVO:                 PROTO_VEVO,
	NDPI_PROTOCOL_PANDORA:              PROTO_PANDORA,
	NDPI_PROTOCOL_QUIC:                 PROTO_QUIC,
	NDPI_PROTOCOL_ZOOM:                 PROTO_ZOOM,
	NDPI_PROTOCOL_EAQ:                  PROTO_EAQ,
	NDPI_PROTOCOL_OOKLA:                PROTO_OOKLA,
	NDPI_PROTOCOL_AMQP:                 PROTO_AMQP,
	NDPI_PROTOCOL_KAKAOTALK:            PROTO_KAKAOTALK,
	NDPI_PROTOCOL_KAKAOTALK_VOICE:      PROTO_KAKAOTALK_VOICE,
	NDPI_PROTOCOL_TWITCH:               PROTO_TWITCH,
	NDPI_PROTOCOL_DOH_DOT:              PROTO_DOH_DOT,
	NDPI_PROTOCOL_WECHAT:               PROTO_WECHAT,
	NDPI_PROTOCOL_MPEGTS:               PROTO_MPEGTS,
	NDPI_PROTOCOL_SNAPCHAT:             PROTO_SNAPCHAT,
	NDPI_PROTOCOL_SINA:                 PROTO_SINA,
	NDPI_PROTOCOL_HANGOUT_DUO:          PROTO_HANGOUT_DUO,
	NDPI_PROTOCOL_IFLIX:                PROTO_IFLIX,
	NDPI_PROTOCOL_GITHUB:               PROTO_GITHUB,
	NDPI_PROTOCOL_BJNP:                 PROTO_BJNP,
	NDPI_PROTOCOL_REDDIT:               PROTO_REDDIT,
	NDPI_PROTOCOL_WIREGUARD:            PROTO_WIREGUARD,
	NDPI_PROTOCOL_SMPP:                 PROTO_SMPP,
	NDPI_PROTOCOL_DNSCRYPT:             PROTO_DNSCRYPT,
	NDPI_PROTOCOL_TINC:                 PROTO_TINC,
	NDPI_PROTOCOL_DEEZER:               PROTO_DEEZER,
	NDPI_PROTOCOL_INSTAGRAM:            PROTO_INSTAGRAM,
	NDPI_PROTOCOL_MICROSOFT:            PROTO_MICROSOFT,
	NDPI_PROTOCOL_STARCRAFT:            PROTO_STARCRAFT,
	NDPI_PROTOCOL_TEREDO:               PROTO_TEREDO,
	NDPI_PROTOCOL_HOTSPOT_SHIELD:       PROTO_HOTSPOT_SHIELD,
	NDPI_PROTOCOL_IMO:                  PROTO_IMO,
	NDPI_PROTOCOL_GOOGLE_DRIVE:         PROTO_GOOGLE_DRIVE,
	NDPI_PROTOCOL_OCS:                  PROTO_OCS,
	NDPI_PROTOCOL_MICROSOFT_365:        PROTO_MICROSOFT_365,
	NDPI_PROTOCOL_CLOUDFLARE:           PROTO_CLOUDFLARE,
	NDPI_PROTOCOL_MS_ONE_DRIVE:         PROTO_MS_ONE_DRIVE,
	NDPI_PROTOCOL_MQTT:                 PROTO_MQTT,
	NDPI_PROTOCOL_RX:                   PROTO_RX,
	NDPI_PROTOCOL_APPLESTORE:           PROTO_APPLESTORE,
	NDPI_PROTOCOL_OPENDNS:              PROTO_OPENDNS,
	NDPI_PROTOCOL_GIT:                  PROTO_GIT,
	NDPI_PROTOCOL_DRDA:                 PROTO_DRDA,
	NDPI_PROTOCOL_PLAYSTORE:            PROTO_PLAYSTORE,
	NDPI_PROTOCOL_SOMEIP:               PROTO_SOMEIP,
	NDPI_PROTOCOL_FIX:                  PROTO_FIX,
	NDPI_PROTOCOL_PLAYSTATION:          PROTO_PLAYSTATION,
	NDPI_PROTOCOL_PASTEBIN:             PROTO_PASTEBIN,
	NDPI_PROTOCOL_LINKEDIN:             PROTO_LINKEDIN,
	NDPI_PROTOCOL_SOUNDCLOUD:           PROTO_SOUNDCLOUD,
	NDPI_PROTOCOL_CSGO:                 PROTO_CSGO,
	NDPI_PROTOCOL_LISP:                 PROTO_LISP,
	NDPI_PROTOCOL_DIAMETER:             PROTO_DIAMETER,
	NDPI_PROTOCOL_APPLE_PUSH:           PROTO_APPLE_PUSH,
	NDPI_PROTOCOL_GOOGLE_SERVICES:      PROTO_GOOGLE_SERVICES,
	NDPI_PROTOCOL_AMAZON_VIDEO:         PROTO_AMAZON_VIDEO,
	NDPI_PROTOCOL_GOOGLE_DOCS:          PROTO_GOOGLE_DOCS,
	NDPI_PROTOCOL_WHATSAPP_FILES:       PROTO_WHATSAPP_FILES,
	NDPI_PROTOCOL_TARGUS_GETDATA:       PROTO_TARGUS_GETDATA,
	NDPI_PROTOCOL_DNP3:                 PROTO_DNP3,
	NDPI_PROTOCOL_IEC60870:             PROTO_IEC60870,
	NDPI_PROTOCOL_BLOOMBERG:            PROTO_BLOOMBERG,
	NDPI_PROTOCOL_CAPWAP:               PROTO_CAPWAP,
	NDPI_PROTOCOL_ZABBIX:               PROTO_ZABBIX,
	NDPI_PROTOCOL_S7COMM:               PROTO_S7COMM,
	NDPI_PROTOCOL_MSTEAMS:              PROTO_MSTEAMS,
	NDPI_PROTOCOL_WEBSOCKET:            PROTO_WEBSOCKET,
	NDPI_PROTOCOL_ANYDESK:              PROTO_ANYDESK,
	NDPI_PROTOCOL_SOAP:                 PROTO_SOAP,
	NDPI_PROTOCOL_APPLE_SIRI:           PROTO_APPLE_SIRI,
	NDPI_PROTOCOL_SNAPCHAT_CALL:        PROTO_SNAPCHAT_CALL,
	NDPI_PROTOCOL_HPVIRTGRP:            PROTO_HPVIRTGRP,
	NDPI_PROTOCOL_GENSHIN_IMPACT:       PROTO_GENSHIN_IMPACT,
	NDPI_PROTOCOL_ACTIVISION:           PROTO_ACTIVISION,
	NDPI_PROTOCOL_FORTICLIENT:          PROTO_FORTICLIENT,
	NDPI_PROTOCOL_Z3950:                PROTO_Z3950,
	NDPI_PROTOCOL_LIKEE:                PROTO_LIKEE,
	NDPI_PROTOCOL_GITLAB:               PROTO_GITLAB,
	NDPI_PROTOCOL_AVAST_SECUREDNS:      PROTO_AVAST_SECUREDNS,
	NDPI_PROTOCOL_CASSANDRA:            PROTO_CASSANDRA,
	NDPI_PROTOCOL_AMAZON_AWS:           PROTO_AMAZON_AWS,
	NDPI_PROTOCOL_SALESFORCE:           PROTO_SALESFORCE,
	NDPI_PROTOCOL_VIMEO:                PROTO_VIMEO,
	NDPI_PROTOCOL_FACEBOOK_VOIP:        PROTO_FACEBOOK_VOIP,
	NDPI_PROTOCOL_SIGNAL_VOIP:          PROTO_SIGNAL_VOIP,
	NDPI_PROTOCOL_FUZE:                 PROTO_FUZE,
	NDPI_PROTOCOL_GTP_U:                PROTO_GTP_U,
	NDPI_PROTOCOL_GTP_C:                PROTO_GTP_C,
	NDPI_PROTOCOL_GTP_PRIME:            PROTO_GTP_PRIME,
	NDPI_PROTOCOL_ALIBABA:              PROTO_ALIBABA,
	NDPI_PROTOCOL_CRASHLYSTICS:         PROTO_CRASHLYSTICS,
	NDPI_PROTOCOL_MICROSOFT_AZURE:      PROTO_MICROSOFT_AZURE,
	NDPI_PROTOCOL_ICLOUD_PRIVATE_RELAY: PROTO_ICLOUD_PRIVATE_RELAY,
	NDPI_PROTOCOL_ETHERNET_IP:          PROTO_ETHERNET_IP,
	NDPI_PROTOCOL_BADOO:                PROTO_BADOO,
	NDPI_PROTOCOL_ACCUWEATHER:          PROTO_ACCUWEATHER,
	NDPI_PROTOCOL_GOOGLE_CLASSROOM:     PROTO_GOOGLE_CLASSROOM,
	NDPI_PROTOCOL_HSRP:                 PROTO_HSRP,
	NDPI_PROTOCOL_CYBERSECURITY:        PROTO_CYBERSECURITY,
	NDPI_PROTOCOL_GOOGLE_CLOUD:         PROTO_GOOGLE_CLOUD,
	NDPI_PROTOCOL_TENCENT:              PROTO_TENCENT,
	NDPI_PROTOCOL_RAKNET:               PROTO_RAKNET,
	NDPI_PROTOCOL_XIAOMI:               PROTO_XIAOMI,
	NDPI_PROTOCOL_EDGECAST:             PROTO_EDGECAST,
	NDPI_PROTOCOL_CACHEFLY:             PROTO_CACHEFLY,
	NDPI_PROTOCOL_SOFTETHER:            PROTO_SOFTETHER,
	NDPI_PROTOCOL_MPEGDASH:             PROTO_MPEGDASH,
	NDPI_PROTOCOL_DAZN:                 PROTO_DAZN,
	NDPI_PROTOCOL_GOTO:                 PROTO_GOTO,
	NDPI_PROTOCOL_RSH:                  PROTO_RSH,
	NDPI_PROTOCOL_1KXUN:                PROTO_1KXUN,
	NDPI_PROTOCOL_IP_PGM:               PROTO_IP_PGM,
	NDPI_PROTOCOL_IP_PIM:               PROTO_IP_PIM,
	NDPI_PROTOCOL_COLLECTD:             PROTO_COLLECTD,
	NDPI_PROTOCOL_TUNNELBEAR:           PROTO_TUNNELBEAR,
	NDPI_PROTOCOL_CLOUDFLARE_WARP:      PROTO_CLOUDFLARE_WARP,
	NDPI_PROTOCOL_I3D:                  PROTO_I3D,
	NDPI_PROTOCOL_RIOTGAMES:            PROTO_RIOTGAMES,
	NDPI_PROTOCOL_PSIPHON:              PROTO_PSIPHON,
	NDPI_PROTOCOL_ULTRASURF:            PROTO_ULTRASURF,
}
