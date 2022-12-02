package main

import (
	"fmt"
)

// standarf 'pdqsort' partition, adopted for data slice of int 
func qsPartitionInt(data []int, a, b, pivot int) (newpivot int, alreadyPartitioned bool) {
	data[a], data[pivot] = data[pivot], data[a]
	i, j := a + 1, b - 1
	for i <= j && data[i] < data[a] {
		i++
	}
	for i <= j && !(data[j] < data[a]) {
		j--
	}

	if i > j {
		data[j], data[a] = data[a], data[j]
		return j, true
	}

	return pivot, false
}

func main() {

	s := []int{7, 5, 11, 3, 15, 18, 24, 28, 23, 17}
	np, ap := qsPartitionInt(s, 0, 10, 4)
	fmt.Println(np, ap)
}