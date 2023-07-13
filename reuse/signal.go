package reuse

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func RunAppSimple() {

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	fmt.Println("starting ...")
	sig := <-quit
	fmt.Println(syscall.Getpid(), sig)
	fmt.Println("... quitting")

}

func RunAppNestedSignal() {

	quit := make(chan bool)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	fmt.Println("starting trunk ...")
	go func() {
		fmt.Println("starting nested ...")
		s := <-sig
		fmt.Println(syscall.Getpid(), s)
		quit <- true
	}()
	fmt.Println("waiting for signal ...")
	<- quit
	fmt.Println("... quitting")

}