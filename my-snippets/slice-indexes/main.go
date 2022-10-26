package main

import (
	"fmt"
)

var (
	posLeft = 5
	posRight = 10
)

func main() {

	s := "some plain text"
	fmt.Println(TrimString(s, posLeft, posRight)) // this first trims from left posLeft symbols from initial string, then trims from right all but posRight symbols
	fmt.Println(s[posLeft:posRight])              // this trims everything but symbols between posLeft and posRight of initial string, in one step
	fmt.Println(s[posLeft:][:posRight][:5])

}

func TrimString(s string, posBegin, posEnd int) string {

	return s[posBegin:][:posEnd] // second pair of brackets 'trims from right' the result of applying first pair of brackets ('trim from left')

} 