### lookup

Succinct DNS record lookup.

I needed simple terse easy to parse DNS record information for a bash script. This worked.

    go build lookup.go

All returned fields are tab separated.

     > lookup www.golang.org
     www.golang.org	A	 172.217.4.177
     www.golang.org	A	 2607:f8b0:4007:80e::2011
     >

     > lookup.go -a www.golang.org
     www.golang.org	A	 172.217.4.177
     www.golang.org	A	 2607:f8b0:4007:80e::2011
     >

     > lookup.go -cname www.cisco.com
     www.cisco.com	CNAME	 e2867.dsca.akamaiedge.net.
     >

     > lookup.go -mx cox.com
     cox.com	MX	mxa-002b3901.gslb.pphosted.com.	 10
     cox.com	MX	mxb-002b3901.gslb.pphosted.com.	 10
     >

     > lookup.go -ns google.com
     google.com	NS	ns3.google.com.
     google.com	NS	ns2.google.com.
     google.com	NS	ns4.google.com.
     google.com	NS	ns1.google.com.
     >

     > lookup.go 216.239.32.10
     216.239.32.10	PTR	ns1.google.com.
     > go run lookup.go -ptr 2607:f8b0:4007:80e::2011
     2607:f8b0:4007:80e::2011	PTR	lax31s01-in-x11.1e100.net.
     >

     > lookup.go -ptr 216.239.32.10
     216.239.32.10	PTR	ns1.google.com.
     > go run lookup.go -ptr 2607:f8b0:4007:80e::2011
     2607:f8b0:4007:80e::2011	PTR	lax31s01-in-x11.1e100.net.
     >

     > lookup.go -srv spud.com -s xmpp-server -p tcp
     lookup _xmpp-server._tcp.spud.com on 127.0.0.53:53: no such host
     >

     > lookup.go -txt costco.com
     costco.com	TXT	have-i-been-pwned-verification=7411effeb11c12300a4c027396b4cf0f
     costco.com	TXT	docusign=3e51deb4-b2f9-46ef-b120-18f6955ebce1
     costco.com	TXT	MS=ms45356090
     costco.com	TXT	docusign=4db9a87a-9dea-417f-bfaa-d45bf7c94291
     costco.com	TXT	ciscocidomainverification=778c9d75e8b6fdb7dc3f095b677d08d21ececebe8462cb691c377eb0c061c825
     costco.com	TXT	facebook-domain-verification=46if6nzbzbffbhikh3uq8sye7ez5bg
     costco.com	TXT	google-site-verification=dffTnhrKhs3V5-UUHvEYJg-RNLJjBz27jCBqmF6sX_E
     costco.com	TXT	google-site-verification=3Kx9tc2_v5CgX3NuVqNwWFZRYanPPJCP3__8KbRxn5Q
     costco.com	TXT	adobe-idp-site-verification=ae9e3f0f-2848-4000-aed4-e1c7e8031815
     costco.com	TXT	v=spf1 include:_spf.costco.com include:_spf.google.com include:_spf1.costco.com      include:_spf2.costco.com include:_spf3.costco.com include:_spf4.costco.com include:spf-001f9d01.pphosted.com -all
     >

