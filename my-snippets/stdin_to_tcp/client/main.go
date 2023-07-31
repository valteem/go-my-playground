package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println("error establishing connection:", err)
		return
	}
	defer conn.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Message: ")
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading message")
			return
		}
		_, err = fmt.Fprintf(conn, string(message) + "\n")
		if err != nil {
			fmt.Println("error sending message")
		}
		if message == "end\n" {
			fmt.Println("closing client ...")
			return 
		}
	}
}