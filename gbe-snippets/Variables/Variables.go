package main

import "fmt"

func main() {
	var b, c int = 2, 3
	fmt.Println(b, c)

	var txt = "simple text"
	fmt.Println(fmt.Println(txt)) // outer Println returns number of bytes printed and an error message
	// from inner Println

	var text string
	fmt.Println(text)
}
