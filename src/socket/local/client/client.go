package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

var socketPath = flag.String("sp", "/tmp/ipc.socket", "The another process's path")

func main() {
	flag.Parse()
	conn, err := net.Dial("unix", *socketPath)
	if err != nil {
		panic("net dial error," + err.Error())
	}
	fmt.Println("connect success, please type some ...")
	data := make([]byte, 1024)
	for true {
		fmt.Print("> ")
		in := bufio.NewReader(os.Stdin)
		b, _, err := in.ReadLine()
		_, err = conn.Write(b)
		if err != nil {
			panic("unix conn write error," + err.Error())
		}
		i, err := conn.Read(data)
		if err != nil {
			panic("unix read ack error" + err.Error())
		}
		fmt.Println(string(data[:i]))
	}
}
