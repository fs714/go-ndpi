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
	NDPI_PROTOCOL_VK                   NdpiProtocol = 22
	NDPI_PROTOCOL_MAIL_POPS            NdpiProtocol = 23
	NDPI_PROTOCOL_TAILSCALE            NdpiProtocol = 24
	NDPI_PROTOCOL_YANDEX               NdpiProtocol = 25
	NDPI_PROTOCOL_NTOP                 NdpiProtocol = 26
	NDPI_PROTOCOL_COAP                 NdpiProtocol = 27
	NDPI_PROTOCOL_VMWARE               NdpiProtocol = 28
	NDPI_PROTOCOL_MAIL_SMTPS           NdpiProtocol = 29
	NDPI_PROTOCOL_DTLS                 NdpiProtocol = 30
	NDPI_PROTOCOL_UBNTAC2              NdpiProtocol = 31
	NDPI_PROTOCOL_BFCP                 NdpiProtocol = 32
	NDPI_PROTOCOL_YANDEX_MAIL          NdpiProtocol = 33
	NDPI_PROTOCOL_YANDEX_MUSIC         NdpiProtocol = 34
	NDPI_PROTOCOL_GNUTELLA             NdpiProtocol = 35
	NDPI_PROTOCOL_EDONKEY              NdpiProtocol = 36
	NDPI_PROTOCOL_BITTORRENT           NdpiProtocol = 37
	NDPI_PROTOCOL_MSTEAMS_CALL         NdpiProtocol = 38
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
	NDPI_PROTOCOL_IQIYI                NdpiProtocol = 54
	NDPI_PROTOCOL_MESHTASTIC           NdpiProtocol = 55
	NDPI_PROTOCOL_YANDEX_MARKET        NdpiProtocol = 56
	NDPI_PROTOCOL_YANDEX_DISK          NdpiProtocol = 57
	NDPI_PROTOCOL_DISCORD              NdpiProtocol = 58
	NDPI_PROTOCOL_ADOBE_CONNECT        NdpiProtocol = 59
	NDPI_PROTOCOL_MONGODB              NdpiProtocol = 60
	NDPI_PROTOCOL_PLURALSIGHT          NdpiProtocol = 61
	NDPI_PROTOCOL_YANDEX_CLOUD         NdpiProtocol = 62
	NDPI_PROTOCOL_OCSP                 NdpiProtocol = 63
	NDPI_PROTOCOL_VXLAN                NdpiProtocol = 64
	NDPI_PROTOCOL_IRC                  NdpiProtocol = 65
	NDPI_PROTOCOL_MERAKI_CLOUD         NdpiProtocol = 66
	NDPI_PROTOCOL_JABBER               NdpiProtocol = 67
	NDPI_PROTOCOL_NATS                 NdpiProtocol = 68
	NDPI_PROTOCOL_AMONG_US             NdpiProtocol = 69
	NDPI_PROTOCOL_YAHOO                NdpiProtocol = 70
	NDPI_PROTOCOL_DISNEYPLUS           NdpiProtocol = 71
	NDPI_PROTOCOL_HART_IP              NdpiProtocol = 72
	NDPI_PROTOCOL_IP_VRRP              NdpiProtocol = 73
	NDPI_PROTOCOL_STEAM                NdpiProtocol = 74
	NDPI_PROTOCOL_MELSEC               NdpiProtocol = 75
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
	NDPI_PROTOCOL_YANDEX_METRIKA       NdpiProtocol = 98
	NDPI_PROTOCOL_YANDEX_DIRECT        NdpiProtocol = 99
	NDPI_PROTOCOL_SIP                  NdpiProtocol = 100
	NDPI_PROTOCOL_TRUPHONE             NdpiProtocol = 101
	NDPI_PROTOCOL_IP_ICMPV6            NdpiProtocol = 102
	NDPI_PROTOCOL_DHCPV6               NdpiProtocol = 103
	NDPI_PROTOCOL_ARMAGETRON           NdpiProtocol = 104
	NDPI_PROTOCOL_CROSSFIRE            NdpiProtocol = 105
	NDPI_PROTOCOL_DOFUS                NdpiProtocol = 106
	NDPI_PROTOCOL_BLACKNUT             NdpiProtocol = 107
	NDPI_PROTOCOL_BOOSTEROID           NdpiProtocol = 108
	NDPI_PROTOCOL_GUILDWARS2           NdpiProtocol = 109
	NDPI_PROTOCOL_AMAZON_ALEXA         NdpiProtocol = 110
	NDPI_PROTOCOL_KERBEROS             NdpiProtocol = 111
	NDPI_PROTOCOL_LDAP                 NdpiProtocol = 112
	NDPI_PROTOCOL_NEXON                NdpiProtocol = 113
	NDPI_PROTOCOL_MSSQL_TDS            NdpiProtocol = 114
	NDPI_PROTOCOL_PPTP                 NdpiProtocol = 115
	NDPI_PROTOCOL_IP_AH                NdpiProtocol = 116
	NDPI_PROTOCOL_IP_ESP               NdpiProtocol = 117
	NDPI_PROTOCOL_SLACK                NdpiProtocol = 118
	NDPI_PROTOCOL_FACEBOOK             NdpiProtocol = 119
	NDPI_PROTOCOL_TWITTER              NdpiProtocol = 120
	NDPI_PROTOCOL_DROPBOX              NdpiProtocol = 121
	NDPI_PROTOCOL_GMAIL                NdpiProtocol = 122
	NDPI_PROTOCOL_GOOGLE_MAPS          NdpiProtocol = 123
	NDPI_PROTOCOL_YOUTUBE              NdpiProtocol = 124
	NDPI_PROTOCOL_MOZILLA              NdpiProtocol = 125
	NDPI_PROTOCOL_GOOGLE               NdpiProtocol = 126
	NDPI_PROTOCOL_MS_RPCH              NdpiProtocol = 127
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
	NDPI_PROTOCOL_EGD                  NdpiProtocol = 149
	NDPI_PROTOCOL_HCL_NOTES            NdpiProtocol = 150
	NDPI_PROTOCOL_SAP                  NdpiProtocol = 151
	NDPI_PROTOCOL_GTP                  NdpiProtocol = 152
	NDPI_PROTOCOL_WSD                  NdpiProtocol = 153
	NDPI_PROTOCOL_LLMNR                NdpiProtocol = 154
	NDPI_PROTOCOL_TOCA_BOCA            NdpiProtocol = 155
	NDPI_PROTOCOL_SPOTIFY              NdpiProtocol = 156
	NDPI_PROTOCOL_FACEBOOK_MESSENGER   NdpiProtocol = 157
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
	NDPI_PROTOCOL_CANONICAL            NdpiProtocol = 169
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
	NDPI_PROTOCOL_RESP                 NdpiProtocol = 182
	NDPI_PROTOCOL_PINTEREST            NdpiProtocol = 183
	NDPI_PROTOCOL_OSPF                 NdpiProtocol = 184
	NDPI_PROTOCOL_TELEGRAM             NdpiProtocol = 185
	NDPI_PROTOCOL_COD_MOBILE           NdpiProtocol = 186
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
	NDPI_PROTOCOL_GOOGLE_MEET          NdpiProtocol = 201
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
	NDPI_PROTOCOL_BLIZZARD             NdpiProtocol = 213
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
	NDPI_PROTOCOL_VALVE_SDR            NdpiProtocol = 235
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
	NDPI_PROTOCOL_THREEMA              NdpiProtocol = 305
	NDPI_PROTOCOL_ALICLOUD             NdpiProtocol = 306
	NDPI_PROTOCOL_AVAST                NdpiProtocol = 307
	NDPI_PROTOCOL_TIVOCONNECT          NdpiProtocol = 308
	NDPI_PROTOCOL_KISMET               NdpiProtocol = 309
	NDPI_PROTOCOL_FASTCGI              NdpiProtocol = 310
	NDPI_PROTOCOL_FTPS                 NdpiProtocol = 311
	NDPI_PROTOCOL_NATPMP               NdpiProtocol = 312
	NDPI_PROTOCOL_SYNCTHING            NdpiProtocol = 313
	NDPI_PROTOCOL_CRYNET               NdpiProtocol = 314
	NDPI_PROTOCOL_LINE                 NdpiProtocol = 315
	NDPI_PROTOCOL_LINE_CALL            NdpiProtocol = 316
	NDPI_PROTOCOL_APPLETVPLUS          NdpiProtocol = 317
	NDPI_PROTOCOL_DIRECTV              NdpiProtocol = 318
	NDPI_PROTOCOL_HBO                  NdpiProtocol = 319
	NDPI_PROTOCOL_VUDU                 NdpiProtocol = 320
	NDPI_PROTOCOL_SHOWTIME             NdpiProtocol = 321
	NDPI_PROTOCOL_DAILYMOTION          NdpiProtocol = 322
	NDPI_PROTOCOL_LIVESTREAM           NdpiProtocol = 323
	NDPI_PROTOCOL_TENCENTVIDEO         NdpiProtocol = 324
	NDPI_PROTOCOL_IHEARTRADIO          NdpiProtocol = 325
	NDPI_PROTOCOL_TIDAL                NdpiProtocol = 326
	NDPI_PROTOCOL_TUNEIN               NdpiProtocol = 327
	NDPI_PROTOCOL_SIRIUSXMRADIO        NdpiProtocol = 328
	NDPI_PROTOCOL_MUNIN                NdpiProtocol = 329
	NDPI_PROTOCOL_ELASTICSEARCH        NdpiProtocol = 330
	NDPI_PROTOCOL_TUYA_LP              NdpiProtocol = 331
	NDPI_PROTOCOL_TPLINK_SHP           NdpiProtocol = 332
	NDPI_PROTOCOL_SOURCE_ENGINE        NdpiProtocol = 333
	NDPI_PROTOCOL_BACNET               NdpiProtocol = 334
	NDPI_PROTOCOL_OICQ                 NdpiProtocol = 335
	NDPI_PROTOCOL_HOTS                 NdpiProtocol = 336
	NDPI_PROTOCOL_FACEBOOK_REEL_STORY  NdpiProtocol = 337
	NDPI_PROTOCOL_SRTP                 NdpiProtocol = 338
	NDPI_PROTOCOL_OPERA_VPN            NdpiProtocol = 339
	NDPI_PROTOCOL_EPICGAMES            NdpiProtocol = 340
	NDPI_PROTOCOL_GEFORCENOW           NdpiProtocol = 341
	NDPI_PROTOCOL_NVIDIA               NdpiProtocol = 342
	NDPI_PROTOCOL_BITCOIN              NdpiProtocol = 343
	NDPI_PROTOCOL_PROTONVPN            NdpiProtocol = 344
	NDPI_PROTOCOL_APACHE_THRIFT        NdpiProtocol = 345
	NDPI_PROTOCOL_ROBLOX               NdpiProtocol = 346
	NDPI_PROTOCOL_SERVICE_LOCATION     NdpiProtocol = 347
	NDPI_PROTOCOL_MULLVAD              NdpiProtocol = 348
	NDPI_PROTOCOL_HTTP2                NdpiProtocol = 349
	NDPI_PROTOCOL_HAPROXY              NdpiProtocol = 350
	NDPI_PROTOCOL_RMCP                 NdpiProtocol = 351
	NDPI_PROTOCOL_CAN                  NdpiProtocol = 352
	NDPI_PROTOCOL_PROTOBUF             NdpiProtocol = 353
	NDPI_PROTOCOL_ETHEREUM             NdpiProtocol = 354
	NDPI_PROTOCOL_TELEGRAM_VOIP        NdpiProtocol = 355
	NDPI_PROTOCOL_SINA_WEIBO           NdpiProtocol = 356
	NDPI_PROTOCOL_TESLA_SERVICES       NdpiProtocol = 357
	NDPI_PROTOCOL_PTPV2                NdpiProtocol = 358
	NDPI_PROTOCOL_RTPS                 NdpiProtocol = 359
	NDPI_PROTOCOL_OPC_UA               NdpiProtocol = 360
	NDPI_PROTOCOL_S7COMM_PLUS          NdpiProtocol = 361
	NDPI_PROTOCOL_FINS                 NdpiProtocol = 362
	NDPI_PROTOCOL_ETHERSIO             NdpiProtocol = 363
	NDPI_PROTOCOL_UMAS                 NdpiProtocol = 364
	NDPI_PROTOCOL_BECKHOFF_ADS         NdpiProtocol = 365
	NDPI_PROTOCOL_ISO9506_1_MMS        NdpiProtocol = 366
	NDPI_PROTOCOL_IEEE_C37118          NdpiProtocol = 367
	NDPI_PROTOCOL_ETHERSBUS            NdpiProtocol = 368
	NDPI_PROTOCOL_MONERO               NdpiProtocol = 369
	NDPI_PROTOCOL_DCERPC               NdpiProtocol = 370
	NDPI_PROTOCOL_PROFINET_IO          NdpiProtocol = 371
	NDPI_PROTOCOL_HISLIP               NdpiProtocol = 372
	NDPI_PROTOCOL_UFTP                 NdpiProtocol = 373
	NDPI_PROTOCOL_OPENFLOW             NdpiProtocol = 374
	NDPI_PROTOCOL_JSON_RPC             NdpiProtocol = 375
	NDPI_PROTOCOL_WEBDAV               NdpiProtocol = 376
	NDPI_PROTOCOL_APACHE_KAFKA         NdpiProtocol = 377
	NDPI_PROTOCOL_NOMACHINE            NdpiProtocol = 378
	NDPI_PROTOCOL_IEC62056             NdpiProtocol = 379
	NDPI_PROTOCOL_HL7                  NdpiProtocol = 380
	NDPI_PROTOCOL_CEPH                 NdpiProtocol = 381
	NDPI_PROTOCOL_GOOGLE_CHAT          NdpiProtocol = 382
	NDPI_PROTOCOL_ROUGHTIME            NdpiProtocol = 383
	NDPI_PROTOCOL_PIA                  NdpiProtocol = 384
	NDPI_PROTOCOL_KCP                  NdpiProtocol = 385
	NDPI_PROTOCOL_DOTA2                NdpiProtocol = 386
	NDPI_PROTOCOL_MUMBLE               NdpiProtocol = 387
	NDPI_PROTOCOL_YOJIMBO              NdpiProtocol = 388
	NDPI_PROTOCOL_ELECTRONICARTS       NdpiProtocol = 389
	NDPI_PROTOCOL_STOMP                NdpiProtocol = 390
	NDPI_PROTOCOL_RADMIN               NdpiProtocol = 391
	NDPI_PROTOCOL_RAFT                 NdpiProtocol = 392
	NDPI_PROTOCOL_CIP                  NdpiProtocol = 393
	NDPI_PROTOCOL_GEARMAN              NdpiProtocol = 394
	NDPI_PROTOCOL_TENCENTGAMES         NdpiProtocol = 395
	NDPI_PROTOCOL_GAIJIN               NdpiProtocol = 396
	NDPI_PROTOCOL_C1222                NdpiProtocol = 397
	NDPI_PROTOCOL_HUAWEI               NdpiProtocol = 398
	NDPI_PROTOCOL_HUAWEI_CLOUD         NdpiProtocol = 399
	NDPI_PROTOCOL_DLEP                 NdpiProtocol = 400
	NDPI_PROTOCOL_BFD                  NdpiProtocol = 401
	NDPI_PROTOCOL_NETEASE_GAMES        NdpiProtocol = 402
	NDPI_PROTOCOL_PATHOFEXILE          NdpiProtocol = 403
	NDPI_PROTOCOL_GOOGLE_CALL          NdpiProtocol = 404
	NDPI_PROTOCOL_PFCP                 NdpiProtocol = 405
	NDPI_PROTOCOL_FLUTE                NdpiProtocol = 406
	NDPI_PROTOCOL_LOLWILDRIFT          NdpiProtocol = 407
	NDPI_PROTOCOL_TESO                 NdpiProtocol = 408
	NDPI_PROTOCOL_LDP                  NdpiProtocol = 409
	NDPI_PROTOCOL_KNXNET_IP            NdpiProtocol = 410
	NDPI_PROTOCOL_BLUESKY              NdpiProtocol = 411
	NDPI_PROTOCOL_MASTODON             NdpiProtocol = 412
	NDPI_PROTOCOL_THREADS              NdpiProtocol = 413
	NDPI_PROTOCOL_VIBER_VOIP           NdpiProtocol = 414
	NDPI_PROTOCOL_ZUG                  NdpiProtocol = 415
	NDPI_PROTOCOL_JRMI                 NdpiProtocol = 416
	NDPI_PROTOCOL_RIPE_ATLAS           NdpiProtocol = 417
	NDPI_PROTOCOL_HLS                  NdpiProtocol = 418
	NDPI_PROTOCOL_CLICKHOUSE           NdpiProtocol = 419
	NDPI_PROTOCOL_NANO                 NdpiProtocol = 420
	NDPI_PROTOCOL_OPENWIRE             NdpiProtocol = 421
	NDPI_PROTOCOL_CNP_IP               NdpiProtocol = 422
	NDPI_PROTOCOL_ATG                  NdpiProtocol = 423
	NDPI_PROTOCOL_TRDP                 NdpiProtocol = 424
	NDPI_PROTOCOL_LUSTRE               NdpiProtocol = 425
	NDPI_PROTOCOL_NORDVPN              NdpiProtocol = 426
	NDPI_PROTOCOL_SURFSHARK            NdpiProtocol = 427
	NDPI_PROTOCOL_CACTUSVPN            NdpiProtocol = 428
	NDPI_PROTOCOL_WINDSCRIBE           NdpiProtocol = 429
	NDPI_PROTOCOL_SONOS                NdpiProtocol = 430
	NDPI_PROTOCOL_DINGTALK             NdpiProtocol = 431
	NDPI_PROTOCOL_PALTALK              NdpiProtocol = 432
	NDPI_PROTOCOL_NAVER                NdpiProtocol = 433
	NDPI_PROTOCOL_SHEIN                NdpiProtocol = 434
	NDPI_PROTOCOL_TEMU                 NdpiProtocol = 435
	NDPI_PROTOCOL_TAOBAO               NdpiProtocol = 436
	NDPI_PROTOCOL_MIKROTIK             NdpiProtocol = 437
	NDPI_PROTOCOL_DICOM                NdpiProtocol = 438
	NDPI_PROTOCOL_PARAMOUNTPLUS        NdpiProtocol = 439
	NDPI_PROTOCOL_YANDEX_ALICE         NdpiProtocol = 440
	NDPI_PROTOCOL_VIVOX                NdpiProtocol = 441
	NDPI_PROTOCOL_DIGITALOCEAN         NdpiProtocol = 442
	NDPI_PROTOCOL_RUTUBE               NdpiProtocol = 443
	NDPI_PROTOCOL_LAGOFAST             NdpiProtocol = 444
	NDPI_PROTOCOL_GEARUP_BOOSTER       NdpiProtocol = 445
	NDPI_PROTOCOL_RUMBLE               NdpiProtocol = 446
	NDPI_PROTOCOL_UBIQUITY             NdpiProtocol = 447
	NDPI_PROTOCOL_MSDO                 NdpiProtocol = 448
	NDPI_PROTOCOL_ROCKSTAR_GAMES       NdpiProtocol = 449
	NDPI_PROTOCOL_KICK                 NdpiProtocol = 450
	NDPI_PROTOCOL_HAMACHI              NdpiProtocol = 451
	NDPI_PROTOCOL_GLBP                 NdpiProtocol = 452
	NDPI_PROTOCOL_EASYWEATHER          NdpiProtocol = 453
	NDPI_PROTOCOL_MUDFISH              NdpiProtocol = 454
	NDPI_PROTOCOL_TRISTATION           NdpiProtocol = 455
	NDPI_PROTOCOL_SAMSUNG_SDP          NdpiProtocol = 456
	NDPI_PROTOCOL_MATTER               NdpiProtocol = 457
	NDPI_PROTOCOL_AWS_COGNITO          NdpiProtocol = 458
	NDPI_PROTOCOL_AWS_API_GATEWAY      NdpiProtocol = 459
	NDPI_PROTOCOL_AWS_KINESIS          NdpiProtocol = 460
	NDPI_PROTOCOL_AWS_EC2              NdpiProtocol = 461
	NDPI_PROTOCOL_AWS_EMR              NdpiProtocol = 462
	NDPI_PROTOCOL_AWS_S3               NdpiProtocol = 463
	NDPI_PROTOCOL_AWS_CLOUDFRONT       NdpiProtocol = 464
	NDPI_PROTOCOL_AWS_DYNAMODB         NdpiProtocol = 465
	NDPI_PROTOCOL_ESPN                 NdpiProtocol = 466
	NDPI_PROTOCOL_AKAMAI               NdpiProtocol = 467
	NDPI_PROTOCOL_JSON                 NdpiProtocol = 468
	NDPI_PROTOCOL_MSGPACK              NdpiProtocol = 469
	NDPI_PROTOCOL_GITHUB_COPILOT       NdpiProtocol = 470
	NDPI_PROTOCOL_GITHUB_PACKAGES      NdpiProtocol = 471
	NDPI_PROTOCOL_GITHUB_ACTIONS       NdpiProtocol = 472
	NDPI_PROTOCOL_YOUTUBE_KIDS         NdpiProtocol = 473
	NDPI_PROTOCOL_SBE                  NdpiProtocol = 474
	NDPI_PROTOCOL_IRIS                 NdpiProtocol = 475
	NDPI_PROTOCOL_LIBP2P               NdpiProtocol = 476
	NDPI_PROTOCOL_NETMOTION            NdpiProtocol = 477
	NDPI_PROTOCOL_PROTON               NdpiProtocol = 478
	NDPI_PROTOCOL_PROTON_MAIL          NdpiProtocol = 479
	NDPI_PROTOCOL_PROTON_PASS          NdpiProtocol = 480
	NDPI_PROTOCOL_PROTON_DRIVE         NdpiProtocol = 481
	NDPI_PROTOCOL_PROTON_CALENDAR      NdpiProtocol = 482
	NDPI_PROTOCOL_PROTON_DOCS          NdpiProtocol = 483
	NDPI_PROTOCOL_PROTON_WALLET        NdpiProtocol = 484
	NDPI_PROTOCOL_PROTON_LUMO          NdpiProtocol = 485
	NDPI_PROTOCOL_PROTON_MEET          NdpiProtocol = 486
	NDPI_PROTOCOL_YGGDRASIL            NdpiProtocol = 487
	NDPI_PROTOCOL_NEBULA               NdpiProtocol = 488
	NDPI_PROTOCOL_DISCORD_CALL         NdpiProtocol = 489
	NDPI_PROTOCOL_FREEFIRE             NdpiProtocol = 490
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
	PROTO_VK                   string = "VK"
	PROTO_MAIL_POPS            string = "MAIL_POPS"
	PROTO_TAILSCALE            string = "TAILSCALE"
	PROTO_YANDEX               string = "YANDEX"
	PROTO_NTOP                 string = "NTOP"
	PROTO_COAP                 string = "COAP"
	PROTO_VMWARE               string = "VMWARE"
	PROTO_MAIL_SMTPS           string = "MAIL_SMTPS"
	PROTO_DTLS                 string = "DTLS"
	PROTO_UBNTAC2              string = "UBNTAC2"
	PROTO_BFCP                 string = "BFCP"
	PROTO_YANDEX_MAIL          string = "YANDEX_MAIL"
	PROTO_YANDEX_MUSIC         string = "YANDEX_MUSIC"
	PROTO_GNUTELLA             string = "GNUTELLA"
	PROTO_EDONKEY              string = "EDONKEY"
	PROTO_BITTORRENT           string = "BITTORRENT"
	PROTO_MSTEAMS_CALL         string = "MSTEAMS_CALL"
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
	PROTO_IQIYI                string = "IQIYI"
	PROTO_MESHTASTIC           string = "MESHTASTIC"
	PROTO_YANDEX_MARKET        string = "YANDEX_MARKET"
	PROTO_YANDEX_DISK          string = "YANDEX_DISK"
	PROTO_DISCORD              string = "DISCORD"
	PROTO_ADOBE_CONNECT        string = "ADOBE_CONNECT"
	PROTO_MONGODB              string = "MONGODB"
	PROTO_PLURALSIGHT          string = "PLURALSIGHT"
	PROTO_YANDEX_CLOUD         string = "YANDEX_CLOUD"
	PROTO_OCSP                 string = "OCSP"
	PROTO_VXLAN                string = "VXLAN"
	PROTO_IRC                  string = "IRC"
	PROTO_MERAKI_CLOUD         string = "MERAKI_CLOUD"
	PROTO_JABBER               string = "JABBER"
	PROTO_NATS                 string = "NATS"
	PROTO_AMONG_US             string = "AMONG_US"
	PROTO_YAHOO                string = "YAHOO"
	PROTO_DISNEYPLUS           string = "DISNEYPLUS"
	PROTO_HART_IP              string = "HART_IP"
	PROTO_IP_VRRP              string = "IP_VRRP"
	PROTO_STEAM                string = "STEAM"
	PROTO_MELSEC               string = "MELSEC"
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
	PROTO_YANDEX_METRIKA       string = "YANDEX_METRIKA"
	PROTO_YANDEX_DIRECT        string = "YANDEX_DIRECT"
	PROTO_SIP                  string = "SIP"
	PROTO_TRUPHONE             string = "TRUPHONE"
	PROTO_IP_ICMPV6            string = "IP_ICMPV6"
	PROTO_DHCPV6               string = "DHCPV6"
	PROTO_ARMAGETRON           string = "ARMAGETRON"
	PROTO_CROSSFIRE            string = "CROSSFIRE"
	PROTO_DOFUS                string = "DOFUS"
	PROTO_BLACKNUT             string = "BLACKNUT"
	PROTO_BOOSTEROID           string = "BOOSTEROID"
	PROTO_GUILDWARS2           string = "GUILDWARS2"
	PROTO_AMAZON_ALEXA         string = "AMAZON_ALEXA"
	PROTO_KERBEROS             string = "KERBEROS"
	PROTO_LDAP                 string = "LDAP"
	PROTO_NEXON                string = "NEXON"
	PROTO_MSSQL_TDS            string = "MSSQL_TDS"
	PROTO_PPTP                 string = "PPTP"
	PROTO_IP_AH                string = "IP_AH"
	PROTO_IP_ESP               string = "IP_ESP"
	PROTO_SLACK                string = "SLACK"
	PROTO_FACEBOOK             string = "FACEBOOK"
	PROTO_TWITTER              string = "TWITTER"
	PROTO_DROPBOX              string = "DROPBOX"
	PROTO_GMAIL                string = "GMAIL"
	PROTO_GOOGLE_MAPS          string = "GOOGLE_MAPS"
	PROTO_YOUTUBE              string = "YOUTUBE"
	PROTO_MOZILLA              string = "MOZILLA"
	PROTO_GOOGLE               string = "GOOGLE"
	PROTO_MS_RPCH              string = "MS_RPCH"
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
	PROTO_EGD                  string = "EGD"
	PROTO_HCL_NOTES            string = "HCL_NOTES"
	PROTO_SAP                  string = "SAP"
	PROTO_GTP                  string = "GTP"
	PROTO_WSD                  string = "WSD"
	PROTO_LLMNR                string = "LLMNR"
	PROTO_TOCA_BOCA            string = "TOCA_BOCA"
	PROTO_SPOTIFY              string = "SPOTIFY"
	PROTO_FACEBOOK_MESSENGER   string = "FACEBOOK_MESSENGER"
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
	PROTO_CANONICAL            string = "CANONICAL"
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
	PROTO_RESP                 string = "RESP"
	PROTO_PINTEREST            string = "PINTEREST"
	PROTO_OSPF                 string = "OSPF"
	PROTO_TELEGRAM             string = "TELEGRAM"
	PROTO_COD_MOBILE           string = "COD_MOBILE"
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
	PROTO_GOOGLE_MEET          string = "GOOGLE_MEET"
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
	PROTO_BLIZZARD             string = "BLIZZARD"
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
	PROTO_VALVE_SDR            string = "VALVE_SDR"
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
	PROTO_THREEMA              string = "THREEMA"
	PROTO_ALICLOUD             string = "ALICLOUD"
	PROTO_AVAST                string = "AVAST"
	PROTO_TIVOCONNECT          string = "TIVOCONNECT"
	PROTO_KISMET               string = "KISMET"
	PROTO_FASTCGI              string = "FASTCGI"
	PROTO_FTPS                 string = "FTPS"
	PROTO_NATPMP               string = "NATPMP"
	PROTO_SYNCTHING            string = "SYNCTHING"
	PROTO_CRYNET               string = "CRYNET"
	PROTO_LINE                 string = "LINE"
	PROTO_LINE_CALL            string = "LINE_CALL"
	PROTO_APPLETVPLUS          string = "APPLETVPLUS"
	PROTO_DIRECTV              string = "DIRECTV"
	PROTO_HBO                  string = "HBO"
	PROTO_VUDU                 string = "VUDU"
	PROTO_SHOWTIME             string = "SHOWTIME"
	PROTO_DAILYMOTION          string = "DAILYMOTION"
	PROTO_LIVESTREAM           string = "LIVESTREAM"
	PROTO_TENCENTVIDEO         string = "TENCENTVIDEO"
	PROTO_IHEARTRADIO          string = "IHEARTRADIO"
	PROTO_TIDAL                string = "TIDAL"
	PROTO_TUNEIN               string = "TUNEIN"
	PROTO_SIRIUSXMRADIO        string = "SIRIUSXMRADIO"
	PROTO_MUNIN                string = "MUNIN"
	PROTO_ELASTICSEARCH        string = "ELASTICSEARCH"
	PROTO_TUYA_LP              string = "TUYA_LP"
	PROTO_TPLINK_SHP           string = "TPLINK_SHP"
	PROTO_SOURCE_ENGINE        string = "SOURCE_ENGINE"
	PROTO_BACNET               string = "BACNET"
	PROTO_OICQ                 string = "OICQ"
	PROTO_HOTS                 string = "HOTS"
	PROTO_FACEBOOK_REEL_STORY  string = "FACEBOOK_REEL_STORY"
	PROTO_SRTP                 string = "SRTP"
	PROTO_OPERA_VPN            string = "OPERA_VPN"
	PROTO_EPICGAMES            string = "EPICGAMES"
	PROTO_GEFORCENOW           string = "GEFORCENOW"
	PROTO_NVIDIA               string = "NVIDIA"
	PROTO_BITCOIN              string = "BITCOIN"
	PROTO_PROTONVPN            string = "PROTONVPN"
	PROTO_APACHE_THRIFT        string = "APACHE_THRIFT"
	PROTO_ROBLOX               string = "ROBLOX"
	PROTO_SERVICE_LOCATION     string = "SERVICE_LOCATION"
	PROTO_MULLVAD              string = "MULLVAD"
	PROTO_HTTP2                string = "HTTP2"
	PROTO_HAPROXY              string = "HAPROXY"
	PROTO_RMCP                 string = "RMCP"
	PROTO_CAN                  string = "CAN"
	PROTO_PROTOBUF             string = "PROTOBUF"
	PROTO_ETHEREUM             string = "ETHEREUM"
	PROTO_TELEGRAM_VOIP        string = "TELEGRAM_VOIP"
	PROTO_SINA_WEIBO           string = "SINA_WEIBO"
	PROTO_TESLA_SERVICES       string = "TESLA_SERVICES"
	PROTO_PTPV2                string = "PTPV2"
	PROTO_RTPS                 string = "RTPS"
	PROTO_OPC_UA               string = "OPC_UA"
	PROTO_S7COMM_PLUS          string = "S7COMM_PLUS"
	PROTO_FINS                 string = "FINS"
	PROTO_ETHERSIO             string = "ETHERSIO"
	PROTO_UMAS                 string = "UMAS"
	PROTO_BECKHOFF_ADS         string = "BECKHOFF_ADS"
	PROTO_ISO9506_1_MMS        string = "ISO9506_1_MMS"
	PROTO_IEEE_C37118          string = "IEEE_C37118"
	PROTO_ETHERSBUS            string = "ETHERSBUS"
	PROTO_MONERO               string = "MONERO"
	PROTO_DCERPC               string = "DCERPC"
	PROTO_PROFINET_IO          string = "PROFINET_IO"
	PROTO_HISLIP               string = "HISLIP"
	PROTO_UFTP                 string = "UFTP"
	PROTO_OPENFLOW             string = "OPENFLOW"
	PROTO_JSON_RPC             string = "JSON_RPC"
	PROTO_WEBDAV               string = "WEBDAV"
	PROTO_APACHE_KAFKA         string = "APACHE_KAFKA"
	PROTO_NOMACHINE            string = "NOMACHINE"
	PROTO_IEC62056             string = "IEC62056"
	PROTO_HL7                  string = "HL7"
	PROTO_CEPH                 string = "CEPH"
	PROTO_GOOGLE_CHAT          string = "GOOGLE_CHAT"
	PROTO_ROUGHTIME            string = "ROUGHTIME"
	PROTO_PIA                  string = "PIA"
	PROTO_KCP                  string = "KCP"
	PROTO_DOTA2                string = "DOTA2"
	PROTO_MUMBLE               string = "MUMBLE"
	PROTO_YOJIMBO              string = "YOJIMBO"
	PROTO_ELECTRONICARTS       string = "ELECTRONICARTS"
	PROTO_STOMP                string = "STOMP"
	PROTO_RADMIN               string = "RADMIN"
	PROTO_RAFT                 string = "RAFT"
	PROTO_CIP                  string = "CIP"
	PROTO_GEARMAN              string = "GEARMAN"
	PROTO_TENCENTGAMES         string = "TENCENTGAMES"
	PROTO_GAIJIN               string = "GAIJIN"
	PROTO_C1222                string = "C1222"
	PROTO_HUAWEI               string = "HUAWEI"
	PROTO_HUAWEI_CLOUD         string = "HUAWEI_CLOUD"
	PROTO_DLEP                 string = "DLEP"
	PROTO_BFD                  string = "BFD"
	PROTO_NETEASE_GAMES        string = "NETEASE_GAMES"
	PROTO_PATHOFEXILE          string = "PATHOFEXILE"
	PROTO_GOOGLE_CALL          string = "GOOGLE_CALL"
	PROTO_PFCP                 string = "PFCP"
	PROTO_FLUTE                string = "FLUTE"
	PROTO_LOLWILDRIFT          string = "LOLWILDRIFT"
	PROTO_TESO                 string = "TESO"
	PROTO_LDP                  string = "LDP"
	PROTO_KNXNET_IP            string = "KNXNET_IP"
	PROTO_BLUESKY              string = "BLUESKY"
	PROTO_MASTODON             string = "MASTODON"
	PROTO_THREADS              string = "THREADS"
	PROTO_VIBER_VOIP           string = "VIBER_VOIP"
	PROTO_ZUG                  string = "ZUG"
	PROTO_JRMI                 string = "JRMI"
	PROTO_RIPE_ATLAS           string = "RIPE_ATLAS"
	PROTO_HLS                  string = "HLS"
	PROTO_CLICKHOUSE           string = "CLICKHOUSE"
	PROTO_NANO                 string = "NANO"
	PROTO_OPENWIRE             string = "OPENWIRE"
	PROTO_CNP_IP               string = "CNP_IP"
	PROTO_ATG                  string = "ATG"
	PROTO_TRDP                 string = "TRDP"
	PROTO_LUSTRE               string = "LUSTRE"
	PROTO_NORDVPN              string = "NORDVPN"
	PROTO_SURFSHARK            string = "SURFSHARK"
	PROTO_CACTUSVPN            string = "CACTUSVPN"
	PROTO_WINDSCRIBE           string = "WINDSCRIBE"
	PROTO_SONOS                string = "SONOS"
	PROTO_DINGTALK             string = "DINGTALK"
	PROTO_PALTALK              string = "PALTALK"
	PROTO_NAVER                string = "NAVER"
	PROTO_SHEIN                string = "SHEIN"
	PROTO_TEMU                 string = "TEMU"
	PROTO_TAOBAO               string = "TAOBAO"
	PROTO_MIKROTIK             string = "MIKROTIK"
	PROTO_DICOM                string = "DICOM"
	PROTO_PARAMOUNTPLUS        string = "PARAMOUNTPLUS"
	PROTO_YANDEX_ALICE         string = "YANDEX_ALICE"
	PROTO_VIVOX                string = "VIVOX"
	PROTO_DIGITALOCEAN         string = "DIGITALOCEAN"
	PROTO_RUTUBE               string = "RUTUBE"
	PROTO_LAGOFAST             string = "LAGOFAST"
	PROTO_GEARUP_BOOSTER       string = "GEARUP_BOOSTER"
	PROTO_RUMBLE               string = "RUMBLE"
	PROTO_UBIQUITY             string = "UBIQUITY"
	PROTO_MSDO                 string = "MSDO"
	PROTO_ROCKSTAR_GAMES       string = "ROCKSTAR_GAMES"
	PROTO_KICK                 string = "KICK"
	PROTO_HAMACHI              string = "HAMACHI"
	PROTO_GLBP                 string = "GLBP"
	PROTO_EASYWEATHER          string = "EASYWEATHER"
	PROTO_MUDFISH              string = "MUDFISH"
	PROTO_TRISTATION           string = "TRISTATION"
	PROTO_SAMSUNG_SDP          string = "SAMSUNG_SDP"
	PROTO_MATTER               string = "MATTER"
	PROTO_AWS_COGNITO          string = "AWS_COGNITO"
	PROTO_AWS_API_GATEWAY      string = "AWS_API_GATEWAY"
	PROTO_AWS_KINESIS          string = "AWS_KINESIS"
	PROTO_AWS_EC2              string = "AWS_EC2"
	PROTO_AWS_EMR              string = "AWS_EMR"
	PROTO_AWS_S3               string = "AWS_S3"
	PROTO_AWS_CLOUDFRONT       string = "AWS_CLOUDFRONT"
	PROTO_AWS_DYNAMODB         string = "AWS_DYNAMODB"
	PROTO_ESPN                 string = "ESPN"
	PROTO_AKAMAI               string = "AKAMAI"
	PROTO_JSON                 string = "JSON"
	PROTO_MSGPACK              string = "MSGPACK"
	PROTO_GITHUB_COPILOT       string = "GITHUB_COPILOT"
	PROTO_GITHUB_PACKAGES      string = "GITHUB_PACKAGES"
	PROTO_GITHUB_ACTIONS       string = "GITHUB_ACTIONS"
	PROTO_YOUTUBE_KIDS         string = "YOUTUBE_KIDS"
	PROTO_SBE                  string = "SBE"
	PROTO_IRIS                 string = "IRIS"
	PROTO_LIBP2P               string = "LIBP2P"
	PROTO_NETMOTION            string = "NETMOTION"
	PROTO_PROTON               string = "PROTON"
	PROTO_PROTON_MAIL          string = "PROTON_MAIL"
	PROTO_PROTON_PASS          string = "PROTON_PASS"
	PROTO_PROTON_DRIVE         string = "PROTON_DRIVE"
	PROTO_PROTON_CALENDAR      string = "PROTON_CALENDAR"
	PROTO_PROTON_DOCS          string = "PROTON_DOCS"
	PROTO_PROTON_WALLET        string = "PROTON_WALLET"
	PROTO_PROTON_LUMO          string = "PROTON_LUMO"
	PROTO_PROTON_MEET          string = "PROTON_MEET"
	PROTO_YGGDRASIL            string = "YGGDRASIL"
	PROTO_NEBULA               string = "NEBULA"
	PROTO_DISCORD_CALL         string = "DISCORD_CALL"
	PROTO_FREEFIRE             string = "FREEFIRE"
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
	NDPI_PROTOCOL_VK:                   PROTO_VK,
	NDPI_PROTOCOL_MAIL_POPS:            PROTO_MAIL_POPS,
	NDPI_PROTOCOL_TAILSCALE:            PROTO_TAILSCALE,
	NDPI_PROTOCOL_YANDEX:               PROTO_YANDEX,
	NDPI_PROTOCOL_NTOP:                 PROTO_NTOP,
	NDPI_PROTOCOL_COAP:                 PROTO_COAP,
	NDPI_PROTOCOL_VMWARE:               PROTO_VMWARE,
	NDPI_PROTOCOL_MAIL_SMTPS:           PROTO_MAIL_SMTPS,
	NDPI_PROTOCOL_DTLS:                 PROTO_DTLS,
	NDPI_PROTOCOL_UBNTAC2:              PROTO_UBNTAC2,
	NDPI_PROTOCOL_BFCP:                 PROTO_BFCP,
	NDPI_PROTOCOL_YANDEX_MAIL:          PROTO_YANDEX_MAIL,
	NDPI_PROTOCOL_YANDEX_MUSIC:         PROTO_YANDEX_MUSIC,
	NDPI_PROTOCOL_GNUTELLA:             PROTO_GNUTELLA,
	NDPI_PROTOCOL_EDONKEY:              PROTO_EDONKEY,
	NDPI_PROTOCOL_BITTORRENT:           PROTO_BITTORRENT,
	NDPI_PROTOCOL_MSTEAMS_CALL:         PROTO_MSTEAMS_CALL,
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
	NDPI_PROTOCOL_IQIYI:                PROTO_IQIYI,
	NDPI_PROTOCOL_MESHTASTIC:           PROTO_MESHTASTIC,
	NDPI_PROTOCOL_YANDEX_MARKET:        PROTO_YANDEX_MARKET,
	NDPI_PROTOCOL_YANDEX_DISK:          PROTO_YANDEX_DISK,
	NDPI_PROTOCOL_DISCORD:              PROTO_DISCORD,
	NDPI_PROTOCOL_ADOBE_CONNECT:        PROTO_ADOBE_CONNECT,
	NDPI_PROTOCOL_MONGODB:              PROTO_MONGODB,
	NDPI_PROTOCOL_PLURALSIGHT:          PROTO_PLURALSIGHT,
	NDPI_PROTOCOL_YANDEX_CLOUD:         PROTO_YANDEX_CLOUD,
	NDPI_PROTOCOL_OCSP:                 PROTO_OCSP,
	NDPI_PROTOCOL_VXLAN:                PROTO_VXLAN,
	NDPI_PROTOCOL_IRC:                  PROTO_IRC,
	NDPI_PROTOCOL_MERAKI_CLOUD:         PROTO_MERAKI_CLOUD,
	NDPI_PROTOCOL_JABBER:               PROTO_JABBER,
	NDPI_PROTOCOL_NATS:                 PROTO_NATS,
	NDPI_PROTOCOL_AMONG_US:             PROTO_AMONG_US,
	NDPI_PROTOCOL_YAHOO:                PROTO_YAHOO,
	NDPI_PROTOCOL_DISNEYPLUS:           PROTO_DISNEYPLUS,
	NDPI_PROTOCOL_HART_IP:              PROTO_HART_IP,
	NDPI_PROTOCOL_IP_VRRP:              PROTO_IP_VRRP,
	NDPI_PROTOCOL_STEAM:                PROTO_STEAM,
	NDPI_PROTOCOL_MELSEC:               PROTO_MELSEC,
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
	NDPI_PROTOCOL_YANDEX_METRIKA:       PROTO_YANDEX_METRIKA,
	NDPI_PROTOCOL_YANDEX_DIRECT:        PROTO_YANDEX_DIRECT,
	NDPI_PROTOCOL_SIP:                  PROTO_SIP,
	NDPI_PROTOCOL_TRUPHONE:             PROTO_TRUPHONE,
	NDPI_PROTOCOL_IP_ICMPV6:            PROTO_IP_ICMPV6,
	NDPI_PROTOCOL_DHCPV6:               PROTO_DHCPV6,
	NDPI_PROTOCOL_ARMAGETRON:           PROTO_ARMAGETRON,
	NDPI_PROTOCOL_CROSSFIRE:            PROTO_CROSSFIRE,
	NDPI_PROTOCOL_DOFUS:                PROTO_DOFUS,
	NDPI_PROTOCOL_BLACKNUT:             PROTO_BLACKNUT,
	NDPI_PROTOCOL_BOOSTEROID:           PROTO_BOOSTEROID,
	NDPI_PROTOCOL_GUILDWARS2:           PROTO_GUILDWARS2,
	NDPI_PROTOCOL_AMAZON_ALEXA:         PROTO_AMAZON_ALEXA,
	NDPI_PROTOCOL_KERBEROS:             PROTO_KERBEROS,
	NDPI_PROTOCOL_LDAP:                 PROTO_LDAP,
	NDPI_PROTOCOL_NEXON:                PROTO_NEXON,
	NDPI_PROTOCOL_MSSQL_TDS:            PROTO_MSSQL_TDS,
	NDPI_PROTOCOL_PPTP:                 PROTO_PPTP,
	NDPI_PROTOCOL_IP_AH:                PROTO_IP_AH,
	NDPI_PROTOCOL_IP_ESP:               PROTO_IP_ESP,
	NDPI_PROTOCOL_SLACK:                PROTO_SLACK,
	NDPI_PROTOCOL_FACEBOOK:             PROTO_FACEBOOK,
	NDPI_PROTOCOL_TWITTER:              PROTO_TWITTER,
	NDPI_PROTOCOL_DROPBOX:              PROTO_DROPBOX,
	NDPI_PROTOCOL_GMAIL:                PROTO_GMAIL,
	NDPI_PROTOCOL_GOOGLE_MAPS:          PROTO_GOOGLE_MAPS,
	NDPI_PROTOCOL_YOUTUBE:              PROTO_YOUTUBE,
	NDPI_PROTOCOL_MOZILLA:              PROTO_MOZILLA,
	NDPI_PROTOCOL_GOOGLE:               PROTO_GOOGLE,
	NDPI_PROTOCOL_MS_RPCH:              PROTO_MS_RPCH,
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
	NDPI_PROTOCOL_EGD:                  PROTO_EGD,
	NDPI_PROTOCOL_HCL_NOTES:            PROTO_HCL_NOTES,
	NDPI_PROTOCOL_SAP:                  PROTO_SAP,
	NDPI_PROTOCOL_GTP:                  PROTO_GTP,
	NDPI_PROTOCOL_WSD:                  PROTO_WSD,
	NDPI_PROTOCOL_LLMNR:                PROTO_LLMNR,
	NDPI_PROTOCOL_TOCA_BOCA:            PROTO_TOCA_BOCA,
	NDPI_PROTOCOL_SPOTIFY:              PROTO_SPOTIFY,
	NDPI_PROTOCOL_FACEBOOK_MESSENGER:   PROTO_FACEBOOK_MESSENGER,
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
	NDPI_PROTOCOL_CANONICAL:            PROTO_CANONICAL,
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
	NDPI_PROTOCOL_RESP:                 PROTO_RESP,
	NDPI_PROTOCOL_PINTEREST:            PROTO_PINTEREST,
	NDPI_PROTOCOL_OSPF:                 PROTO_OSPF,
	NDPI_PROTOCOL_TELEGRAM:             PROTO_TELEGRAM,
	NDPI_PROTOCOL_COD_MOBILE:           PROTO_COD_MOBILE,
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
	NDPI_PROTOCOL_GOOGLE_MEET:          PROTO_GOOGLE_MEET,
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
	NDPI_PROTOCOL_BLIZZARD:             PROTO_BLIZZARD,
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
	NDPI_PROTOCOL_VALVE_SDR:            PROTO_VALVE_SDR,
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
	NDPI_PROTOCOL_THREEMA:              PROTO_THREEMA,
	NDPI_PROTOCOL_ALICLOUD:             PROTO_ALICLOUD,
	NDPI_PROTOCOL_AVAST:                PROTO_AVAST,
	NDPI_PROTOCOL_TIVOCONNECT:          PROTO_TIVOCONNECT,
	NDPI_PROTOCOL_KISMET:               PROTO_KISMET,
	NDPI_PROTOCOL_FASTCGI:              PROTO_FASTCGI,
	NDPI_PROTOCOL_FTPS:                 PROTO_FTPS,
	NDPI_PROTOCOL_NATPMP:               PROTO_NATPMP,
	NDPI_PROTOCOL_SYNCTHING:            PROTO_SYNCTHING,
	NDPI_PROTOCOL_CRYNET:               PROTO_CRYNET,
	NDPI_PROTOCOL_LINE:                 PROTO_LINE,
	NDPI_PROTOCOL_LINE_CALL:            PROTO_LINE_CALL,
	NDPI_PROTOCOL_APPLETVPLUS:          PROTO_APPLETVPLUS,
	NDPI_PROTOCOL_DIRECTV:              PROTO_DIRECTV,
	NDPI_PROTOCOL_HBO:                  PROTO_HBO,
	NDPI_PROTOCOL_VUDU:                 PROTO_VUDU,
	NDPI_PROTOCOL_SHOWTIME:             PROTO_SHOWTIME,
	NDPI_PROTOCOL_DAILYMOTION:          PROTO_DAILYMOTION,
	NDPI_PROTOCOL_LIVESTREAM:           PROTO_LIVESTREAM,
	NDPI_PROTOCOL_TENCENTVIDEO:         PROTO_TENCENTVIDEO,
	NDPI_PROTOCOL_IHEARTRADIO:          PROTO_IHEARTRADIO,
	NDPI_PROTOCOL_TIDAL:                PROTO_TIDAL,
	NDPI_PROTOCOL_TUNEIN:               PROTO_TUNEIN,
	NDPI_PROTOCOL_SIRIUSXMRADIO:        PROTO_SIRIUSXMRADIO,
	NDPI_PROTOCOL_MUNIN:                PROTO_MUNIN,
	NDPI_PROTOCOL_ELASTICSEARCH:        PROTO_ELASTICSEARCH,
	NDPI_PROTOCOL_TUYA_LP:              PROTO_TUYA_LP,
	NDPI_PROTOCOL_TPLINK_SHP:           PROTO_TPLINK_SHP,
	NDPI_PROTOCOL_SOURCE_ENGINE:        PROTO_SOURCE_ENGINE,
	NDPI_PROTOCOL_BACNET:               PROTO_BACNET,
	NDPI_PROTOCOL_OICQ:                 PROTO_OICQ,
	NDPI_PROTOCOL_HOTS:                 PROTO_HOTS,
	NDPI_PROTOCOL_FACEBOOK_REEL_STORY:  PROTO_FACEBOOK_REEL_STORY,
	NDPI_PROTOCOL_SRTP:                 PROTO_SRTP,
	NDPI_PROTOCOL_OPERA_VPN:            PROTO_OPERA_VPN,
	NDPI_PROTOCOL_EPICGAMES:            PROTO_EPICGAMES,
	NDPI_PROTOCOL_GEFORCENOW:           PROTO_GEFORCENOW,
	NDPI_PROTOCOL_NVIDIA:               PROTO_NVIDIA,
	NDPI_PROTOCOL_BITCOIN:              PROTO_BITCOIN,
	NDPI_PROTOCOL_PROTONVPN:            PROTO_PROTONVPN,
	NDPI_PROTOCOL_APACHE_THRIFT:        PROTO_APACHE_THRIFT,
	NDPI_PROTOCOL_ROBLOX:               PROTO_ROBLOX,
	NDPI_PROTOCOL_SERVICE_LOCATION:     PROTO_SERVICE_LOCATION,
	NDPI_PROTOCOL_MULLVAD:              PROTO_MULLVAD,
	NDPI_PROTOCOL_HTTP2:                PROTO_HTTP2,
	NDPI_PROTOCOL_HAPROXY:              PROTO_HAPROXY,
	NDPI_PROTOCOL_RMCP:                 PROTO_RMCP,
	NDPI_PROTOCOL_CAN:                  PROTO_CAN,
	NDPI_PROTOCOL_PROTOBUF:             PROTO_PROTOBUF,
	NDPI_PROTOCOL_ETHEREUM:             PROTO_ETHEREUM,
	NDPI_PROTOCOL_TELEGRAM_VOIP:        PROTO_TELEGRAM_VOIP,
	NDPI_PROTOCOL_SINA_WEIBO:           PROTO_SINA_WEIBO,
	NDPI_PROTOCOL_TESLA_SERVICES:       PROTO_TESLA_SERVICES,
	NDPI_PROTOCOL_PTPV2:                PROTO_PTPV2,
	NDPI_PROTOCOL_RTPS:                 PROTO_RTPS,
	NDPI_PROTOCOL_OPC_UA:               PROTO_OPC_UA,
	NDPI_PROTOCOL_S7COMM_PLUS:          PROTO_S7COMM_PLUS,
	NDPI_PROTOCOL_FINS:                 PROTO_FINS,
	NDPI_PROTOCOL_ETHERSIO:             PROTO_ETHERSIO,
	NDPI_PROTOCOL_UMAS:                 PROTO_UMAS,
	NDPI_PROTOCOL_BECKHOFF_ADS:         PROTO_BECKHOFF_ADS,
	NDPI_PROTOCOL_ISO9506_1_MMS:        PROTO_ISO9506_1_MMS,
	NDPI_PROTOCOL_IEEE_C37118:          PROTO_IEEE_C37118,
	NDPI_PROTOCOL_ETHERSBUS:            PROTO_ETHERSBUS,
	NDPI_PROTOCOL_MONERO:               PROTO_MONERO,
	NDPI_PROTOCOL_DCERPC:               PROTO_DCERPC,
	NDPI_PROTOCOL_PROFINET_IO:          PROTO_PROFINET_IO,
	NDPI_PROTOCOL_HISLIP:               PROTO_HISLIP,
	NDPI_PROTOCOL_UFTP:                 PROTO_UFTP,
	NDPI_PROTOCOL_OPENFLOW:             PROTO_OPENFLOW,
	NDPI_PROTOCOL_JSON_RPC:             PROTO_JSON_RPC,
	NDPI_PROTOCOL_WEBDAV:               PROTO_WEBDAV,
	NDPI_PROTOCOL_APACHE_KAFKA:         PROTO_APACHE_KAFKA,
	NDPI_PROTOCOL_NOMACHINE:            PROTO_NOMACHINE,
	NDPI_PROTOCOL_IEC62056:             PROTO_IEC62056,
	NDPI_PROTOCOL_HL7:                  PROTO_HL7,
	NDPI_PROTOCOL_CEPH:                 PROTO_CEPH,
	NDPI_PROTOCOL_GOOGLE_CHAT:          PROTO_GOOGLE_CHAT,
	NDPI_PROTOCOL_ROUGHTIME:            PROTO_ROUGHTIME,
	NDPI_PROTOCOL_PIA:                  PROTO_PIA,
	NDPI_PROTOCOL_KCP:                  PROTO_KCP,
	NDPI_PROTOCOL_DOTA2:                PROTO_DOTA2,
	NDPI_PROTOCOL_MUMBLE:               PROTO_MUMBLE,
	NDPI_PROTOCOL_YOJIMBO:              PROTO_YOJIMBO,
	NDPI_PROTOCOL_ELECTRONICARTS:       PROTO_ELECTRONICARTS,
	NDPI_PROTOCOL_STOMP:                PROTO_STOMP,
	NDPI_PROTOCOL_RADMIN:               PROTO_RADMIN,
	NDPI_PROTOCOL_RAFT:                 PROTO_RAFT,
	NDPI_PROTOCOL_CIP:                  PROTO_CIP,
	NDPI_PROTOCOL_GEARMAN:              PROTO_GEARMAN,
	NDPI_PROTOCOL_TENCENTGAMES:         PROTO_TENCENTGAMES,
	NDPI_PROTOCOL_GAIJIN:               PROTO_GAIJIN,
	NDPI_PROTOCOL_C1222:                PROTO_C1222,
	NDPI_PROTOCOL_HUAWEI:               PROTO_HUAWEI,
	NDPI_PROTOCOL_HUAWEI_CLOUD:         PROTO_HUAWEI_CLOUD,
	NDPI_PROTOCOL_DLEP:                 PROTO_DLEP,
	NDPI_PROTOCOL_BFD:                  PROTO_BFD,
	NDPI_PROTOCOL_NETEASE_GAMES:        PROTO_NETEASE_GAMES,
	NDPI_PROTOCOL_PATHOFEXILE:          PROTO_PATHOFEXILE,
	NDPI_PROTOCOL_GOOGLE_CALL:          PROTO_GOOGLE_CALL,
	NDPI_PROTOCOL_PFCP:                 PROTO_PFCP,
	NDPI_PROTOCOL_FLUTE:                PROTO_FLUTE,
	NDPI_PROTOCOL_LOLWILDRIFT:          PROTO_LOLWILDRIFT,
	NDPI_PROTOCOL_TESO:                 PROTO_TESO,
	NDPI_PROTOCOL_LDP:                  PROTO_LDP,
	NDPI_PROTOCOL_KNXNET_IP:            PROTO_KNXNET_IP,
	NDPI_PROTOCOL_BLUESKY:              PROTO_BLUESKY,
	NDPI_PROTOCOL_MASTODON:             PROTO_MASTODON,
	NDPI_PROTOCOL_THREADS:              PROTO_THREADS,
	NDPI_PROTOCOL_VIBER_VOIP:           PROTO_VIBER_VOIP,
	NDPI_PROTOCOL_ZUG:                  PROTO_ZUG,
	NDPI_PROTOCOL_JRMI:                 PROTO_JRMI,
	NDPI_PROTOCOL_RIPE_ATLAS:           PROTO_RIPE_ATLAS,
	NDPI_PROTOCOL_HLS:                  PROTO_HLS,
	NDPI_PROTOCOL_CLICKHOUSE:           PROTO_CLICKHOUSE,
	NDPI_PROTOCOL_NANO:                 PROTO_NANO,
	NDPI_PROTOCOL_OPENWIRE:             PROTO_OPENWIRE,
	NDPI_PROTOCOL_CNP_IP:               PROTO_CNP_IP,
	NDPI_PROTOCOL_ATG:                  PROTO_ATG,
	NDPI_PROTOCOL_TRDP:                 PROTO_TRDP,
	NDPI_PROTOCOL_LUSTRE:               PROTO_LUSTRE,
	NDPI_PROTOCOL_NORDVPN:              PROTO_NORDVPN,
	NDPI_PROTOCOL_SURFSHARK:            PROTO_SURFSHARK,
	NDPI_PROTOCOL_CACTUSVPN:            PROTO_CACTUSVPN,
	NDPI_PROTOCOL_WINDSCRIBE:           PROTO_WINDSCRIBE,
	NDPI_PROTOCOL_SONOS:                PROTO_SONOS,
	NDPI_PROTOCOL_DINGTALK:             PROTO_DINGTALK,
	NDPI_PROTOCOL_PALTALK:              PROTO_PALTALK,
	NDPI_PROTOCOL_NAVER:                PROTO_NAVER,
	NDPI_PROTOCOL_SHEIN:                PROTO_SHEIN,
	NDPI_PROTOCOL_TEMU:                 PROTO_TEMU,
	NDPI_PROTOCOL_TAOBAO:               PROTO_TAOBAO,
	NDPI_PROTOCOL_MIKROTIK:             PROTO_MIKROTIK,
	NDPI_PROTOCOL_DICOM:                PROTO_DICOM,
	NDPI_PROTOCOL_PARAMOUNTPLUS:        PROTO_PARAMOUNTPLUS,
	NDPI_PROTOCOL_YANDEX_ALICE:         PROTO_YANDEX_ALICE,
	NDPI_PROTOCOL_VIVOX:                PROTO_VIVOX,
	NDPI_PROTOCOL_DIGITALOCEAN:         PROTO_DIGITALOCEAN,
	NDPI_PROTOCOL_RUTUBE:               PROTO_RUTUBE,
	NDPI_PROTOCOL_LAGOFAST:             PROTO_LAGOFAST,
	NDPI_PROTOCOL_GEARUP_BOOSTER:       PROTO_GEARUP_BOOSTER,
	NDPI_PROTOCOL_RUMBLE:               PROTO_RUMBLE,
	NDPI_PROTOCOL_UBIQUITY:             PROTO_UBIQUITY,
	NDPI_PROTOCOL_MSDO:                 PROTO_MSDO,
	NDPI_PROTOCOL_ROCKSTAR_GAMES:       PROTO_ROCKSTAR_GAMES,
	NDPI_PROTOCOL_KICK:                 PROTO_KICK,
	NDPI_PROTOCOL_HAMACHI:              PROTO_HAMACHI,
	NDPI_PROTOCOL_GLBP:                 PROTO_GLBP,
	NDPI_PROTOCOL_EASYWEATHER:          PROTO_EASYWEATHER,
	NDPI_PROTOCOL_MUDFISH:              PROTO_MUDFISH,
	NDPI_PROTOCOL_TRISTATION:           PROTO_TRISTATION,
	NDPI_PROTOCOL_SAMSUNG_SDP:          PROTO_SAMSUNG_SDP,
	NDPI_PROTOCOL_MATTER:               PROTO_MATTER,
	NDPI_PROTOCOL_AWS_COGNITO:          PROTO_AWS_COGNITO,
	NDPI_PROTOCOL_AWS_API_GATEWAY:      PROTO_AWS_API_GATEWAY,
	NDPI_PROTOCOL_AWS_KINESIS:          PROTO_AWS_KINESIS,
	NDPI_PROTOCOL_AWS_EC2:              PROTO_AWS_EC2,
	NDPI_PROTOCOL_AWS_EMR:              PROTO_AWS_EMR,
	NDPI_PROTOCOL_AWS_S3:               PROTO_AWS_S3,
	NDPI_PROTOCOL_AWS_CLOUDFRONT:       PROTO_AWS_CLOUDFRONT,
	NDPI_PROTOCOL_AWS_DYNAMODB:         PROTO_AWS_DYNAMODB,
	NDPI_PROTOCOL_ESPN:                 PROTO_ESPN,
	NDPI_PROTOCOL_AKAMAI:               PROTO_AKAMAI,
	NDPI_PROTOCOL_JSON:                 PROTO_JSON,
	NDPI_PROTOCOL_MSGPACK:              PROTO_MSGPACK,
	NDPI_PROTOCOL_GITHUB_COPILOT:       PROTO_GITHUB_COPILOT,
	NDPI_PROTOCOL_GITHUB_PACKAGES:      PROTO_GITHUB_PACKAGES,
	NDPI_PROTOCOL_GITHUB_ACTIONS:       PROTO_GITHUB_ACTIONS,
	NDPI_PROTOCOL_YOUTUBE_KIDS:         PROTO_YOUTUBE_KIDS,
	NDPI_PROTOCOL_SBE:                  PROTO_SBE,
	NDPI_PROTOCOL_IRIS:                 PROTO_IRIS,
	NDPI_PROTOCOL_LIBP2P:               PROTO_LIBP2P,
	NDPI_PROTOCOL_NETMOTION:            PROTO_NETMOTION,
	NDPI_PROTOCOL_PROTON:               PROTO_PROTON,
	NDPI_PROTOCOL_PROTON_MAIL:          PROTO_PROTON_MAIL,
	NDPI_PROTOCOL_PROTON_PASS:          PROTO_PROTON_PASS,
	NDPI_PROTOCOL_PROTON_DRIVE:         PROTO_PROTON_DRIVE,
	NDPI_PROTOCOL_PROTON_CALENDAR:      PROTO_PROTON_CALENDAR,
	NDPI_PROTOCOL_PROTON_DOCS:          PROTO_PROTON_DOCS,
	NDPI_PROTOCOL_PROTON_WALLET:        PROTO_PROTON_WALLET,
	NDPI_PROTOCOL_PROTON_LUMO:          PROTO_PROTON_LUMO,
	NDPI_PROTOCOL_PROTON_MEET:          PROTO_PROTON_MEET,
	NDPI_PROTOCOL_YGGDRASIL:            PROTO_YGGDRASIL,
	NDPI_PROTOCOL_NEBULA:               PROTO_NEBULA,
	NDPI_PROTOCOL_DISCORD_CALL:         PROTO_DISCORD_CALL,
	NDPI_PROTOCOL_FREEFIRE:             PROTO_FREEFIRE,
}
