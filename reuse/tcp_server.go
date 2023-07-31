package reuse

import (
	"fmt"
	"log"
	"net"
)

type TCPServer struct{
	Port string
}

func (s *TCPServer) read(conn net.Conn) {

	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		message := string(buffer[:n])

		// if message == "/quit" {
		// 	fmt.Println("quitting ...")
		// 	return
		// }

		if n > 0 {
			fmt.Println(n, message)
		}

		if err != nil {
			log.Println(err)
			return
		}	
	}

}

func (s *TCPServer) Run() {

	listener, err := net.Listen("tcp", s.Port)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go s.read(conn)
	}

}