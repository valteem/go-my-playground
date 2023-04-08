package main

import (
	"fmt"
)

func recIntDecr(i *int) {
	*i--
	fmt.Println("Value after decrement is", *i)
	if *i > 0 {
		recIntDecr(i)
	}
}

func recSum(i int) int {
	if i == 0 {
		fmt.Println(i, "and it stops here")
		return 0
	} else {
		fmt.Println(i, "and it goes on")
		return i + recSum(i-1)
	}
}

func tailRecSum(args ...int) int {
	var s int
	if len(args) > 1 {
		s = args[1]
	} else {
		s = 0
	}
	if args[0] == 0 {
		fmt.Println(args[0], "and it stops here")
		return s
	} else {
		fmt.Println(args[0], "and it goes on")
		return tailRecSum(args[0] - 1, args[0] + s)
	}
}

func main() {
	x := 5
	y := &x
	recIntDecr(y)

	fmt.Println("Regular recursion returns", recSum(5))
	fmt.Println("Tail recursion returns", tailRecSum(5))
}
