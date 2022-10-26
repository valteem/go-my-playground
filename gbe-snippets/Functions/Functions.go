package main

import "fmt"

func add(term1 int, term2 int ) int {
	return term1 + term2
}

func add3 (term1 int, term2 int, term3 int) int {
	return term1 + term2 + term3
}

func main() {
	a1 := 1
	a2 := 11
	a3 := 21
	
	fmt.Println("Adding two arguments result: ", add(a1,a2))
	fmt.Println("Adding three arguments result: ", add3(a1,a2,a3))

}