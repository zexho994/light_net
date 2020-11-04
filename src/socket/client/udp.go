package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// upd client
func UDPHandle(ip string, port int) {
	i := net.ParseIP(ip)
	desAddr := net.UDPAddr{
		IP:   i,
		Port: port,
	}

	conn, err := net.DialUDP("udp", nil, &desAddr)
	if err != nil {
		fmt.Println("dial udp error,", err.Error())
		return
	}
	ackMsg := make([]byte, 1024)
	for {
		fmt.Print(">")
		// 消息通过后续输入
		in := bufio.NewReader(os.Stdin)
		b, _, err := in.ReadLine()
		_, err = conn.Write(b)

		if err != nil {
			fmt.Println("write to udp error,", err.Error())
		}

		msgLen, remote, err := conn.ReadFromUDP(ackMsg)
		if err != nil {
			fmt.Println("read ack msg from udp error,", err.Error())
			return
		}
		fmt.Println(remote.String(), ":", string(ackMsg[:msgLen]))
	}
}
