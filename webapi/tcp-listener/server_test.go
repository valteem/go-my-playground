package main

import (
	"io"
	"net"
	"slices"
	"strconv"
	"testing"
)

const (
	msgNum  = 10
	srvAddr = "localhost"
	srvPort = "3001"
)

func handleMsg(buf *[]byte, t *testing.T) func(net.Conn) {
	return func(c net.Conn) {
		tmp := make([]byte, 256)
		for {
			n, err := c.Read(tmp)
			if err != nil {
				if err != io.EOF {
					t.Fatalf("failed to read from connection: %v", err)
				}
				break
			}
			*buf = append(*buf, tmp[:n]...)
		}
	}
}

func TestServer(t *testing.T) {

	bufOut := make([]byte, 0, 128)
	bufIn := make([]byte, 0, 128)

	go RunServer(srvAddr, srvPort, handleMsg(&bufOut, t))

	conn, err := net.Dial("tcp", srvAddr+":"+srvPort)
	if err != nil {
		t.Fatalf("failed to connect to server: %v", err)
	}

	for i := 0; i < msgNum; i++ {
		payload := []byte(strconv.Itoa(i) + ":")
		_, err := conn.Write(payload)
		if err != nil {
			t.Fatalf("failed to send a message #%d: %v", i, err)
		}
		bufIn = append(bufIn, payload...)
	}

	conn.Close() // sends EOF

	if !slices.Equal(bufIn, bufOut) {
		t.Errorf("expect\n%v\nget\n%v\n", bufIn, bufOut)
	}

}
