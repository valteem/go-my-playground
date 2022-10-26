package main

import (
	"fmt"
	"sync"
)

const nmax = 1000

func main() {

	var wg sync.WaitGroup
	wg.Add(2) // 2 because of waiting for both sending and reading of all integers

	dst := make(chan int)
	n := 1
	go func() {
		for n <= nmax {
				dst <- n
				n++
			}
		wg.Done() // waiting for all integers to be sent to the channel
	}()
	
	go func () {
		for k:= range(dst) {
		fmt.Println(k)
		if k >= nmax {wg.Done()} // waiting for all integers to be read from the channel
		}
	}()

	wg.Wait()
}