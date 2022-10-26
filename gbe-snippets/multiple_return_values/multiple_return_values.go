package main

import "fmt"

func multretval(arg1 int, arg2 int) (int,int) {
	return arg1 + arg2, arg1 - arg2
}

func main() {
	a1 := 7
	a2 := 8
	fmt.Println(multretval(a1,a2))
}