// Build with -tags=odd

//go:build odd

package main

import "fmt"

func Output() {
	fmt.Println("Odd case")
}
