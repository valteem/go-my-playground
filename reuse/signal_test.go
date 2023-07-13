package reuse_test

import (
	"fmt"
	"syscall"
	"testing"
	"time"

	"github.com/valteem/reuse"
)

func TestRunAppSimple(t *testing.T) {

	go reuse.RunAppSimple()
	time.Sleep(1 * time.Second)
	fmt.Println(syscall.Getpid())
	if err := syscall.Kill(syscall.Getpid(), syscall.SIGTERM); err != nil { // terminate
		fmt.Println("failed to send signal", err)
	}
	time.Sleep(1 * time.Second)

}

func TestRunAppNestedSignal(t *testing.T) {

	go reuse.RunAppNestedSignal()
	time.Sleep(1 * time.Second)
	fmt.Println(syscall.Getpid())
	if err := syscall.Kill(syscall.Getpid(), syscall.SIGINT); err != nil { // interrupt
		fmt.Println("failed to send signal", err)
	}
	time.Sleep(1 * time.Second)

}