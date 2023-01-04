package main

import (
	"fmt"
)

func changeSlice(slc []int, singleValue int, ptr *int) {
	for i, v := range slc {
		slc[i] = 2 * v
	}
	singleValue *= 2
	*ptr *= 2
}

func reshapeSlice(slc []string) {
	for ; len(slc) > 0; slc = slc[1:] {
		p := &slc[0]
		*p = "value"
		if len(slc) > 0 {
			fmt.Println("reshaped", len(slc), &slc[0])
		}
	}
}

func refillSlice(slc []string) {
	fmt.Println("inner", &slc[0])
	for i, _ := range slc {
		slc[i] = "value"
	}
}

func main() {
	s := []int{1, 5, 7}
	i := 72
	j := 72
	changeSlice(s, i, &j)
	fmt.Println(s, i, j)
	
	s1 := make([]string, 3)
	fmt.Println("outer before", &s1[0])
	reshapeSlice(s1)
    fmt.Println(s1)
	fmt.Println("outer after", &s1[0])

	s2 := make([]string, 3)
	fmt.Println("outer before", &s2[0])
	refillSlice(s2)
	fmt.Println("outer after", &s2[0])
    fmt.Println(s2)
}