package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	nmax = 100
	num = 15
)

func main() {

	s1 := make([]int, num)
	s2 := make([]int, num)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < num; i++ {
		s1[i] = rand.Intn(nmax)
		s2[i] = s1[i]
	}
	
	fmt.Println(s1)
	smartSort(s1)
	fmt.Println(s1)
	bruteSort(s2)
	fmt.Println(s2)
	
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