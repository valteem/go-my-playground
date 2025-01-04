package main

import (
	"flag"
	"fmt"
)

var (
	left   = flag.String("left", "apples", "left-side operand")
	right  = flag.String("right", "onions", "right-side operand")
	compop = flag.String("op", "sweeter", "comparison operation")
)

func main() {
	flag.Parse()
	fmt.Printf("%s are %s than %s\n", *left, *compop, *right)

	// fmt.Printf("%s %s %s\n", *left, *compop, *right)
}
