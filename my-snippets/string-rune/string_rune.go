package main

import "fmt"

func main() {

	str := "newstring"
	fmt.Printf("Iterate over a string by UTF (?) code:\n")
	for i, s := range str {
		fmt.Printf("Symbol %d starts at position %d\n", s, i)
	}
	fmt.Printf("Iterate over a string by symbol:\n")
	for i, s := range str {
		fmt.Printf("Symbol %c starts at position %d\n", s, i)
	}
	fmt.Println("Iterate over a string as a slice of bytes")
	for j := 0; j < len(str); j++ {
		fmt.Println(j, str[j], rune(str[j]), string(rune(str[j])))
	}

	fmt.Println([]rune(str))

}
