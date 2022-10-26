package main

import "fmt"

func main() {
	queue := make (chan string, 2)
	queue <- "first item"
	queue <- "second item"
	close(queue	)

	fmt.Println("Queue length: ", len(queue))
	fmt.Println("Queue capacity: ", cap(queue))

	for i := range(queue) {
		fmt.Println(i)
	}
}