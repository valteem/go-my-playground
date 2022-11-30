package main

import (
	"fmt"
)

func main() {

	const (
		indexMax = 10
		indexBreak = 5
	)

	for i := 0; i < indexMax; i++ {
		for j := 0; j < indexMax; j++ {
			if i == indexBreak {continue}
			if j == indexBreak {continue}
			fmt.Println(i, j)
		}
	}
}