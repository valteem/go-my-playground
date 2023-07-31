package reuse_test

import (
	"log"
	"net"
	"testing"
	"time"

	"github.com/valteem/reuse"
)

func TestTCPServer(t *testing.T) {

	m := []string{"message1\n", "message2\n", "/quit"}

	s := &reuse.TCPServer{Port: ":54467"}

	go s.Run()
	time.Sleep(time.Second * time.Duration(1))

	conn, err := net.Dial("tcp", "127.0.0.1:54467")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for i := 0; i <2; i++ {
		for _, v := range m {
			_, err = conn.Write([]byte(v))
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	time.Sleep(time.Second * time.Duration(5))

}