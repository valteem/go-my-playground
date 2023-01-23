// https://stackoverflow.com/questions/70358216/how-can-i-access-a-struct-field-with-generics-type-t-has-no-field-or-method
package main

import (
	"fmt"
)

type Shape interface{
//	Point | Rect // cannot uset type Shape outside a type constraint
	GetX() int
}

type Point struct {
	X int
	Y int
}

func (p Point) GetX() int {
	return p.X
}

type Rect struct {
	X int
	Y int
}

func (r Rect) GetX() int {
	return r.X
}

// func GetX[P Point | Rect] (p P) {
// 	return p.X
// }

func PrintX(s Shape) string {
	return fmt.Sprintf("%d", s.GetX())
}

func main() {

	p := Point{X: 1, Y: 2}
	r := Rect{X: 1, Y: 2}
	// fmt.Println(GetX(p), GetX(r))
	fmt.Println(PrintX(p), PrintX(r))

}
