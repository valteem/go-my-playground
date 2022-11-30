package main

import (
	"fmt"
)

func main() {

	s:= []int{17, 4, 11, 7, 19, 3, 14, 8, 15, 44, 33, 27, 51, 8}
	
	smartSort(s)
	fmt.Println(s)
	
}

func bruteSort(s []int) {
	for j := 1; j < len(s); j++ {
		for i := j; i > 0; i-- {
			if s[i] < s[i-1] {s[i], s[i-1] = s[i-1], s[i]}			
		}
	}
}

func smartSort(s []int) {
	for j := 1; j < len(s); j++ {
		x := s[j]
		i := j - 1
		for (i > -1) && s[i] > x {
			s[i+1] = s[i]
			i--
		}
		s[i+1] = x
	}
}