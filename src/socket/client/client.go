package main

import (
	"flag"
	"fmt"
)

var host = flag.String("host", "127.0.0.1", "listening host , ip4 or ip6 or localhost")
var port = flag.Int("port", 9944, "listening port , 5000 ~ 65535")
var pf = flag.String("pf", "tcp", "<tcp> or <udp> or <all>")
var msg = flag.String("msg", " ", "The contents that you want send to")

func main() {
	flag.Parse()

	if *pf == "tcp" { // 处理tcp
		TCPHandle(*host, *port)
	} else if *pf == "udp" { // 处理udp
		UDPHandle(*host, *port)
	} else { // 不支持
		fmt.Println("pf is not support")
	}

}
