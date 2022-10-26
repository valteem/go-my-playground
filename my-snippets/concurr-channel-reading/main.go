// 10 goroutines reading from a channel 100 times

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan int, 1)
	q := make(chan bool)
	for i := 1; i <= 10; i++ {
		go func(v int) {
			counter := 0
			for {
				select {
				case <-q:
					wg.Done()
					return
				default:
					counter++
					fmt.Println("Goroutine #", v, "message", <-ch, "read", counter, "times")
				}
			}
		}(i)
	}

	for j := 1; j <= 100; j++ {
		ch <- j
	}
	q <- true
	wg.Wait()
//	close(ch)
}
