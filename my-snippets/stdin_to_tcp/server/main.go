package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("error setting listener")
		return
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("error opening port")
	}

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("error reading message", message)
		}
		s := string(message)
		if s == "end\n" {
			fmt.Println("closing listener ...")
			return
		} else {
			fmt.Println(s)
		}
	}
}