package main

import (
	"fmt"
	"net"
)

func main() {
	conn, error := net.Dial("udp", "8.8.8.8:53")
	if error != nil {
		fmt.Println(error)
	}

	defer conn.Close()
	ipAddress := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(ipAddress)
}
