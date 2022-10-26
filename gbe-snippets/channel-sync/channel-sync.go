package main

import (
	"fmt"
	"time"
)

func Worker (done chan bool) {
	fmt.Println("working ... ")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {
	c := make(chan bool)

	go Worker(c)

	<- c
}