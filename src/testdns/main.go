package main

import(
    "testdns/dnspkg"
	"github.com/golang/protobuf/proto"
)

func main() {
	hostnames := []string{"www.sample.com", "www.google.com", "www.yahoo.com"}
	listHostname := &dnspkg.DnsLookupRequest{}
    listHostname.HostNames = hostnames;
	data, _ := proto.Marshal(listHostname)
	_, _ = dnspkg.ProcessDNSLookupRequest(data)
}
