package main

import (
	"fmt"
)

func main() {

	s := []int{7, 5, 18, 16, 3, 15, 18, 24, 23, 17}
	np, ap := qsPartitionInt(s, 0, 10, 5)
	fmt.Println(np, ap, s)
}