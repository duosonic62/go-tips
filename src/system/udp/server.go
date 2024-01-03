package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("UDP Server")
	conn, err := net.ListenPacket("udp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buffer := make([]byte, 1500)
	for {
		length, remoteAdress, err := conn.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received from %v: %v\n", remoteAdress, string(buffer[:length]))
		_, err = conn.WriteTo([]byte("Hello from server"), remoteAdress)
		if err != nil {
			panic(err)
		}
	}
}
