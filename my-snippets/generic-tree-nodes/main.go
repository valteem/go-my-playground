package main

import (
	"fmt"
)

type Node[T any] struct {
	Value T
	Left *Node[T]
	Right *Node[T]
}

func main() {

	var ti Node[int]
	ti.Value = 1
	ti.Left = nil
	ti.Right = nil
	fmt.Println(ti)

	var ts Node[string]
	ts.Value = "new"
	fmt.Println(ts)
}