package main

import (
	"fmt"
)

func main() {
	var prints []func()
	for i := 1; i <= 3; i++ {
		prints = append(prints, func() { fmt.Println(i) })
	}
	for _, print := range prints {
		print()
	}
}
