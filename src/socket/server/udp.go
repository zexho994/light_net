package main

import (
	"fmt"
	"net"
)

func udpHandle(host string, port int) {
	i := net.ParseIP(host)
	udpAddr := net.UDPAddr{
		IP:   i,
		Port: port,
	}
	listen, err := net.ListenUDP("udp", &udpAddr)
	if err != nil {
		fmt.Println("listen udp error ,", err)
	}
	fmt.Println("start listening udp...")
	msg := make([]byte, 1024)
	for {
		msgLen, remote, err := listen.ReadFromUDP(msg)
		if err != nil {
			fmt.Println("udp read msg error ,", err.Error())
			return
		}

		fmt.Println(remote.IP.String()+" msg:", string(msg[:msgLen]))
		_, err = listen.WriteToUDP([]byte("ack-udp"), remote)

		if err != nil {
			fmt.Println("udp write to udp error,", err.Error())
		}
	}
}
