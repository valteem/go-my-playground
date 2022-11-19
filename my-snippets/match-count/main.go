package main

import (
	"fmt"
)

func matchLetter(str string, letter byte) int {
	count := 0
	for ; len(str) > 0; str = str[1:] {
		if str[0] == letter {
			count++
		}
	}
	return count
}

func main() {

	s := "many apples"
	fmt.Println(matchLetter(s, 'a'))
	
}