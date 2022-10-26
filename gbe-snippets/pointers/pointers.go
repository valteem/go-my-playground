package main

import "fmt"

func zeroval(num int) {
	num = 0
}

func zeroptr(num *int) {
	*num = 0
}

func main() {
	val := 1
	zeroval(val)
	fmt.Println("After passing by value: ", val)
	zeroptr(&val)
	fmt.Println("After passing by ref: ", val)
	fmt.Println("Ref: ", &val)
}