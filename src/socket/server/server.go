package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

var host = flag.String("host", "localhost", "host, ip4 or ip6 or localhost")
var port = flag.Int("port", 9944, "port 5000 ~ 65535")
var pf = flag.String("pf", "tcp", "protocol family, <tcp> or <udp>")

// 创建socket
func main1() {
	flag.Parse()
	listen, err := net.Listen(*pf, *host+":"+strconv.Itoa(*port))
	if err != nil {
		fmt.Println("Error listening", err)
		os.Exit(1)
	}

	defer listen.Close()
	fmt.Println("Listening on " + *host + ":" + strconv.Itoa(*port))

	for {

		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error Accept,", err)
			return
		}

		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	for {
		_, _ = io.Copy(conn, conn)
	}
}
