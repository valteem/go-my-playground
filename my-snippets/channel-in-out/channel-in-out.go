package main

import 	(
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	conn := make(chan string, 2)
	state :=make(chan int, 2)
	
	go func () {
		conn <- "in"
		fmt.Println("Sending in")
	}()

	go func() {
		fmt.Println("read once:",<-conn)
		state <- 1
	}()

	go func() {
		conn <- "out"
		fmt.Println("Sending out")
	}()

	go func() {
	 fmt.Println("read twice:", <-conn)
	 state <- 1
	}()

	go func() {
		defer wg.Done()
		count := 0
		for input := range state {
			count += input
			if input == 2 {
				return
			}
		}

	}()

	wg.Wait()

}
