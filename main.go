package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/fatih/color"
	"github.com/miekg/dns"
)

var ver = "0.0.1"
var ns = "8.8.8.8"

func main() {
	usage := `基于 Golang 的 DNS 查询工具.

Usage:
  nslookup <domain>
  nslookup <domain> <dns>
  nslookup -h | --help
  nslookup --version

Options:
  -h --help     Show this screen.
  --version     Show version.`

	arguments, _ := docopt.ParseDoc(usage)

	domain, _ := arguments.String("<domain>")
	dns, _ := arguments.String("<dns>")
	version, _ := arguments.Bool("--version")

	if version {
		color.Cyan("version: " + ver)
		return
	}

	if dns == "" {
		dns = ns
	}

	record := nslookup(domain, dns)
	fmt.Println(record)
}

func nslookup(target, server string) (res []string) {
	c := dns.Client{}
	m := dns.Msg{}
	m.SetQuestion(target+".", dns.TypeA)

	ns := server+":53"
	r, t, err := c.Exchange(&m, ns)
	if err != nil {
		color.Red("nameserver %s error: %v", ns, err)
		return res
	}

	color.Cyan("nameserver %s took %v", ns, t)

	if len(r.Answer) == 0 {
		return res
	}

	for _, ans := range r.Answer {
		Arecord := ans.(*dns.A)

		res = append(res, fmt.Sprintf("%s", Arecord))
	}

	return
}
