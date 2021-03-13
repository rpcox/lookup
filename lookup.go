// lookup.go - Short quick DNS record lookup.  Perfect for use in a bash script.
package main

import (
  "context"
  "flag"
  "fmt"
  "net"
  "time"
)

var a_rec, cname, mx, ns, ptr, srv, proto, serv, txt, help string

func init() {
  flag.StringVar(&a_rec, "a",     "", "A record lookup. Hostname required")
  flag.StringVar(&cname, "cname", "", "CNAME record lookup. Hostname required")
  flag.StringVar(&mx,    "mx",    "", "MX record lookup. Domain name required")
  flag.StringVar(&ns,    "ns",    "", "NS record lookup. Domain name required")
  flag.StringVar(&ptr,   "ptr",   "", "PTR record lookup. IP address required {IPv4 or IPv6}")
  flag.StringVar(&srv,   "srv",   "", "Domain name required. See -p and -s")
  flag.StringVar(&proto, "p",     "", "Protocol required {tcp or udp}.  See -srv")
  flag.StringVar(&serv,  "s",     "", "Service required {e.g., xmpp-server}. See -srv")
  flag.StringVar(&txt,   "txt",   "", "Domain name required")
  flag.StringVar(&help,  "help",  "", "Show usage")
}

func main() {
  flag.Parse()

  // Sometimes nameservers timeout
  ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
  defer cancel()

  select {
  case <- time.After(1 * time.Second):
            fmt.Println("timeout")
  case <- ctx.Done():
            fmt.Println(ctx.Err())
  default:

    if a_rec != "" {                               // A
      list, err := net.LookupIP(a_rec)
      if ( err != nil ) { fmt.Println(err) }
      for _, ip := range list {
        fmt.Println(a_rec + "\tA\t", ip)
      }
    } else if cname != ""  {                       // CNAME
      name, err := net.LookupCNAME(cname)
      if ( err != nil ) { fmt.Println(err) }
      fmt.Println(cname + "\tCNAME\t", name)
    } else if mx != ""  {                          // MX
      mx_records, err := net.LookupMX(mx)
      if ( err != nil ) { fmt.Println(err) }
      for _, record := range mx_records {
        fmt.Println(mx + "\tMX\t" + record.Host + "\t", record.Pref)
      }
    } else if ns != ""  {                          // NS
      ns_records, err := net.LookupNS(ns)
      if ( err != nil ) { fmt.Println(err) }
      for _, record := range ns_records {
        fmt.Println(ns + "\tNS\t" + record.Host)
      }
    } else if ptr != ""  {                         // PTR
      names, err := net.LookupAddr(ptr)
      if ( err != nil ) { fmt.Println(err) }
      for _, name := range names {
        fmt.Println(ptr + "\tPTR\t" + name)
      }
    } else if srv != ""  {                         // SRV
      if proto != "" && serv != "" {
        name, services, err := net.LookupSRV(serv, proto, srv)
        for _, service := range services {
          fmt.Printf("%s\t%s\t%v\t%v\t%d\t%d\n", srv, name, service.Target, service.Port, service.Priority, service.Weight)
        }
        if ( err == nil ) {
        } else {
          fmt.Println(err)
        }
      } else {
        fmt.Println("service (-s), protocol (-p) and domain (-d) required")
      }
    } else if txt != ""  {                         // TXT
      records, err := net.LookupTXT(txt)
      if ( err != nil ) { fmt.Println(err) }
      for _, record := range records {
        fmt.Println(txt + "\tTXT\t" + record)
      }
    } else if help == ""  {                         // Usage
      fmt.Println("\nNAME\n    lookup - succinct DNS record retrieval\n")
      fmt.Println("SYNOPSIS\n    lookup [OPTION]\n")
      fmt.Println("DESCRIPTION\n    lookup [OPTION]\n")
      flag.PrintDefaults()
      fmt.Println()
    } else {
      fmt.Println("Try \"lookup -help\"")
    }
  }
}
