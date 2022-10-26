package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	go func () { messages <- "ping" } ()

	time.Sleep(time.Second)

	go func() {
		mssg := <- messages
		fmt.Println("go ", mssg)
	} ()

	msg := <- messages

	fmt.Println("main ", msg)

	c := make(chan string)

	for i := 0; i < 5; i++ {
		go func(i int, co chan<-string) {
			for j := 0; j < 5; j++ {
				co <- fmt.Sprintf("message from action %d.%d", i, j)
			}
		}(i, c)
	}

	for m := 0; m < 25; m++ {
		fmt.Println(<-c)
	}
}