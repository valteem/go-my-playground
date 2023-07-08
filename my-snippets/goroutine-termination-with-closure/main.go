package main

import (
	"fmt"
	"time"
)

var done = []bool{false, false, false, false, false}
var count = []int{0, 0, 0, 0, 0}

func worker(i int) {
	fmt.Println(i, "worker starting ...")
	for {
		if done[i] {
			return
		}
		count[i] ++
		fmt.Println(i, "... working...")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	for i := 0; i < 5; i++ {
		go worker(i)
	}
	time.Sleep(10 * time.Second)
	for i := range done{
		done[i] = true
	}
	fmt.Println(count)
}