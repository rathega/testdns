/* Test dns server
   Usage: go run dns_server.go
   Function: starts dns server in the serverip:port
             serves all dns requests with sample dns responses
             new hostnames can be added to map for testing
   Test dns request: dig @10.105.227.182 -p 5353 A www.sample.com
                     dig @<serverip> -p <serverport> <dnsType> <hostname>
*/

package main

import (
	"fmt"
	"log"

	"github.com/miekg/dns"
)

// ip(mac/vm ip), port(any unused port) to start the dns server
var serverip = "10.11.12.13"
var serverport = "5354"

// ipv4 hostname to ip map
// hostname(key) should be a fqdn
var ipv4ip = map[string][]string{
	"www.sample.com.": {"192.168.0.1", "192.168.0.2", "192.168.0.3", "192.168.0.4", "192.168.0.5", "192.168.0.6", "192.168.0.7"},
}

// ipv4 ip to ttl map
// if an ip does not have ttl entry, default value 3600 is used
var ipv4ttl = map[string]int{
	"192.168.0.1": 130,
	"192.168.0.2": 140,
	"192.168.0.3": 140,
	"192.168.0.4": 150,
	"192.168.0.5": 160,
	"192.168.0.6": 160,
}

// ipv6 hostname to ip map
// hostname(key) should be a fqdn
var ipv6ip = map[string][]string{
	"www.sample.com.": {"2404:6800:4007:80c::2004", "2001:4860:4860::6464", "2001:4860:4860::64", "2001:4860:4860::8888", "2001:4860:4860::8844"},
}

// ipv6 ip to ttl map
// if an ip does not have ttl entry, default value 3600 is used
var ipv6ttl = map[string]int{
	"2404:6800:4007:80c::2004": 20,
	"2001:4860:4860::6464":     40,
	"2001:4860:4860::64":       40,
	"2001:4860:4860::8888":     50,
	"2001:4860:4860::8844":     50,
}

func parseQuery(m *dns.Msg) {
	for _, q := range m.Question {
		switch q.Qtype {
		case dns.TypeA:
			log.Printf("\nQuery for %s for TypeA\n", q.Name)
			ip := ipv4ip[q.Name]
			if len(ip) <= 0 {
				return
			}
			for _, oneip := range ip {
				ttl := 3600
				if _, found := ipv4ttl[oneip]; found {
					ttl = ipv4ttl[oneip]
				}
				rr, err := dns.NewRR(fmt.Sprintf("%s %d A %s", q.Name, ttl, oneip))
				if err == nil {
					log.Printf("%+v\n", rr)
					m.Answer = append(m.Answer, rr)
				}
			}
		case dns.TypeAAAA:
			log.Printf("Query for %s for TypeAAAA\n", q.Name)
			ip := ipv6ip[q.Name]
			if len(ip) <= 0 {
				return
			}
			for _, oneip := range ip {
				ttl := 3600
				if _, found := ipv6ttl[oneip]; found {
					ttl = ipv6ttl[oneip]
				}
				rr, err := dns.NewRR(fmt.Sprintf("%s %d AAAA %s", q.Name, ttl, oneip))
				if err == nil {
					log.Printf("%+v\n", rr)
					m.Answer = append(m.Answer, rr)
				}
			}
		case dns.TypeANY:
			log.Printf("Query for %s for TypeANY\n", q.Name)
			ip := ipv4ip[q.Name]
			if len(ip) <= 0 {
				return
			}
			for _, oneip := range ip {
				ttl := 3600
				if _, found := ipv4ttl[oneip]; found {
					ttl = ipv4ttl[oneip]
				}
				rr, err := dns.NewRR(fmt.Sprintf("%s %d A %s", q.Name, ttl, oneip))
				if err == nil {
					log.Printf("%+v\n", rr)
					m.Answer = append(m.Answer, rr)
				}
			}
			ip = ipv6ip[q.Name]
			if len(ip) <= 0 {
				return
			}
			for _, oneip := range ip {
				ttl := 3600
				if _, found := ipv6ttl[oneip]; found {
					ttl = ipv6ttl[oneip]
				}
				rr, err := dns.NewRR(fmt.Sprintf("%s %d AAAA %s", q.Name, ttl, oneip))
				if err == nil {
					log.Printf("%+v\n", rr)
					m.Answer = append(m.Answer, rr)
				}
			}
		}
	}
}

func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Compress = false
	switch r.Opcode {
	case dns.OpcodeQuery:
		parseQuery(m)
	}
	w.WriteMsg(m)
}

func main() {
	// attach request handler func
	// '.' serves all host requests
	// '.com' serves only host.com requests
	dns.HandleFunc(".", handleDNSRequest)
	// start server
	server := &dns.Server{Addr: serverip + ":" + serverport, Net: "udp"}
	log.Printf("Starting at %s %s\n", serverip, serverport)
	err := server.ListenAndServe()
	defer server.Shutdown()
	if err != nil {
		log.Fatalf("Failed to start server: %s\n ", err.Error())
	}
}
