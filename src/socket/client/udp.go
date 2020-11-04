package main

import (
	"fmt"
	"net"
)

// upd client
func UDPClient(ip string, port int) {
	i := net.ParseIP(ip)
	ua := net.UDPAddr{
		IP:   i,
		Port: port,
	}

	listen, err := net.ListenUDP("udp", &ua)
	if err != nil {
		fmt.Println("listen upd file,", err.Error())
	}

	msg := make([]byte, 1024)
	for {
		// read msg
		n, remoteAddr, err := listen.ReadFromUDP(msg)
		if err != nil {
			fmt.Println("read upd msg error , ", err.Error())
			return
		}

		// print msg
		fmt.Printf("<%s> %s \n", remoteAddr, msg[:n])

		// send ack
		b := []byte{'a', 'c', 'k'}
		_, err = listen.WriteToUDP(b, &ua)
		if err != nil {
			fmt.Println("write msg to udp error ,", err.Error())
		}

	}

}
