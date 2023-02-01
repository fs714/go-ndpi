package types

// Based on github.com/google/gopacket/layers/enums.go
type IPProto uint8

func (p *IPProto) ToName() string {
	name, ok := IPProtocolNameMap[*p]
	if !ok {
		name = ""
	}

	return name
}

const (
	IPProtocolIPv6HopByHop    IPProto = 0
	IPProtocolICMPv4          IPProto = 1
	IPProtocolIGMP            IPProto = 2
	IPProtocolIPv4            IPProto = 4
	IPProtocolTCP             IPProto = 6
	IPProtocolUDP             IPProto = 17
	IPProtocolRUDP            IPProto = 27
	IPProtocolIPv6            IPProto = 41
	IPProtocolIPv6Routing     IPProto = 43
	IPProtocolIPv6Fragment    IPProto = 44
	IPProtocolGRE             IPProto = 47
	IPProtocolESP             IPProto = 50
	IPProtocolAH              IPProto = 51
	IPProtocolICMPv6          IPProto = 58
	IPProtocolNoNextHeader    IPProto = 59
	IPProtocolIPv6Destination IPProto = 60
	IPProtocolOSPF            IPProto = 89
	IPProtocolIPIP            IPProto = 94
	IPProtocolEtherIP         IPProto = 97
	IPProtocolVRRP            IPProto = 112
	IPProtocolSCTP            IPProto = 132
	IPProtocolUDPLite         IPProto = 136
	IPProtocolMPLSInIP        IPProto = 137
)

const (
	IPProtocolIPv6HopByHopName    string = "IPv6HopByHop"
	IPProtocolICMPv4Name          string = "ICMPv4"
	IPProtocolIGMPName            string = "IGMP"
	IPProtocolIPv4Name            string = "IPv4"
	IPProtocolTCPName             string = "TCP"
	IPProtocolUDPName             string = "UDP"
	IPProtocolRUDPName            string = "RUDP"
	IPProtocolIPv6Name            string = "IPv6"
	IPProtocolIPv6RoutingName     string = "IPv6Routing"
	IPProtocolIPv6FragmentName    string = "IPv6Fragment"
	IPProtocolGREName             string = "GRE"
	IPProtocolESPName             string = "ESP"
	IPProtocolAHName              string = "AH"
	IPProtocolICMPv6Name          string = "ICMPv6"
	IPProtocolNoNextHeaderName    string = "NoNextHeader"
	IPProtocolIPv6DestinationName string = "IPv6Destination"
	IPProtocolOSPFName            string = "OSPF"
	IPProtocolIPIPName            string = "IPIP"
	IPProtocolEtherIPName         string = "EtherIP"
	IPProtocolVRRPName            string = "VRRP"
	IPProtocolSCTPName            string = "SCTP"
	IPProtocolUDPLiteName         string = "UDPLite"
	IPProtocolMPLSInIPName        string = "MPLSInIP"
)

var IPProtocolNameMap = map[IPProto]string{
	IPProtocolIPv6HopByHop:    IPProtocolIPv6HopByHopName,
	IPProtocolICMPv4:          IPProtocolICMPv4Name,
	IPProtocolIGMP:            IPProtocolIGMPName,
	IPProtocolIPv4:            IPProtocolIPv4Name,
	IPProtocolTCP:             IPProtocolTCPName,
	IPProtocolUDP:             IPProtocolUDPName,
	IPProtocolRUDP:            IPProtocolRUDPName,
	IPProtocolIPv6:            IPProtocolIPv6Name,
	IPProtocolIPv6Routing:     IPProtocolIPv6RoutingName,
	IPProtocolIPv6Fragment:    IPProtocolIPv6FragmentName,
	IPProtocolGRE:             IPProtocolGREName,
	IPProtocolESP:             IPProtocolESPName,
	IPProtocolAH:              IPProtocolAHName,
	IPProtocolICMPv6:          IPProtocolICMPv6Name,
	IPProtocolNoNextHeader:    IPProtocolNoNextHeaderName,
	IPProtocolIPv6Destination: IPProtocolIPv6DestinationName,
	IPProtocolOSPF:            IPProtocolOSPFName,
	IPProtocolIPIP:            IPProtocolIPIPName,
	IPProtocolEtherIP:         IPProtocolEtherIPName,
	IPProtocolVRRP:            IPProtocolVRRPName,
	IPProtocolSCTP:            IPProtocolSCTPName,
	IPProtocolUDPLite:         IPProtocolUDPLiteName,
	IPProtocolMPLSInIP:        IPProtocolMPLSInIPName,
}
