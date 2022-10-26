package main

import "fmt"
import "time"

func main() {
	cl1 := make(chan string, 3)

	go func () {
		time.Sleep(time.Second * 1)
		cl1 <- "first message"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		cl1 <- "second message"
	}()

	str := <- cl1
	fmt.Println(str)

	cl1 <- "third message"
	fmt.Println(<-cl1)
}