package main

import "fmt"

func main() {
	queue := make(chan string, 5)
	queue <- "first message"
	queue <- "second message"
	close(queue)

	for{
		out, more := <- queue
		if more {
			fmt.Println(out)
		} else {
			fmt.Println("No more messages")
		}
	}
	
}