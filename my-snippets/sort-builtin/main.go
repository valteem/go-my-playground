package main

import (
	"sort"
)

type S []int

func (s S) Len() int {
	return len(s)
}

func (s S) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s S) Less(i, j int) bool {
	return s[i] < s[j]
}

func main() {

	s:= S{17, 4, 11, 7, 19, 3, 14, 8, 15, 44, 33, 27, 51, 8}
	sort.Sort(s)
}