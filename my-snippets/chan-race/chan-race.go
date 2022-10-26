package main

import (
	"fmt"
	"strconv"
	"sync"
)

func signalReceiver1st(ch <-chan string, nch chan<- int) {
	for msg := range ch {
		fmt.Printf("1st receiver %s\n", msg)
		nch <- 1
	}
}

func signalReceiver2nd(ch <-chan string, nch chan<- int) {
	for msg := range ch {
		fmt.Printf("2nd receiver %s\n", msg)
		nch <- 1
	}
}

func signalCount(nch <-chan int, maxCount int, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	for input := range nch {
		count += input
		if count == maxCount {
			fmt.Println("All signals read")
			return
		}
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	maxIter := 10

	c := make(chan string)
//	defer close(c) // https://stackoverflow.com/questions/8593645/is-it-ok-to-leave-a-channel-open
	n := make(chan int)
//	defer close(n) // https://stackoverflow.com/questions/8593645/is-it-ok-to-leave-a-channel-open

	go signalReceiver1st(c, n)
	go signalReceiver2nd(c, n)
	go signalCount(n, maxIter, &wg)

	for count := 1; count <= maxIter; count++ {
		msg := "signal " + strconv.Itoa(count)
		c <- msg
	}
	
	wg.Wait()
}