package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

var host = flag.String("host", "localhost", "listening host , ip4 or ip6 or localhost")
var port = flag.String("port", "9944", "listening port , 5000 ~ 65535")
var pf = flag.String("pf", "tcp", "<tcp> or <udp> or <all>")

// 创建socket
func main() {
	flag.Parse()
	listen, err := net.Listen(*pf, *host+":"+*port)
	if err != nil {
		panic("Error listening")
		os.Exit(1)
	}

	defer listen.Close()
	fmt.Println("Listening on " + *host + ":" + *port)

	for {

		conn, err := listen.Accept()
		if err != nil {
			panic("Error accepting")
			os.Exit(1)
		}

		fmt.Println("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	for {
		_, _ = io.Copy(conn, conn)
	}
}
