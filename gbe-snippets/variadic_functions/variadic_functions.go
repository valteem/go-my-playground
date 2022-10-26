package main

import "fmt"

func vfsum(args ...int) int {
	rsum := 0
	for _, val := range args {
		rsum += val
	} 
	return rsum
}

func main() {
	
	fmt.Println("Sum of separate arguments: ",vfsum(10,30,40))

	slcInt := []int{10,30,40}
	fmt.Println("Sum of slice elements: ", vfsum(slcInt...))
}