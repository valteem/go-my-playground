package main

import "fmt"

func main() {
	b := 11
	fmt.Println("Function scope: b = " , b)
	c := 31
	fmt.Println("Function scope: c = " , c)

//	for i := 0; i < 1; i++ { // no loop is needed, just braces will do
	{
		b := 12
		fmt.Println("Braces scope: b = ", b)
		fmt.Println("Braces scope - no local declaration: c = ", c)
	}

	fmt.Println("Function scope (after braces):  b = " , b)
}