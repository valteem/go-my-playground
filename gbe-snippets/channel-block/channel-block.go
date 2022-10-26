package main

import "fmt"

func Square(c chan int) {
	for i := 0; i <= cap(c); i++ {
	num := <- c
	fmt.Println(num*num)
	}
}

func main() {
	cl := make(chan int, 3)

	fmt.Println("starting ...")
	
	go Square(cl)

	cl <- 1
	cl <- 2
	cl <- 3
	cl <- 4
	cl <- 5

	fmt.Println("...  stopping")
}