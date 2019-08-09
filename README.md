# testdns
test miekg/dns package
1. cd src/testdns
2. source set-go-path
3. dep ensure -v
4. --specify the dns server ip, port, protocol details in dnspkg/dnsclient.go--
5. go run main.go

# start sample dns server
1. --specify the ip(machine ip), ip(any unused port) in dnsserver/dnsserver.go--
2. --specify sample hostname to ip address mappings in dnsserver/dnsserver.go--
3. go run dnsserver/dnsserver.go
