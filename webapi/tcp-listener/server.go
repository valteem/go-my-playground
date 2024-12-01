// https://go.dev/src/net/example_test.go

package main

import (
	"net"
)

func RunServer(addr, port string, handler func(net.Conn)) error {

	server, err := net.Listen("tcp", addr+":"+port)
	if err != nil {
		return err
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			return err
		}
		go handler(conn)
	}

}
