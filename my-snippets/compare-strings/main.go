package main

import (
	"fmt"
)

func main() {
	fmt.Println("\"abc\" > \"123\": ", "abc" > "123")
	fmt.Println("\"abc\" > \"ab\": ", "abc" > "ab")
	fmt.Println("\"abc\" > \"abcd\": ", "abc" > "abcd")
}
