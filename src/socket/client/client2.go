package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	flag.Parse()

	if *pf == "tcp" {
		// 1.建立连接
		conn, err := net.Dial(*pf, *host+":"+strconv.Itoa(*port))
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
	} else if *pf == "udp" {
		UDPClient(*host, *port)
	} else {
		fmt.Println("pf is not support")
	}
}
