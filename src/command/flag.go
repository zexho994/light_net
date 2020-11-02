package command

import (
	"flag"
	"fmt"
)

func Cmd() {
	ip := flag.String("af", "ip4", "")
	protocol := flag.String("pf", "tcp", "")

	flag.Parse()

	fmt.Println("addr family", *ip)
	fmt.Println("protocol family", *protocol)
}
