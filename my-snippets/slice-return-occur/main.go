package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

// This section is solely dedicated to how int(uint(n)>>1) works
	var k int64 
	for k= 1; k <= 32; k++ {
		fmt.Println(k, int(uint(k)>>1), strconv.FormatInt(k, 2), strconv.FormatInt(int64(uint(k)>>1), 2))
	}

// This is search example
	nums := []int{11, 22, 33, 44, 55}

	lv := 44

	index := sort.Search(len(nums), func(i int) bool { return nums[i] >= lv }) // inline function is used by sort.Search() to compare two elements
// sort.Search() assumes the slice is already sorted

	fmt.Println(index, nums[index])
}
