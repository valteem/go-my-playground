package main

import (
	"fmt"
	"time"
)

func fp(from string, rept int) {
	for i:=0; i<rept; i++ {
		fmt.Println(from, " : ", i)
	}
}

func main() {
	fp("dir",10)

	go fp("go1", 10)

	go func(msg string) {
        fmt.Println(msg)
    }("going")

	go fp("g02", 10)

	time.Sleep(10 * time.Second)
	fmt.Println("done")
}