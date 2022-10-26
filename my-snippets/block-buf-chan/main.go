package main

import (
	"fmt"
)

func main() {

	maxIter := 10

	ch := make(chan int, 3)

	for i := 0; i < maxIter; i++ {
		fmt.Printf("Sending %d value to the channel\n", i)
		ch <- i
	}
}