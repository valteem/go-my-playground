// After https://github.com/pkg/math/blob/master/uint64.go

package main

import "fmt"

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minIntN(v ...int) int {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		return minInt(v[0], v[1])
	default:
		l := len(v) / 2
		return minIntN(minIntN(v[:l]...), minIntN(v[l:]...))
	}
}

func main() {
	v := []int{5, 8, 3, 16, 1, 4, 21, 2}
	fmt.Println(minIntN(v...))
}
