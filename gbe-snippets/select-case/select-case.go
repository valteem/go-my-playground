package main

import (
	"fmt"
/*	"time" */
)

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func(c chan string) {
/*		time.Sleep(time.Second) */
		c <- "first message"
	}(c1)

	go func(c chan string) {
/*		time.Sleep(time.Second) */
		c <- "second message"
	}(c2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <- c1:
			fmt.Println(msg1)
		case msg2 := <- c2:
			fmt.Println(msg2)
		}
/*	msg1 := <- c1
	fmt.Println(msg1)
	msg2 := <- c2
	fmt.Println(msg2) */
	}
}