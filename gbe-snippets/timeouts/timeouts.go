package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func(c chan string) {
		time.Sleep(2 * time.Second)
		c <- "message 1"
	}(c1)

	go func(c chan string) {
		time.Sleep(3 * time.Second)
		c <- "message 2"
	}(c2)

	select {
	case res := <- c1:
		fmt.Println(res)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout 1")
	}

	select {
	case res := <- c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 1")
	}
}