package main

import (
	"fmt"
	"strconv"
)

func signalReceiver(n int, ch <-chan string, q <-chan bool) {

	var msg string

	for {
		select {
		case <-q:
			return
		default:
	        msg = <-ch
		    fmt.Printf("%d receiver %s\n", n, msg)
		}

	}
	
}

func main() {
	
	max_iter := 100

	c := make(chan string)
	q := make(chan bool)

	go signalReceiver(1, c, q)
	go signalReceiver(2, c, q)
	go signalReceiver(3, c, q)
	go signalReceiver(4, c, q)

	for count := 1; count <= max_iter; count++ {
		msg := "signal " + strconv.Itoa(count)
		c <- msg
	}
	q <- true
	close(q)
	close(c)

}