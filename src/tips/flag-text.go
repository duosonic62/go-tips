package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	var ip net.IP
	// p encoding.TextUnmarshaler, name string, value encoding.TextMarshaler, usage string
	flag.TextVar(&ip, "ip", net.IPv4(127, 0, 0, 1), "ip address")
	flag.Parse()

	fmt.Println(ip)
}
