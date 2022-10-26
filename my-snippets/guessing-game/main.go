package main

import (
	"fmt"
	"sort"
)

const nmax = 100

func main() {
	var s string
	fmt.Printf("Pick an integer from 0 to %d:\n", nmax)
	answer := sort.Search(nmax, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}
