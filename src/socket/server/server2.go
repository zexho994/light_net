package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	flag.Parse()
	if *pf == "tcp" {
		// 建立tcp服务
		ip := net.ParseIP(*host)
		tcpAddr := net.TCPAddr{
			IP:   ip,
			Port: *port,
		}
		listen, err := net.ListenTCP(*pf, &tcpAddr)
		if err != nil {
			fmt.Printf("listen failed,err:%v\n", err)
			os.Exit(1)
		}
		fmt.Println("connect success")

		process(listen)
	} else if *pf == "udp" { // udp

	} else { // err

	}

}

func process(listen *net.TCPListener) {
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("accept tcp error,", err.Error())
		}

		data := make([]byte, 1024)
		for {
			i, err := conn.Read(data)
			if err != nil {
				fmt.Println("read msg err,", err.Error())
				return
			}
			fmt.Println(conn.RemoteAddr().String(), "msg: ", string(data[0:i]))
			if string(data[:i]) == "exit" {
				_, _ = conn.Write([]byte{'e', 'x', 'i', 't'})
				_ = conn.Close()
				return
			}
			_, _ = conn.Write(data[0:i])
		}
	}
}
