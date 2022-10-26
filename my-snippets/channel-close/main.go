package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wgc sync.WaitGroup
	wgc.Add(1)

	numWorkers := 5

	maxDelay := 10

	closing := make(chan int)
	watching := make(chan int, numWorkers)

	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= numWorkers; i++ {
		go func(p int) {
			for {
				select {
				case sg := <-closing:
					fmt.Println("Goroutine", p, "receives closing signal", sg)
					delay := rand.Float32() * float32(maxDelay)
					time.Sleep(time.Duration( delay * float32(time.Second) ))
					fmt.Println("Goroutine", p, "worked", delay, "seconds before terminating")
					watching <- 1
					return
				default:
					fmt.Println("Goroutine", p, "- nothing happens")
				}
			}
		}(i)
	}

	go func() {
		defer wgc.Done() // this is the only reason to run Done() in a separate goroutine
		// making sure all other goroutines receive closing signal and quit
		count := 0
		for input := range watching {
			count += input
			fmt.Println("Count of stopped goroutines is", count)
			if count == numWorkers {
				fmt.Println("Done!")
				return
			}
		}
	}()

	close(closing)
	wgc.Wait()
}