package main

import (
	"flag"
	"fmt"
	"net"
)

var socketPath = flag.String("sp", "/tmp/ipc.socket", "The another process's socket path")

func main() {
	flag.Parse()
	listen, err := net.Listen("unix", *socketPath)
	if err != nil {
		panic("listen unix error," + err.Error())
	}
	fmt.Println("server start listening...")
	data := make([]byte, 1024)
	for {
		conn := handleAccept(listen)
		handleRead(conn, data)
		replayAck(conn)
	}

}

func handleAccept(listener net.Listener) net.Conn {
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listen accept error,", err.Error())
		return nil
	}
	return conn
}

func handleRead(conn net.Conn, data []byte) {
	i, err := conn.Read(data)
	if err != nil {
		fmt.Println("conn read data error,", err.Error())
		return
	}
	fmt.Println("msg: ", string(data[:i]))
}

func replayAck(conn net.Conn) {
	_, err := conn.Write([]byte("ack"))
	if err != nil {
		fmt.Println("conn write ack error,", err.Error())
	}
}
