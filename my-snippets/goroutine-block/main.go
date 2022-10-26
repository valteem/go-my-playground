package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch := make(chan string)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		startTime := time.Now()
		ch <- "message" // goroutine blocks here, waiting for another goroutine wake up and perform reading from channel
		fmt.Println(time.Since(startTime), "elapsed since send operation") // shows that this goroutine is blocked until another goroutine performs read operation
		wg.Done()
	}()

	go func() {
		time.Sleep(1 * time.Second)
		m := <-ch // this read happens after some time, thus blocking another goroutine from moving past send operation
		fmt.Println("Message received:", m)
		wg.Done()
	}()

	wg.Wait()

}