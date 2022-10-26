package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	name string
	width float64
	height float64
}

type circle struct {
	name string
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2.0 * (r.width + r.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2.0 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("Geometry: ",g, " Area: ",g.area()," Perimeter: ", g.perim())
}

func main() {
	rct := rect{name:"Green rectangle",width:2.0, height:7.0}
	crc := circle{name:"Yelllow circle",radius:1.0}

	measure(rct)
	measure(crc)

}