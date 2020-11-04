package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

var host = flag.String("host", "127.0.0.1", "listening host , ip4 or ip6 or localhost")
var port = flag.Int("port", 9944, "listening port , 5000 ~ 65535")
var pf = flag.String("pf", "tcp", "<tcp> or <udp> or <all>")

func main1() {
	flag.Parse()
	conn, err := net.Dial(*pf, *host+":"+strconv.Itoa(*port))
	if err != nil {
		fmt.Println("Error connection", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connecting to " + *host + ":" + strconv.Itoa(*port))

	// 使用wait group进行同步
	var wg sync.WaitGroup
	wg.Add(1)

	in := strings.Join(flag.Args(), " ")

	if in == "exit" {
		wg.Add(1)
	} else {
		_, _ = conn.Write([]byte(in))
	}

	wg.Wait()
}

func write(msg string, conn net.Conn) {
	_, _ = conn.Write([]byte(msg))
}

func handleWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 10; i > 0; i-- {
		_, e := conn.Write([]byte("hello " + strconv.Itoa(i) + "\r\n"))
		if e != nil {
			fmt.Println("Error to send message because of ", e.Error())
			break
		}
	}
}

func handleRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	reader := bufio.NewReader(conn)
	for i := 1; i <= 10; i++ {
		line, err := reader.ReadString(byte('\n'))
		if err != nil {
			fmt.Print("Error to read message because of ", err)
			return
		}
		fmt.Print(line)
	}
}
