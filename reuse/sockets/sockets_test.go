package sockets

import (
	"fmt"
	"net"
	"slices"

	"testing"
)

func TestUnixSocket(t *testing.T) {

	msg := []byte("some message")

	l, err := net.Listen("unix", "/tmp/echo.sock")
	if err != nil {
		t.Fatalf("failed create unix endpoint: %v", err)
	}
	defer l.Close() // removes /tmp/echo.sock

	errCh := make(chan error)

	go func() {

		defer close(errCh)

		conn, err := net.Dial("unix", "/tmp/echo.sock")
		if err != nil {
			errCh <- err
			return
		}
		defer conn.Close()

		n, err := conn.Write(msg)
		if err != nil {
			errCh <- err
			return
		}

		if n != len(msg) {
			errCh <- fmt.Errorf("write to unix connection: %d bytes, expect %d", n, len(msg))
		}

	}()

	for e := range errCh {
		t.Fatalf("client error: %v", e)
	}

	conn, err := l.Accept()
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		t.Fatalf("failed to read data from connection: %v", err)
	}
	if n != len(msg) {
		t.Errorf("expect to read %d bytes, get %d", len(msg), n)
	}

	output := buf[:n]
	if !slices.Equal(output, msg) {
		t.Errorf("")
	}

}
