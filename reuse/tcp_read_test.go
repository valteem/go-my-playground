// https://stackoverflow.com/a/3017456
// echo '2' > /tcp/dev/127.0.0.1:54467

// WIP/TODO: get 'error reading from connection' after first send

package reuse_test

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/valteem/reuse"
)

const (
	timeToSleep = 30
)

func TestTCPRead(t *testing.T) {

	server, err := net.Listen("tcp", ":54467")
	if err != nil {
		fmt.Println("error binding listener to socket")
	}
	defer server.Close()
	fmt.Println("listening ...")

	conn, err := server.Accept()
	if err != nil {
		fmt.Println("error accepting new connection")
	}
	ctx, cancel := context.WithCancelCause(context.Background())
	go reuse.ReadFromConn(conn, ctx, cancel)
	time.Sleep(time.Second * timeToSleep)
	cancel(fmt.Errorf("close connection after timeout"))

}