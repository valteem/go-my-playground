// go build -gcflags "all=-N -l"

package main

import (
	"fmt"
	"log"
	"net"
)

const (
	BUF_SIZE = 16
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
			r := make([]byte, 0)
			b := make([]byte, BUF_SIZE)
			for {
				n, err := conn.Read(b)
				if err != nil {
					log.Fatalf("error reading from connection: %s", err)
				}
				r = append(r, b[:n]...)
				if n < BUF_SIZE {
					break
				}
			}
			fmt.Println(string(r[:]))
		}()
	}
}
