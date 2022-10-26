package main

import (
	"fmt"
	"math"
)

const s string = "Outer constant"

func main() {
	var s1 string = "Local prefix + " + s
	fmt.Println(s1)

	var locNum1 float64 = 16.0
	fmt.Println(math.Sqrt(locNum1))

}
