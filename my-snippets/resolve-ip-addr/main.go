package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	localAddr, err := net.ResolveIPAddr("ip4", "localhost")
	if err != nil {
		log.Fatal("error resolving IP address")
		return
	}
	fmt.Println(localAddr)

}