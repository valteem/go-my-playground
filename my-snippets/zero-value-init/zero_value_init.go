package main

import "fmt"

func main() {

	var i int
	var s string

	type combo struct {
		txt string
		num int
	}
	var c combo

	fmt.Println("i = ", i, "s = ", s, "c.txt = ", c.txt, "c.num = ", c.num)
}