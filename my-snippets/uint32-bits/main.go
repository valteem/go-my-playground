package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int64
	for i := 1; i < 32; i++ {
		n = 1 << i
		fmt.Println(i, n, strconv.FormatInt(n,2))
	}

	x := int64(0b11111111) // 255 = 2 ^ 8 - 1
	for i := 0; i < 8; i++ {
	    fmt.Println(x, x>>i, strconv.FormatInt(x>>i, 2))
	}
}