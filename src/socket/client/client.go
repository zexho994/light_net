package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
)

var host = flag.String("host", "localhost", "listening host , ip4 or ip6 or localhost")
var port = flag.String("port", "9944", "listening port , 5000 ~ 65535")
var pf = flag.String("pf", "tcp", "<tcp> or <udp> or <all>")

func main() {
	flag.Parse()
	conn, err := net.Dial(*pf, *host+":"+*port)
	if err != nil {
		fmt.Println("Error connection", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connecting to " + *host + ":" + *port)

	// 使用waitgroup进行同步
	var wg sync.WaitGroup
	wg.Add(2)

}
