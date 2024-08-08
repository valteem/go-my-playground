// go build -gcflags "all=-N -l"

package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	l, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 44566})
	if err != nil {
		log.Fatalf("error setting up tcp listener: %s", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("error listening: %s", err)
		}
		go func() {
			b := make([]byte, 1024)
			conn.Read(b)
			if err != nil {
				log.Fatalf("error reading from connection: %s", err)
			}
			fmt.Println(string(b[:]))
		}()
	}
}
