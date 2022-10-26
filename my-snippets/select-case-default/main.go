package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	count := 0

	term := make(chan bool)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-term:
				fmt.Printf("Quit after %d counts\n", count)
				return
			default:
				count ++
				fmt.Printf("Count %d\n", count)
			}
		}
	}()

	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Terminating ...")
		term <- true
	}()

	wg.Wait()
	
}