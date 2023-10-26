package main

import (
	"fmt"
	"sync"
)

const (
	numReads = 5
)

func main() {

	c := make(chan int, 3)

	var wg sync.WaitGroup
	wg.Add(numReads)

	c <- 1
	c <- 2
	c <- 3

	close(c)

// A receive operation on a closed channel can always proceed immediately, yielding the element type's zero value after any previously sent values have been received
	for i := 1; i <= numReads; i++ {
		go func(i int) {
			fmt.Println("Reading #", i, "from channel returns", <-c)
			wg.Done()
		}(i)
	}

	wg.Wait()

}