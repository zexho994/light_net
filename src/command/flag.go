package command

import (
	"flag"
)

var host = flag.String("host", "localhost", "host")
var port = flag.String("port", "9944", "port")

func Cmd() {
	flag.Parse()
}
