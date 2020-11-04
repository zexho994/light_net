package main

import (
	"flag"
	"fmt"
)

var host = flag.String("host", "localhost", "host, ip4 or ip6 or localhost")
var port = flag.Int("port", 9944, "port 5000 ~ 65535")
var pf = flag.String("pf", "tcp", "protocol family, <tcp> or <udp>")

func main() {
	flag.Parse()
	if *pf == "tcp" {
		tcpHandle(*host, *port)
	} else if *pf == "udp" { // udp
		udpHandle(*host, *port)
	} else { // err
		fmt.Println("pf is not support")
		return
	}
}
