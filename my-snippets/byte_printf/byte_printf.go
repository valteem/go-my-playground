// https://pkg.go.dev/fmt

package main

import "fmt"

func main() {

	txt := "Some text"
	bt := []byte(txt)
	fmt.Printf("%s\n", bt)
	fmt.Printf("%v\n", txt)
	fmt.Println(bt)
	fmt.Printf("%v\n", bt) // v - default format
	fmt.Printf("%o\n", bt) // o - base 8 format (83: 123 (base 8) = 1*8*8 + 2*8 + 3 = 64 + 16 + 3 = 83 (base 10))
	t1 := "W"
	b1 := []byte(t1)
	fmt.Printf("%x %x\n", t1, b1) // x - base 16 format (hexadecimal)
	fmt.Printf("%s %o\n", t1, b1) // 

}