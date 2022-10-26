// If we know number of messages to send and to read in advance, no 'select' or 'context' is needed
// Use context with cancel to stop goroutines at some given moment, 

package main

import (
	"fmt"
	"sync"
)

const (
	numWorkers = 5
	numSignals = 100
)

func main() {

	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 1; i <= numWorkers; i++ {
		go func(p int) {
			for rm := range ch {
				fmt.Println("Message", rm, "in goroutine", p)
				wg.Done()
			}
		}(i)
	}

	for j := 1; j <= numSignals; j++ {
		wg.Add(1)
		ch <- j
	}

	wg.Wait()

}