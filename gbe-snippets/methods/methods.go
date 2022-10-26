package main

import "fmt"

type rect struct {
	width float64
	height float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2.0 * (r.width + r.height)
}

func main() {

	var rct rect
	rct.width = 1.0
	rct.height = 3.0
	fmt.Println(rct.area())
	fmt.Println(rct.perim())

	rptr := &rct
	fmt.Println(rptr.area())
	fmt.Println(rptr.perim())
}