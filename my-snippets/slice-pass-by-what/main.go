package main

import "log"

func changeSlice(slc []int, singleValue int, ptr *int) {
	for i, v := range slc {
		slc[i] = 2 * v
	}
	singleValue *= 2
	*ptr *= 2
}

func main() {
	s := []int{1, 5, 7}
	i := 72
	j := 72
	changeSlice(s, i, &j)
	log.Println(s, i, j)
}