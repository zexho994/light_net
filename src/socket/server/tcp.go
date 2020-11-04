package main

import (
	"fmt"
	"net"
	"os"
)

func tcpHandle(host string, port int) {
	// 建立tcp服务
	ip := net.ParseIP(host)
	tcpAddr := net.TCPAddr{
		IP:   ip,
		Port: port,
	}
	listen, err := net.ListenTCP("tcp", &tcpAddr)
	if err != nil {
		fmt.Printf("listen failed,err:%v\n", err)
		os.Exit(1)
	}
	fmt.Println("connect success")

	go process(listen)
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
			_, _ = conn.Write([]byte("ack"))
			if string(data[:i]) == "exit" {
				_ = conn.Close()
				return
			}
		}
	}
}
