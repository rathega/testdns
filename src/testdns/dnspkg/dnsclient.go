package dnspkg 

import (
	"net"
	"time"
	"strconv"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/miekg/dns"
)

// DNSLookupResponse struct populates the ip address for the fqdn received from the dns server
type DNSLookupResponse struct {
	fqdn     string
	ipv4Addr string
	ipv6Addr string
}

type DnsServer struct {
    IP       string `yaml:"ip"`
    Port     int    `yaml:"port"`
    Protocol string `yaml:"protocol"`
}

var getConfigDnsServers = []DnsServer{DnsServer{IP: "64.104.76.247", Port: 53, Protocol: "udp"}}
var getConfigQueryType = "ipv4"
var getConfigTimeoutMs = 1000

// ProcessDNSLookupRequest function processes the dns lookup request
func ProcessDNSLookupRequest(msg []byte) ([]byte, string) {
	dnsLookup := &DnsLookupRequest{}
	err := proto.Unmarshal(msg, dnsLookup)
	if err != nil {
		fmt.Printf("\nerror converting bytes to pb: %v", err)
		return nil, "error converting bytes to pb"
	}
	hostNames := dnsLookup.GetHostNames()
	fmt.Printf("\nHostnames: %v", hostNames)
	ch := make(chan *DNSLookupResponse)
	// addresses := make([]string, len(hostNames))
	for _, hostName := range hostNames {
		fmt.Printf("\nReceived dnslookup request for host: %v", hostName)
		go queryDNS(hostName, ch)
	}

	dnsLookupGrpcResponse := &DnsLookupResponse{}
	hostNameToIPAddressMapping := dnsLookupGrpcResponse.GetHostNameMapping()

	for range hostNames {
		// addresses = append(addresses, <-ch)
		dnsLookupResponse := <-ch
		hostNameToIPAddressMappingElement := HostNameToIpAddressMapping{Ipv4Addr: dnsLookupResponse.ipv4Addr, Ipv6Addr: dnsLookupResponse.ipv6Addr, HostName: dnsLookupResponse.fqdn}
		hostNameToIPAddressMapping = append(hostNameToIPAddressMapping, &hostNameToIPAddressMappingElement)
	}
	dnsLookupGrpcResponse.HostNameMapping = hostNameToIPAddressMapping
	fmt.Printf("\nHostNameMapping set to %v\n", hostNameToIPAddressMapping)

	rspMsg, err := proto.Marshal(dnsLookupGrpcResponse)
	if err != nil {
		fmt.Printf("\nerror converting bytes to pb: %v", err)
		return nil, "error converting bytes to pb"
	}
	return rspMsg, ""

}

func queryDNS(addr string, ch chan *DNSLookupResponse) {
	m1 := new(dns.Msg)
	m1.Id = dns.Id()
	m1.RecursionDesired = true

	m1.Question = make([]dns.Question, 1)
	qType := dns.TypeANY
	queryType := getConfigQueryType
	switch queryType {
	case "ipv4":
		qType = dns.TypeA
	case "ipv6":
		qType = dns.TypeAAAA
	}
	fmt.Printf("\nqType for %v is %v", addr, qType)
	m1.Question[0] = dns.Question{Name: dns.Fqdn(addr), Qtype: qType, Qclass: dns.ClassINET}
	c := new(dns.Client)
	c.Timeout = time.Duration(getConfigTimeoutMs) * time.Millisecond
	dnsServers := getConfigDnsServers
	fmt.Printf("\nCurrent DNS servers list : %v", dnsServers)
	dnsLookupResponse := &DNSLookupResponse{}
	dnsLookupResponse.fqdn = addr
	ipv4Address, ipv6Address := "", ""
	for _, dnsServer := range dnsServers {
		url := dnsServer.IP + ":" + strconv.Itoa(dnsServer.Port)
		c.Net = dnsServer.Protocol
		fmt.Printf("\nQuerying DNS server %v", url)
		//startTime := time.Now().UnixNano()
		in, _, err := c.Exchange(m1, url)
		//endTime := time.Now().UnixNano()
		if err != nil {
			if err, ok := err.(net.Error); ok && err.Timeout() {
				fmt.Printf("\nTimeout occurred while receiving response from %v", url)
				continue
			}
			fmt.Printf("\nError occurred while sending dns lookup query to DNS server: %v, Error: %v", dnsServer, err)
			continue
		}
		if in != nil {
			entry := false
            fmt.Printf("\nReceived dns lookup response for host: %v", addr)
			if qType == dns.TypeA {
				ipv4Address = getIpv4Address(in)
				if ipv4Address != "" {
					entry = true
					break
				}
			}
			if qType == dns.TypeAAAA {
				ipv6Address = getIpv6Address(in)
				if ipv6Address != "" {
					entry = true
					break
				}
			}
			if qType == dns.TypeANY {
				if ipv4Address == "" {
					ipv4Address = getIpv4Address(in)
				}
				if ipv6Address == "" {
					ipv6Address = getIpv6Address(in)
				}
				if ipv4Address != "" || ipv6Address != "" {
					entry = true
					break
				}
			}
			result := "entry_found"
			if !entry {
				result = "noentry"
			}
			fmt.Printf("\nresult: %v", result)
		} else {
			fmt.Printf("\nNo response received from dns server %v. Try with next configured dns server", url)
		}
	}
	dnsLookupResponse.ipv4Addr = ipv4Address
	dnsLookupResponse.ipv6Addr = ipv6Address
	fmt.Printf("\nSending dnsLookup response : %v,", *dnsLookupResponse)
	ch <- dnsLookupResponse

}

func getIpv4Address(in *dns.Msg) string {
	for _, rr := range in.Answer {
		_, ok := rr.(*dns.A)
		if ok {
			Arecord := rr.(*dns.A)
            fmt.Printf("\nDns Answer fields: %+v", Arecord.Header())
            fmt.Printf("\n\n**Dns TTL field** %v\n\n", Arecord.Header().Ttl)
			addresss := Arecord.A.String()
			return addresss
		}
	}
	return ""
}

func getIpv6Address(in *dns.Msg) string {
	for _, rr := range in.Answer {
		_, ok := rr.(*dns.AAAA)
		if ok {
			AAAArecord := rr.(*dns.AAAA)
            fmt.Printf("\nDns Answer fields: %+v", AAAArecord.Header())
            fmt.Printf("\n\n**Dns TTL field**: %v\n\n", AAAArecord.Header().Ttl)
			addresss := AAAArecord.AAAA.String()
			return addresss
		}
	}
	return ""
}
