package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

/**
和目标建立连接，然后发送输入数据
*/
func TCPHandle(host string, port int) {
	// 1.建立连接
	conn, err := net.Dial("tcp", host+":"+strconv.Itoa(port))
	if err != nil {
		fmt.Printf("conn server failed , err:%v\n", err)
		return
	}

	rb := make([]byte, 1024)
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

		ri, err := conn.Read(rb)
		if err != nil {
			fmt.Println("read error ," + err.Error())
			return
		}
		fmt.Println(string(rb[0:ri]))
	}
}
