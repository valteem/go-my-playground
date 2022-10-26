package main

import "fmt"

type Child struct {}

// Child has method Print()
func (c *Child) Print() {
	fmt.Println("Child")
}

type Parent struct {
	Child
}

// Parent has method Print() too
func (p *Parent) Print() {
	fmt.Println("Parent")
}

func main() {

	var x Parent
	x.Print()		// uses Parent definition of Print()
	x.Child.Print()	// uses Child definition of Print()
	
}