package main

import "fmt"

func printInput(ch1, ch2, quit chan int) {
	var rec1 int
	var rec2 int
	for {
	select {
	case rec1 = <-ch1:
		fmt.Println("Channel 1 receives",rec1)
	case rec2 = <-ch2:
		fmt.Println("Channel 2 receives", rec2)
	case <- quit:
		return
	}
    }
}

func main() {

	c1 := make(chan int)
	c2 := make(chan int)
	quit := make(chan int)

	numMessage := 10

	go func() {
		for i:= 1; i <= numMessage; i++ {
			c1 <- i
		}
		quit <- 0
	}()
	go func() {
		for j:= -numMessage; j <= -1; j++ {
			c2 <- j
		}
		quit <- 0
	}()

	printInput(c1, c2, quit)
}