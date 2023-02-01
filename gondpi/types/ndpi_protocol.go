package types

type NdpiProtocol uint16

func (p *NdpiProtocol) ToName() string {
	name, ok := NdpiProtocolNameMap[*p]
	if !ok {
		name = ""
	}

	return name
}

const (
	NDPI_PROTOCOL_UNKNOWN              NdpiProtocol = 0
	NDPI_PROTOCOL_FTP_CONTROL          NdpiProtocol = 1
	NDPI_PROTOCOL_MAIL_POP             NdpiProtocol = 2
	NDPI_PROTOCOL_MAIL_SMTP            NdpiProtocol = 3
	NDPI_PROTOCOL_MAIL_IMAP            NdpiProtocol = 4
	NDPI_PROTOCOL_DNS                  NdpiProtocol = 5
	NDPI_PROTOCOL_IPP                  NdpiProtocol = 6
	NDPI_PROTOCOL_HTTP                 NdpiProtocol = 7
	NDPI_PROTOCOL_MDNS                 NdpiProtocol = 8
	NDPI_PROTOCOL_NTP                  NdpiProtocol = 9
	NDPI_PROTOCOL_NETBIOS              NdpiProtocol = 10
	NDPI_PROTOCOL_NFS                  NdpiProtocol = 11
	NDPI_PROTOCOL_SSDP                 NdpiProtocol = 12
	NDPI_PROTOCOL_BGP                  NdpiProtocol = 13
	NDPI_PROTOCOL_SNMP                 NdpiProtocol = 14
	NDPI_PROTOCOL_XDMCP                NdpiProtocol = 15
	NDPI_PROTOCOL_SMBV1                NdpiProtocol = 16
	NDPI_PROTOCOL_SYSLOG               NdpiProtocol = 17
	NDPI_PROTOCOL_DHCP                 NdpiProtocol = 18
	NDPI_PROTOCOL_POSTGRES             NdpiProtocol = 19
	NDPI_PROTOCOL_MYSQL                NdpiProtocol = 20
	NDPI_PROTOCOL_MS_OUTLOOK           NdpiProtocol = 21
	NDPI_PROTOCOL_DIRECT_DOWNLOAD_LINK NdpiProtocol = 22
	NDPI_PROTOCOL_MAIL_POPS            NdpiProtocol = 23
	NDPI_PROTOCOL_APPLEJUICE           NdpiProtocol = 24
	NDPI_PROTOCOL_DIRECTCONNECT        NdpiProtocol = 25
	NDPI_PROTOCOL_NTOP                 NdpiProtocol = 26
	NDPI_PROTOCOL_COAP                 NdpiProtocol = 27
	NDPI_PROTOCOL_VMWARE               NdpiProtocol = 28
	NDPI_PROTOCOL_MAIL_SMTPS           NdpiProtocol = 29
	NDPI_PROTOCOL_DTLS                 NdpiProtocol = 30
	NDPI_PROTOCOL_UBNTAC2              NdpiProtocol = 31
	NDPI_PROTOCOL_KONTIKI              NdpiProtocol = 32
	NDPI_PROTOCOL_OPENFT               NdpiProtocol = 33
	NDPI_PROTOCOL_FASTTRACK            NdpiProtocol = 34
	NDPI_PROTOCOL_GNUTELLA             NdpiProtocol = 35
	NDPI_PROTOCOL_EDONKEY              NdpiProtocol = 36
	NDPI_PROTOCOL_BITTORRENT           NdpiProtocol = 37
	NDPI_PROTOCOL_SKYPE_TEAMS_CALL     NdpiProtocol = 38
	NDPI_PROTOCOL_SIGNAL               NdpiProtocol = 39
	NDPI_PROTOCOL_MEMCACHED            NdpiProtocol = 40
	NDPI_PROTOCOL_SMBV23               NdpiProtocol = 41
	NDPI_PROTOCOL_MINING               NdpiProtocol = 42
	NDPI_PROTOCOL_NEST_LOG_SINK        NdpiProtocol = 43
	NDPI_PROTOCOL_MODBUS               NdpiProtocol = 44
	NDPI_PROTOCOL_WHATSAPP_CALL        NdpiProtocol = 45
	NDPI_PROTOCOL_DATASAVER            NdpiProtocol = 46
	NDPI_PROTOCOL_XBOX                 NdpiProtocol = 47
	NDPI_PROTOCOL_QQ                   NdpiProtocol = 48
	NDPI_PROTOCOL_TIKTOK               NdpiProtocol = 49
	NDPI_PROTOCOL_RTSP                 NdpiProtocol = 50
	NDPI_PROTOCOL_MAIL_IMAPS           NdpiProtocol = 51
	NDPI_PROTOCOL_ICECAST              NdpiProtocol = 52
	NDPI_PROTOCOL_CPHA                 NdpiProtocol = 53
	NDPI_PROTOCOL_PPSTREAM             NdpiProtocol = 54
	NDPI_PROTOCOL_ZATTOO               NdpiProtocol = 55
	NDPI_PROTOCOL_SHOUTCAST            NdpiProtocol = 56
	NDPI_PROTOCOL_SOPCAST              NdpiProtocol = 57
	NDPI_PROTOCOL_DISCORD              NdpiProtocol = 58
	NDPI_PROTOCOL_TVUPLAYER            NdpiProtocol = 59
	NDPI_PROTOCOL_MONGODB              NdpiProtocol = 60
	NDPI_PROTOCOL_PLURALSIGHT          NdpiProtocol = 61
	NDPI_PROTOCOL_THUNDER              NdpiProtocol = 62
	NDPI_PROTOCOL_OCSP                 NdpiProtocol = 63
	NDPI_PROTOCOL_VXLAN                NdpiProtocol = 64
	NDPI_PROTOCOL_IRC                  NdpiProtocol = 65
	NDPI_PROTOCOL_AYIYA                NdpiProtocol = 66
	NDPI_PROTOCOL_JABBER               NdpiProtocol = 67
	NDPI_PROTOCOL_NATS                 NdpiProtocol = 68
	NDPI_PROTOCOL_AMONG_US             NdpiProtocol = 69
	NDPI_PROTOCOL_YAHOO                NdpiProtocol = 70
	NDPI_PROTOCOL_DISNEYPLUS           NdpiProtocol = 71
	NDPI_PROTOCOL_GOOGLE_PLUS          NdpiProtocol = 72
	NDPI_PROTOCOL_IP_VRRP              NdpiProtocol = 73
	NDPI_PROTOCOL_STEAM                NdpiProtocol = 74
	NDPI_PROTOCOL_HALFLIFE2            NdpiProtocol = 75
	NDPI_PROTOCOL_WORLDOFWARCRAFT      NdpiProtocol = 76
	NDPI_PROTOCOL_TELNET               NdpiProtocol = 77
	NDPI_PROTOCOL_STUN                 NdpiProtocol = 78
	NDPI_PROTOCOL_IPSEC                NdpiProtocol = 79
	NDPI_PROTOCOL_IP_GRE               NdpiProtocol = 80
	NDPI_PROTOCOL_IP_ICMP              NdpiProtocol = 81
	NDPI_PROTOCOL_IP_IGMP              NdpiProtocol = 82
	NDPI_PROTOCOL_IP_EGP               NdpiProtocol = 83
	NDPI_PROTOCOL_IP_SCTP              NdpiProtocol = 84
	NDPI_PROTOCOL_IP_OSPF              NdpiProtocol = 85
	NDPI_PROTOCOL_IP_IP_IN_IP          NdpiProtocol = 86
	NDPI_PROTOCOL_RTP                  NdpiProtocol = 87
	NDPI_PROTOCOL_RDP                  NdpiProtocol = 88
	NDPI_PROTOCOL_VNC                  NdpiProtocol = 89
	NDPI_PROTOCOL_TUMBLR               NdpiProtocol = 90
	NDPI_PROTOCOL_TLS                  NdpiProtocol = 91
	NDPI_PROTOCOL_SSH                  NdpiProtocol = 92
	NDPI_PROTOCOL_USENET               NdpiProtocol = 93
	NDPI_PROTOCOL_MGCP                 NdpiProtocol = 94
	NDPI_PROTOCOL_IAX                  NdpiProtocol = 95
	NDPI_PROTOCOL_TFTP                 NdpiProtocol = 96
	NDPI_PROTOCOL_AFP                  NdpiProtocol = 97
	NDPI_PROTOCOL_STEALTHNET           NdpiProtocol = 98
	NDPI_PROTOCOL_AIMINI               NdpiProtocol = 99
	NDPI_PROTOCOL_SIP                  NdpiProtocol = 100
	NDPI_PROTOCOL_TRUPHONE             NdpiProtocol = 101
	NDPI_PROTOCOL_IP_ICMPV6            NdpiProtocol = 102
	NDPI_PROTOCOL_DHCPV6               NdpiProtocol = 103
	NDPI_PROTOCOL_ARMAGETRON           NdpiProtocol = 104
	NDPI_PROTOCOL_CROSSFIRE            NdpiProtocol = 105
	NDPI_PROTOCOL_DOFUS                NdpiProtocol = 106
	NDPI_PROTOCOL_FIESTA               NdpiProtocol = 107
	NDPI_PROTOCOL_FLORENSIA            NdpiProtocol = 108
	NDPI_PROTOCOL_GUILDWARS            NdpiProtocol = 109
	NDPI_PROTOCOL_AMAZON_ALEXA         NdpiProtocol = 110
	NDPI_PROTOCOL_KERBEROS             NdpiProtocol = 111
	NDPI_PROTOCOL_LDAP                 NdpiProtocol = 112
	NDPI_PROTOCOL_MAPLESTORY           NdpiProtocol = 113
	NDPI_PROTOCOL_MSSQL_TDS            NdpiProtocol = 114
	NDPI_PROTOCOL_PPTP                 NdpiProtocol = 115
	NDPI_PROTOCOL_WARCRAFT3            NdpiProtocol = 116
	NDPI_PROTOCOL_WORLD_OF_KUNG_FU     NdpiProtocol = 117
	NDPI_PROTOCOL_SLACK                NdpiProtocol = 118
	NDPI_PROTOCOL_FACEBOOK             NdpiProtocol = 119
	NDPI_PROTOCOL_TWITTER              NdpiProtocol = 120
	NDPI_PROTOCOL_DROPBOX              NdpiProtocol = 121
	NDPI_PROTOCOL_GMAIL                NdpiProtocol = 122
	NDPI_PROTOCOL_GOOGLE_MAPS          NdpiProtocol = 123
	NDPI_PROTOCOL_YOUTUBE              NdpiProtocol = 124
	NDPI_PROTOCOL_SKYPE_TEAMS          NdpiProtocol = 125
	NDPI_PROTOCOL_GOOGLE               NdpiProtocol = 126
	NDPI_PROTOCOL_RPC                  NdpiProtocol = 127
	NDPI_PROTOCOL_NETFLOW              NdpiProtocol = 128
	NDPI_PROTOCOL_SFLOW                NdpiProtocol = 129
	NDPI_PROTOCOL_HTTP_CONNECT         NdpiProtocol = 130
	NDPI_PROTOCOL_HTTP_PROXY           NdpiProtocol = 131
	NDPI_PROTOCOL_CITRIX               NdpiProtocol = 132
	NDPI_PROTOCOL_NETFLIX              NdpiProtocol = 133
	NDPI_PROTOCOL_LASTFM               NdpiProtocol = 134
	NDPI_PROTOCOL_WAZE                 NdpiProtocol = 135
	NDPI_PROTOCOL_YOUTUBE_UPLOAD       NdpiProtocol = 136
	NDPI_PROTOCOL_HULU                 NdpiProtocol = 137
	NDPI_PROTOCOL_CHECKMK              NdpiProtocol = 138
	NDPI_PROTOCOL_AJP                  NdpiProtocol = 139
	NDPI_PROTOCOL_APPLE                NdpiProtocol = 140
	NDPI_PROTOCOL_WEBEX                NdpiProtocol = 141
	NDPI_PROTOCOL_WHATSAPP             NdpiProtocol = 142
	NDPI_PROTOCOL_APPLE_ICLOUD         NdpiProtocol = 143
	NDPI_PROTOCOL_VIBER                NdpiProtocol = 144
	NDPI_PROTOCOL_APPLE_ITUNES         NdpiProtocol = 145
	NDPI_PROTOCOL_RADIUS               NdpiProtocol = 146
	NDPI_PROTOCOL_WINDOWS_UPDATE       NdpiProtocol = 147
	NDPI_PROTOCOL_TEAMVIEWER           NdpiProtocol = 148
	NDPI_PROTOCOL_TUENTI               NdpiProtocol = 149
	NDPI_PROTOCOL_LOTUS_NOTES          NdpiProtocol = 150
	NDPI_PROTOCOL_SAP                  NdpiProtocol = 151
	NDPI_PROTOCOL_GTP                  NdpiProtocol = 152
	NDPI_PROTOCOL_WSD                  NdpiProtocol = 153
	NDPI_PROTOCOL_LLMNR                NdpiProtocol = 154
	NDPI_PROTOCOL_TOCA_BOCA            NdpiProtocol = 155
	NDPI_PROTOCOL_SPOTIFY              NdpiProtocol = 156
	NDPI_PROTOCOL_MESSENGER            NdpiProtocol = 157
	NDPI_PROTOCOL_H323                 NdpiProtocol = 158
	NDPI_PROTOCOL_OPENVPN              NdpiProtocol = 159
	NDPI_PROTOCOL_NOE                  NdpiProtocol = 160
	NDPI_PROTOCOL_CISCOVPN             NdpiProtocol = 161
	NDPI_PROTOCOL_TEAMSPEAK            NdpiProtocol = 162
	NDPI_PROTOCOL_TOR                  NdpiProtocol = 163
	NDPI_PROTOCOL_SKINNY               NdpiProtocol = 164
	NDPI_PROTOCOL_RTCP                 NdpiProtocol = 165
	NDPI_PROTOCOL_RSYNC                NdpiProtocol = 166
	NDPI_PROTOCOL_ORACLE               NdpiProtocol = 167
	NDPI_PROTOCOL_CORBA                NdpiProtocol = 168
	NDPI_PROTOCOL_UBUNTUONE            NdpiProtocol = 169
	NDPI_PROTOCOL_WHOIS_DAS            NdpiProtocol = 170
	NDPI_PROTOCOL_SD_RTN               NdpiProtocol = 171
	NDPI_PROTOCOL_SOCKS                NdpiProtocol = 172
	NDPI_PROTOCOL_NINTENDO             NdpiProtocol = 173
	NDPI_PROTOCOL_RTMP                 NdpiProtocol = 174
	NDPI_PROTOCOL_FTP_DATA             NdpiProtocol = 175
	NDPI_PROTOCOL_WIKIPEDIA            NdpiProtocol = 176
	NDPI_PROTOCOL_ZMQ                  NdpiProtocol = 177
	NDPI_PROTOCOL_AMAZON               NdpiProtocol = 178
	NDPI_PROTOCOL_EBAY                 NdpiProtocol = 179
	NDPI_PROTOCOL_CNN                  NdpiProtocol = 180
	NDPI_PROTOCOL_MEGACO               NdpiProtocol = 181
	NDPI_PROTOCOL_REDIS                NdpiProtocol = 182
	NDPI_PROTOCOL_PINTEREST            NdpiProtocol = 183
	NDPI_PROTOCOL_VHUA                 NdpiProtocol = 184
	NDPI_PROTOCOL_TELEGRAM             NdpiProtocol = 185
	NDPI_PROTOCOL_VEVO                 NdpiProtocol = 186
	NDPI_PROTOCOL_PANDORA              NdpiProtocol = 187
	NDPI_PROTOCOL_QUIC                 NdpiProtocol = 188
	NDPI_PROTOCOL_ZOOM                 NdpiProtocol = 189
	NDPI_PROTOCOL_EAQ                  NdpiProtocol = 190
	NDPI_PROTOCOL_OOKLA                NdpiProtocol = 191
	NDPI_PROTOCOL_AMQP                 NdpiProtocol = 192
	NDPI_PROTOCOL_KAKAOTALK            NdpiProtocol = 193
	NDPI_PROTOCOL_KAKAOTALK_VOICE      NdpiProtocol = 194
	NDPI_PROTOCOL_TWITCH               NdpiProtocol = 195
	NDPI_PROTOCOL_DOH_DOT              NdpiProtocol = 196
	NDPI_PROTOCOL_WECHAT               NdpiProtocol = 197
	NDPI_PROTOCOL_MPEGTS               NdpiProtocol = 198
	NDPI_PROTOCOL_SNAPCHAT             NdpiProtocol = 199
	NDPI_PROTOCOL_SINA                 NdpiProtocol = 200
	NDPI_PROTOCOL_HANGOUT_DUO          NdpiProtocol = 201
	NDPI_PROTOCOL_IFLIX                NdpiProtocol = 202
	NDPI_PROTOCOL_GITHUB               NdpiProtocol = 203
	NDPI_PROTOCOL_BJNP                 NdpiProtocol = 204
	NDPI_PROTOCOL_REDDIT               NdpiProtocol = 205
	NDPI_PROTOCOL_WIREGUARD            NdpiProtocol = 206
	NDPI_PROTOCOL_SMPP                 NdpiProtocol = 207
	NDPI_PROTOCOL_DNSCRYPT             NdpiProtocol = 208
	NDPI_PROTOCOL_TINC                 NdpiProtocol = 209
	NDPI_PROTOCOL_DEEZER               NdpiProtocol = 210
	NDPI_PROTOCOL_INSTAGRAM            NdpiProtocol = 211
	NDPI_PROTOCOL_MICROSOFT            NdpiProtocol = 212
	NDPI_PROTOCOL_STARCRAFT            NdpiProtocol = 213
	NDPI_PROTOCOL_TEREDO               NdpiProtocol = 214
	NDPI_PROTOCOL_HOTSPOT_SHIELD       NdpiProtocol = 215
	NDPI_PROTOCOL_IMO                  NdpiProtocol = 216
	NDPI_PROTOCOL_GOOGLE_DRIVE         NdpiProtocol = 217
	NDPI_PROTOCOL_OCS                  NdpiProtocol = 218
	NDPI_PROTOCOL_MICROSOFT_365        NdpiProtocol = 219
	NDPI_PROTOCOL_CLOUDFLARE           NdpiProtocol = 220
	NDPI_PROTOCOL_MS_ONE_DRIVE         NdpiProtocol = 221
	NDPI_PROTOCOL_MQTT                 NdpiProtocol = 222
	NDPI_PROTOCOL_RX                   NdpiProtocol = 223
	NDPI_PROTOCOL_APPLESTORE           NdpiProtocol = 224
	NDPI_PROTOCOL_OPENDNS              NdpiProtocol = 225
	NDPI_PROTOCOL_GIT                  NdpiProtocol = 226
	NDPI_PROTOCOL_DRDA                 NdpiProtocol = 227
	NDPI_PROTOCOL_PLAYSTORE            NdpiProtocol = 228
	NDPI_PROTOCOL_SOMEIP               NdpiProtocol = 229
	NDPI_PROTOCOL_FIX                  NdpiProtocol = 230
	NDPI_PROTOCOL_PLAYSTATION          NdpiProtocol = 231
	NDPI_PROTOCOL_PASTEBIN             NdpiProtocol = 232
	NDPI_PROTOCOL_LINKEDIN             NdpiProtocol = 233
	NDPI_PROTOCOL_SOUNDCLOUD           NdpiProtocol = 234
	NDPI_PROTOCOL_CSGO                 NdpiProtocol = 235
	NDPI_PROTOCOL_LISP                 NdpiProtocol = 236
	NDPI_PROTOCOL_DIAMETER             NdpiProtocol = 237
	NDPI_PROTOCOL_APPLE_PUSH           NdpiProtocol = 238
	NDPI_PROTOCOL_GOOGLE_SERVICES      NdpiProtocol = 239
	NDPI_PROTOCOL_AMAZON_VIDEO         NdpiProtocol = 240
	NDPI_PROTOCOL_GOOGLE_DOCS          NdpiProtocol = 241
	NDPI_PROTOCOL_WHATSAPP_FILES       NdpiProtocol = 242
	NDPI_PROTOCOL_TARGUS_GETDATA       NdpiProtocol = 243
	NDPI_PROTOCOL_DNP3                 NdpiProtocol = 244
	NDPI_PROTOCOL_IEC60870             NdpiProtocol = 245
	NDPI_PROTOCOL_BLOOMBERG            NdpiProtocol = 246
	NDPI_PROTOCOL_CAPWAP               NdpiProtocol = 247
	NDPI_PROTOCOL_ZABBIX               NdpiProtocol = 248
	NDPI_PROTOCOL_S7COMM               NdpiProtocol = 249
	NDPI_PROTOCOL_MSTEAMS              NdpiProtocol = 250
	NDPI_PROTOCOL_WEBSOCKET            NdpiProtocol = 251
	NDPI_PROTOCOL_ANYDESK              NdpiProtocol = 252
	NDPI_PROTOCOL_SOAP                 NdpiProtocol = 253
	NDPI_PROTOCOL_APPLE_SIRI           NdpiProtocol = 254
	NDPI_PROTOCOL_SNAPCHAT_CALL        NdpiProtocol = 255
	NDPI_PROTOCOL_HPVIRTGRP            NdpiProtocol = 256
	NDPI_PROTOCOL_GENSHIN_IMPACT       NdpiProtocol = 257
	NDPI_PROTOCOL_ACTIVISION           NdpiProtocol = 258
	NDPI_PROTOCOL_FORTICLIENT          NdpiProtocol = 259
	NDPI_PROTOCOL_Z3950                NdpiProtocol = 260
	NDPI_PROTOCOL_LIKEE                NdpiProtocol = 261
	NDPI_PROTOCOL_GITLAB               NdpiProtocol = 262
	NDPI_PROTOCOL_AVAST_SECUREDNS      NdpiProtocol = 263
	NDPI_PROTOCOL_CASSANDRA            NdpiProtocol = 264
	NDPI_PROTOCOL_AMAZON_AWS           NdpiProtocol = 265
	NDPI_PROTOCOL_SALESFORCE           NdpiProtocol = 266
	NDPI_PROTOCOL_VIMEO                NdpiProtocol = 267
	NDPI_PROTOCOL_FACEBOOK_VOIP        NdpiProtocol = 268
	NDPI_PROTOCOL_SIGNAL_VOIP          NdpiProtocol = 269
	NDPI_PROTOCOL_FUZE                 NdpiProtocol = 270
	NDPI_PROTOCOL_GTP_U                NdpiProtocol = 271
	NDPI_PROTOCOL_GTP_C                NdpiProtocol = 272
	NDPI_PROTOCOL_GTP_PRIME            NdpiProtocol = 273
	NDPI_PROTOCOL_ALIBABA              NdpiProtocol = 274
	NDPI_PROTOCOL_CRASHLYSTICS         NdpiProtocol = 275
	NDPI_PROTOCOL_MICROSOFT_AZURE      NdpiProtocol = 276
	NDPI_PROTOCOL_ICLOUD_PRIVATE_RELAY NdpiProtocol = 277
	NDPI_PROTOCOL_ETHERNET_IP          NdpiProtocol = 278
	NDPI_PROTOCOL_BADOO                NdpiProtocol = 279
	NDPI_PROTOCOL_ACCUWEATHER          NdpiProtocol = 280
	NDPI_PROTOCOL_GOOGLE_CLASSROOM     NdpiProtocol = 281
	NDPI_PROTOCOL_HSRP                 NdpiProtocol = 282
	NDPI_PROTOCOL_CYBERSECURITY        NdpiProtocol = 283
	NDPI_PROTOCOL_GOOGLE_CLOUD         NdpiProtocol = 284
	NDPI_PROTOCOL_TENCENT              NdpiProtocol = 285
	NDPI_PROTOCOL_RAKNET               NdpiProtocol = 286
	NDPI_PROTOCOL_XIAOMI               NdpiProtocol = 287
	NDPI_PROTOCOL_EDGECAST             NdpiProtocol = 288
	NDPI_PROTOCOL_CACHEFLY             NdpiProtocol = 289
	NDPI_PROTOCOL_SOFTETHER            NdpiProtocol = 290
	NDPI_PROTOCOL_MPEGDASH             NdpiProtocol = 291
	NDPI_PROTOCOL_DAZN                 NdpiProtocol = 292
	NDPI_PROTOCOL_GOTO                 NdpiProtocol = 293
	NDPI_PROTOCOL_RSH                  NdpiProtocol = 294
	NDPI_PROTOCOL_1KXUN                NdpiProtocol = 295
	NDPI_PROTOCOL_IP_PGM               NdpiProtocol = 296
	NDPI_PROTOCOL_IP_PIM               NdpiProtocol = 297
	NDPI_PROTOCOL_COLLECTD             NdpiProtocol = 298
	NDPI_PROTOCOL_TUNNELBEAR           NdpiProtocol = 299
	NDPI_PROTOCOL_CLOUDFLARE_WARP      NdpiProtocol = 300
	NDPI_PROTOCOL_I3D                  NdpiProtocol = 301
	NDPI_PROTOCOL_RIOTGAMES            NdpiProtocol = 302
	NDPI_PROTOCOL_PSIPHON              NdpiProtocol = 303
	NDPI_PROTOCOL_ULTRASURF            NdpiProtocol = 304
)

const (
	PROTO_UNKNOWN              string = "UNKNOWN"
	PROTO_FTP_CONTROL          string = "FTP_CONTROL"
	PROTO_MAIL_POP             string = "MAIL_POP"
	PROTO_MAIL_SMTP            string = "MAIL_SMTP"
	PROTO_MAIL_IMAP            string = "MAIL_IMAP"
	PROTO_DNS                  string = "DNS"
	PROTO_IPP                  string = "IPP"
	PROTO_HTTP                 string = "HTTP"
	PROTO_MDNS                 string = "MDNS"
	PROTO_NTP                  string = "NTP"
	PROTO_NETBIOS              string = "NETBIOS"
	PROTO_NFS                  string = "NFS"
	PROTO_SSDP                 string = "SSDP"
	PROTO_BGP                  string = "BGP"
	PROTO_SNMP                 string = "SNMP"
	PROTO_XDMCP                string = "XDMCP"
	PROTO_SMBV1                string = "SMBV1"
	PROTO_SYSLOG               string = "SYSLOG"
	PROTO_DHCP                 string = "DHCP"
	PROTO_POSTGRES             string = "POSTGRES"
	PROTO_MYSQL                string = "MYSQL"
	PROTO_MS_OUTLOOK           string = "MS_OUTLOOK"
	PROTO_DIRECT_DOWNLOAD_LINK string = "DIRECT_DOWNLOAD_LINK"
	PROTO_MAIL_POPS            string = "MAIL_POPS"
	PROTO_APPLEJUICE           string = "APPLEJUICE"
	PROTO_DIRECTCONNECT        string = "DIRECTCONNECT"
	PROTO_NTOP                 string = "NTOP"
	PROTO_COAP                 string = "COAP"
	PROTO_VMWARE               string = "VMWARE"
	PROTO_MAIL_SMTPS           string = "MAIL_SMTPS"
	PROTO_DTLS                 string = "DTLS"
	PROTO_UBNTAC2              string = "UBNTAC2"
	PROTO_KONTIKI              string = "KONTIKI"
	PROTO_OPENFT               string = "OPENFT"
	PROTO_FASTTRACK            string = "FASTTRACK"
	PROTO_GNUTELLA             string = "GNUTELLA"
	PROTO_EDONKEY              string = "EDONKEY"
	PROTO_BITTORRENT           string = "BITTORRENT"
	PROTO_SKYPE_TEAMS_CALL     string = "SKYPE_TEAMS_CALL"
	PROTO_SIGNAL               string = "SIGNAL"
	PROTO_MEMCACHED            string = "MEMCACHED"
	PROTO_SMBV23               string = "SMBV23"
	PROTO_MINING               string = "MINING"
	PROTO_NEST_LOG_SINK        string = "NEST_LOG_SINK"
	PROTO_MODBUS               string = "MODBUS"
	PROTO_WHATSAPP_CALL        string = "WHATSAPP_CALL"
	PROTO_DATASAVER            string = "DATASAVER"
	PROTO_XBOX                 string = "XBOX"
	PROTO_QQ                   string = "QQ"
	PROTO_TIKTOK               string = "TIKTOK"
	PROTO_RTSP                 string = "RTSP"
	PROTO_MAIL_IMAPS           string = "MAIL_IMAPS"
	PROTO_ICECAST              string = "ICECAST"
	PROTO_CPHA                 string = "CPHA"
	PROTO_PPSTREAM             string = "PPSTREAM"
	PROTO_ZATTOO               string = "ZATTOO"
	PROTO_SHOUTCAST            string = "SHOUTCAST"
	PROTO_SOPCAST              string = "SOPCAST"
	PROTO_DISCORD              string = "DISCORD"
	PROTO_TVUPLAYER            string = "TVUPLAYER"
	PROTO_MONGODB              string = "MONGODB"
	PROTO_PLURALSIGHT          string = "PLURALSIGHT"
	PROTO_THUNDER              string = "THUNDER"
	PROTO_OCSP                 string = "OCSP"
	PROTO_VXLAN                string = "VXLAN"
	PROTO_IRC                  string = "IRC"
	PROTO_AYIYA                string = "AYIYA"
	PROTO_JABBER               string = "JABBER"
	PROTO_NATS                 string = "NATS"
	PROTO_AMONG_US             string = "AMONG_US"
	PROTO_YAHOO                string = "YAHOO"
	PROTO_DISNEYPLUS           string = "DISNEYPLUS"
	PROTO_GOOGLE_PLUS          string = "GOOGLE_PLUS"
	PROTO_IP_VRRP              string = "IP_VRRP"
	PROTO_STEAM                string = "STEAM"
	PROTO_HALFLIFE2            string = "HALFLIFE2"
	PROTO_WORLDOFWARCRAFT      string = "WORLDOFWARCRAFT"
	PROTO_TELNET               string = "TELNET"
	PROTO_STUN                 string = "STUN"
	PROTO_IPSEC                string = "IPSEC"
	PROTO_IP_GRE               string = "IP_GRE"
	PROTO_IP_ICMP              string = "IP_ICMP"
	PROTO_IP_IGMP              string = "IP_IGMP"
	PROTO_IP_EGP               string = "IP_EGP"
	PROTO_IP_SCTP              string = "IP_SCTP"
	PROTO_IP_OSPF              string = "IP_OSPF"
	PROTO_IP_IP_IN_IP          string = "IP_IP_IN_IP"
	PROTO_RTP                  string = "RTP"
	PROTO_RDP                  string = "RDP"
	PROTO_VNC                  string = "VNC"
	PROTO_TUMBLR               string = "TUMBLR"
	PROTO_TLS                  string = "TLS"
	PROTO_SSH                  string = "SSH"
	PROTO_USENET               string = "USENET"
	PROTO_MGCP                 string = "MGCP"
	PROTO_IAX                  string = "IAX"
	PROTO_TFTP                 string = "TFTP"
	PROTO_AFP                  string = "AFP"
	PROTO_STEALTHNET           string = "STEALTHNET"
	PROTO_AIMINI               string = "AIMINI"
	PROTO_SIP                  string = "SIP"
	PROTO_TRUPHONE             string = "TRUPHONE"
	PROTO_IP_ICMPV6            string = "IP_ICMPV6"
	PROTO_DHCPV6               string = "DHCPV6"
	PROTO_ARMAGETRON           string = "ARMAGETRON"
	PROTO_CROSSFIRE            string = "CROSSFIRE"
	PROTO_DOFUS                string = "DOFUS"
	PROTO_FIESTA               string = "FIESTA"
	PROTO_FLORENSIA            string = "FLORENSIA"
	PROTO_GUILDWARS            string = "GUILDWARS"
	PROTO_AMAZON_ALEXA         string = "AMAZON_ALEXA"
	PROTO_KERBEROS             string = "KERBEROS"
	PROTO_LDAP                 string = "LDAP"
	PROTO_MAPLESTORY           string = "MAPLESTORY"
	PROTO_MSSQL_TDS            string = "MSSQL_TDS"
	PROTO_PPTP                 string = "PPTP"
	PROTO_WARCRAFT3            string = "WARCRAFT3"
	PROTO_WORLD_OF_KUNG_FU     string = "WORLD_OF_KUNG_FU"
	PROTO_SLACK                string = "SLACK"
	PROTO_FACEBOOK             string = "FACEBOOK"
	PROTO_TWITTER              string = "TWITTER"
	PROTO_DROPBOX              string = "DROPBOX"
	PROTO_GMAIL                string = "GMAIL"
	PROTO_GOOGLE_MAPS          string = "GOOGLE_MAPS"
	PROTO_YOUTUBE              string = "YOUTUBE"
	PROTO_SKYPE_TEAMS          string = "SKYPE_TEAMS"
	PROTO_GOOGLE               string = "GOOGLE"
	PROTO_RPC                  string = "RPC"
	PROTO_NETFLOW              string = "NETFLOW"
	PROTO_SFLOW                string = "SFLOW"
	PROTO_HTTP_CONNECT         string = "HTTP_CONNECT"
	PROTO_HTTP_PROXY           string = "HTTP_PROXY"
	PROTO_CITRIX               string = "CITRIX"
	PROTO_NETFLIX              string = "NETFLIX"
	PROTO_LASTFM               string = "LASTFM"
	PROTO_WAZE                 string = "WAZE"
	PROTO_YOUTUBE_UPLOAD       string = "YOUTUBE_UPLOAD"
	PROTO_HULU                 string = "HULU"
	PROTO_CHECKMK              string = "CHECKMK"
	PROTO_AJP                  string = "AJP"
	PROTO_APPLE                string = "APPLE"
	PROTO_WEBEX                string = "WEBEX"
	PROTO_WHATSAPP             string = "WHATSAPP"
	PROTO_APPLE_ICLOUD         string = "APPLE_ICLOUD"
	PROTO_VIBER                string = "VIBER"
	PROTO_APPLE_ITUNES         string = "APPLE_ITUNES"
	PROTO_RADIUS               string = "RADIUS"
	PROTO_WINDOWS_UPDATE       string = "WINDOWS_UPDATE"
	PROTO_TEAMVIEWER           string = "TEAMVIEWER"
	PROTO_TUENTI               string = "TUENTI"
	PROTO_LOTUS_NOTES          string = "LOTUS_NOTES"
	PROTO_SAP                  string = "SAP"
	PROTO_GTP                  string = "GTP"
	PROTO_WSD                  string = "WSD"
	PROTO_LLMNR                string = "LLMNR"
	PROTO_TOCA_BOCA            string = "TOCA_BOCA"
	PROTO_SPOTIFY              string = "SPOTIFY"
	PROTO_MESSENGER            string = "MESSENGER"
	PROTO_H323                 string = "H323"
	PROTO_OPENVPN              string = "OPENVPN"
	PROTO_NOE                  string = "NOE"
	PROTO_CISCOVPN             string = "CISCOVPN"
	PROTO_TEAMSPEAK            string = "TEAMSPEAK"
	PROTO_TOR                  string = "TOR"
	PROTO_SKINNY               string = "SKINNY"
	PROTO_RTCP                 string = "RTCP"
	PROTO_RSYNC                string = "RSYNC"
	PROTO_ORACLE               string = "ORACLE"
	PROTO_CORBA                string = "CORBA"
	PROTO_UBUNTUONE            string = "UBUNTUONE"
	PROTO_WHOIS_DAS            string = "WHOIS_DAS"
	PROTO_SD_RTN               string = "SD_RTN"
	PROTO_SOCKS                string = "SOCKS"
	PROTO_NINTENDO             string = "NINTENDO"
	PROTO_RTMP                 string = "RTMP"
	PROTO_FTP_DATA             string = "FTP_DATA"
	PROTO_WIKIPEDIA            string = "WIKIPEDIA"
	PROTO_ZMQ                  string = "ZMQ"
	PROTO_AMAZON               string = "AMAZON"
	PROTO_EBAY                 string = "EBAY"
	PROTO_CNN                  string = "CNN"
	PROTO_MEGACO               string = "MEGACO"
	PROTO_REDIS                string = "REDIS"
	PROTO_PINTEREST            string = "PINTEREST"
	PROTO_VHUA                 string = "VHUA"
	PROTO_TELEGRAM             string = "TELEGRAM"
	PROTO_VEVO                 string = "VEVO"
	PROTO_PANDORA              string = "PANDORA"
	PROTO_QUIC                 string = "QUIC"
	PROTO_ZOOM                 string = "ZOOM"
	PROTO_EAQ                  string = "EAQ"
	PROTO_OOKLA                string = "OOKLA"
	PROTO_AMQP                 string = "AMQP"
	PROTO_KAKAOTALK            string = "KAKAOTALK"
	PROTO_KAKAOTALK_VOICE      string = "KAKAOTALK_VOICE"
	PROTO_TWITCH               string = "TWITCH"
	PROTO_DOH_DOT              string = "DOH_DOT"
	PROTO_WECHAT               string = "WECHAT"
	PROTO_MPEGTS               string = "MPEGTS"
	PROTO_SNAPCHAT             string = "SNAPCHAT"
	PROTO_SINA                 string = "SINA"
	PROTO_HANGOUT_DUO          string = "HANGOUT_DUO"
	PROTO_IFLIX                string = "IFLIX"
	PROTO_GITHUB               string = "GITHUB"
	PROTO_BJNP                 string = "BJNP"
	PROTO_REDDIT               string = "REDDIT"
	PROTO_WIREGUARD            string = "WIREGUARD"
	PROTO_SMPP                 string = "SMPP"
	PROTO_DNSCRYPT             string = "DNSCRYPT"
	PROTO_TINC                 string = "TINC"
	PROTO_DEEZER               string = "DEEZER"
	PROTO_INSTAGRAM            string = "INSTAGRAM"
	PROTO_MICROSOFT            string = "MICROSOFT"
	PROTO_STARCRAFT            string = "STARCRAFT"
	PROTO_TEREDO               string = "TEREDO"
	PROTO_HOTSPOT_SHIELD       string = "HOTSPOT_SHIELD"
	PROTO_IMO                  string = "IMO"
	PROTO_GOOGLE_DRIVE         string = "GOOGLE_DRIVE"
	PROTO_OCS                  string = "OCS"
	PROTO_MICROSOFT_365        string = "MICROSOFT_365"
	PROTO_CLOUDFLARE           string = "CLOUDFLARE"
	PROTO_MS_ONE_DRIVE         string = "MS_ONE_DRIVE"
	PROTO_MQTT                 string = "MQTT"
	PROTO_RX                   string = "RX"
	PROTO_APPLESTORE           string = "APPLESTORE"
	PROTO_OPENDNS              string = "OPENDNS"
	PROTO_GIT                  string = "GIT"
	PROTO_DRDA                 string = "DRDA"
	PROTO_PLAYSTORE            string = "PLAYSTORE"
	PROTO_SOMEIP               string = "SOMEIP"
	PROTO_FIX                  string = "FIX"
	PROTO_PLAYSTATION          string = "PLAYSTATION"
	PROTO_PASTEBIN             string = "PASTEBIN"
	PROTO_LINKEDIN             string = "LINKEDIN"
	PROTO_SOUNDCLOUD           string = "SOUNDCLOUD"
	PROTO_CSGO                 string = "CSGO"
	PROTO_LISP                 string = "LISP"
	PROTO_DIAMETER             string = "DIAMETER"
	PROTO_APPLE_PUSH           string = "APPLE_PUSH"
	PROTO_GOOGLE_SERVICES      string = "GOOGLE_SERVICES"
	PROTO_AMAZON_VIDEO         string = "AMAZON_VIDEO"
	PROTO_GOOGLE_DOCS          string = "GOOGLE_DOCS"
	PROTO_WHATSAPP_FILES       string = "WHATSAPP_FILES"
	PROTO_TARGUS_GETDATA       string = "TARGUS_GETDATA"
	PROTO_DNP3                 string = "DNP3"
	PROTO_IEC60870             string = "IEC60870"
	PROTO_BLOOMBERG            string = "BLOOMBERG"
	PROTO_CAPWAP               string = "CAPWAP"
	PROTO_ZABBIX               string = "ZABBIX"
	PROTO_S7COMM               string = "S7COMM"
	PROTO_MSTEAMS              string = "MSTEAMS"
	PROTO_WEBSOCKET            string = "WEBSOCKET"
	PROTO_ANYDESK              string = "ANYDESK"
	PROTO_SOAP                 string = "SOAP"
	PROTO_APPLE_SIRI           string = "APPLE_SIRI"
	PROTO_SNAPCHAT_CALL        string = "SNAPCHAT_CALL"
	PROTO_HPVIRTGRP            string = "HPVIRTGRP"
	PROTO_GENSHIN_IMPACT       string = "GENSHIN_IMPACT"
	PROTO_ACTIVISION           string = "ACTIVISION"
	PROTO_FORTICLIENT          string = "FORTICLIENT"
	PROTO_Z3950                string = "Z3950"
	PROTO_LIKEE                string = "LIKEE"
	PROTO_GITLAB               string = "GITLAB"
	PROTO_AVAST_SECUREDNS      string = "AVAST_SECUREDNS"
	PROTO_CASSANDRA            string = "CASSANDRA"
	PROTO_AMAZON_AWS           string = "AMAZON_AWS"
	PROTO_SALESFORCE           string = "SALESFORCE"
	PROTO_VIMEO                string = "VIMEO"
	PROTO_FACEBOOK_VOIP        string = "FACEBOOK_VOIP"
	PROTO_SIGNAL_VOIP          string = "SIGNAL_VOIP"
	PROTO_FUZE                 string = "FUZE"
	PROTO_GTP_U                string = "GTP_U"
	PROTO_GTP_C                string = "GTP_C"
	PROTO_GTP_PRIME            string = "GTP_PRIME"
	PROTO_ALIBABA              string = "ALIBABA"
	PROTO_CRASHLYSTICS         string = "CRASHLYSTICS"
	PROTO_MICROSOFT_AZURE      string = "MICROSOFT_AZURE"
	PROTO_ICLOUD_PRIVATE_RELAY string = "ICLOUD_PRIVATE_RELAY"
	PROTO_ETHERNET_IP          string = "ETHERNET_IP"
	PROTO_BADOO                string = "BADOO"
	PROTO_ACCUWEATHER          string = "ACCUWEATHER"
	PROTO_GOOGLE_CLASSROOM     string = "GOOGLE_CLASSROOM"
	PROTO_HSRP                 string = "HSRP"
	PROTO_CYBERSECURITY        string = "CYBERSECURITY"
	PROTO_GOOGLE_CLOUD         string = "GOOGLE_CLOUD"
	PROTO_TENCENT              string = "TENCENT"
	PROTO_RAKNET               string = "RAKNET"
	PROTO_XIAOMI               string = "XIAOMI"
	PROTO_EDGECAST             string = "EDGECAST"
	PROTO_CACHEFLY             string = "CACHEFLY"
	PROTO_SOFTETHER            string = "SOFTETHER"
	PROTO_MPEGDASH             string = "MPEGDASH"
	PROTO_DAZN                 string = "DAZN"
	PROTO_GOTO                 string = "GOTO"
	PROTO_RSH                  string = "RSH"
	PROTO_1KXUN                string = "1KXUN"
	PROTO_IP_PGM               string = "IP_PGM"
	PROTO_IP_PIM               string = "IP_PIM"
	PROTO_COLLECTD             string = "COLLECTD"
	PROTO_TUNNELBEAR           string = "TUNNELBEAR"
	PROTO_CLOUDFLARE_WARP      string = "CLOUDFLARE_WARP"
	PROTO_I3D                  string = "I3D"
	PROTO_RIOTGAMES            string = "RIOTGAMES"
	PROTO_PSIPHON              string = "PSIPHON"
	PROTO_ULTRASURF            string = "ULTRASURF"
)

var NdpiProtocolNameMap = map[NdpiProtocol]string{
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
