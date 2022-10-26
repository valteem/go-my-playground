package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("os.Args are ")
	fmt.Println(os.Args) // all arguments. including command name
	fmt.Println(os.Args[1:]) // only 'real' arguments (omitting command name)
}