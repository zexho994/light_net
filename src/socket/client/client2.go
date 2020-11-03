package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	flag.Parse()
	// 1.建立连接
	conn, err := net.Dial(*pf, *host+":"+*port)
	if err != nil {
		fmt.Printf("conn server failed , err:%v\n", err)
		return
	}

	for {
		fmt.Print("> ")
		in := bufio.NewReader(os.Stdin)
		b, _, err := in.ReadLine()
		if err != nil {
			fmt.Println("type error:", err.Error())
		}
		_, err = conn.Write(b)
		if err != nil {
			fmt.Println("write error ,", err.Error())
		}
	}
}
