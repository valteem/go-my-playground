package main

import (
	"fmt"
	"sync"
)

type Work struct {
	number int
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	numWorkers := 5

	ch := make(chan Work, numWorkers)
	defer close(ch) // closing channel prevents later 'fatal error: all goroutines are asleep - deadlock!'

	go func() {
		fmt.Println("range ...")
		for i := 0; i < numWorkers; i++ {
			ch <- Work{number:i}
		}
	}()

	go func () {
		count := 0
		for w := range ch {
			fmt.Println(w)
			count += 1
			if count == numWorkers {
				fmt.Println("... over")
				wg.Done()
			}
		}
	}()

	wg.Wait()

}