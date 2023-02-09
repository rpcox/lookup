// lookup - Succinct DNS record lookup. Useful in a bash script.
package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func ShowVersion() {
	ver := "0.1.0"
	fmt.Println("lookup v", ver)
	os.Exit(0)
}

func ShowUsage(code int) {
	fmt.Println("\nNAME\n    lookup - succinct DNS record retrieval")
	fmt.Println()
	fmt.Println("SYNOPSIS\n    lookup [OPTION]")
	fmt.Println()
	fmt.Println("DESCRIPTION\n    lookup [OPTION]")
	fmt.Println()
	flag.PrintDefaults()
	fmt.Println()
	os.Exit(code)
}

func isValidIP(object string) bool {
	if net.ParseIP(object) != nil {
		return true
	} else {
		return false
	}
}

func PtrLookup(s string) {
	names, err := net.LookupAddr(s)
	if err != nil {
		fmt.Println(err)
	}

	for _, name := range names {
		fmt.Println(s + "\tPTR\t" + name)
	}
}

func ALookup(s string) {
	list, err := net.LookupIP(s)
	if err != nil {
		fmt.Println(err)
	}

	for _, ip := range list {
		fmt.Println(s+"\tA\t", ip)
	}
}

func CNameLookup(s string) {
	name, err := net.LookupCNAME(s)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s+"\tCNAME\t", name)
}

func MxLookup(s string) {
	mx_records, err := net.LookupMX(s)
	if err != nil {
		fmt.Println(err)
	}

	for _, record := range mx_records {
		fmt.Println(s+"\tMX\t"+record.Host+"\t", record.Pref)
	}
}

func NsLookup(s string) {
	ns_records, err := net.LookupNS(s)
	if err != nil {
		fmt.Println(err)
	}

	for _, record := range ns_records {
		fmt.Println(s + "\tNS\t" + record.Host)
	}
}

func SrvLookup(srv, proto, serv string) {
	if proto != "" && serv != "" {
		name, services, err := net.LookupSRV(serv, proto, srv)
		for _, service := range services {
			fmt.Printf("%s\t%s\t%v\t%v\t%d\t%d\n", srv, name, service.Target, service.Port, service.Priority, service.Weight)
		}

		if err != nil {
			fmt.Println(err)
		}

	} else {
		fmt.Println("service (-s), protocol (-p) and domain (-d) required")
	}
}

func TxtLookup(s string) {
	records, err := net.LookupTXT(s)
	if err != nil {
		fmt.Println(err)
	}

	for _, record := range records {
		fmt.Println(s + "\tTXT\t" + record)
	}
}

func main() {
	var (
		a_rec   = flag.String("a", "", "A record lookup. Hostname required")
		cname   = flag.String("cname", "", "CNAME record lookup. Hostname required")
		mx      = flag.String("mx", "", "MX record lookup. Domain name required")
		ns      = flag.String("ns", "", "NS record lookup. Domain name required")
		ptr     = flag.String("ptr", "", "PTR record lookup. IP address required {IPv4 or IPv6}")
		srv     = flag.String("srv", "", "Domain name required. See -p and -s")
		proto   = flag.String("p", "", "Protocol required {tcp or udp}.  See -srv")
		serv    = flag.String("s", "", "Service required {e.g., xmpp-server}. See -srv")
		to      = flag.Int("timeout", 3, "Timeout value (seconds)")
		txt     = flag.String("txt", "", "Domain name required")
		help    = flag.Bool("help", false, "Show usage")
		version = flag.Bool("version", false, "Show version")
	)

	flag.Parse()

	if *help {
		ShowUsage(0)
	}

	if *version {
		ShowVersion()
	}

	netObject := ""

	if flag.NArg() > 0 {
		netObject = os.Args[1]
	}

	// Sometimes nameservers timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(int(time.Millisecond)*(*to*1000)))
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	default:

		if netObject != "" {

			if isValidIP(netObject) {
				PtrLookup(netObject)
			} else {
				ALookup(netObject)
			}

		} else {

			if *a_rec != "" {
				ALookup(*a_rec)
			} else if *cname != "" {
				CNameLookup(*cname)
			} else if *mx != "" {
				MxLookup(*mx)
			} else if *ns != "" {
				NsLookup(*ns)
			} else if *ptr != "" {
				PtrLookup(*ptr)
			} else if *srv != "" {
				SrvLookup(*srv, *proto, *serv)
			} else if *txt != "" {
				TxtLookup(*txt)
			} else {
				fmt.Println("Try \"lookup -help\"")
			}
		}
	}
}
