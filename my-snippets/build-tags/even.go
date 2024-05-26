// Build with -tags=even

//go:build even

package main

import "fmt"

func Output() {
	fmt.Println("Even case")
}
