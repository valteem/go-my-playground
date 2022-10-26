package main

import "fmt"

func ping (c chan<-string, s string) {
	c <- s
}

func pong(cRecv<- chan string, cSend chan<-string) {
	
	msg := <- cRecv
	cSend <- msg

}

func main() {

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "message to pass")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}