package main

import "fmt"

type Tyre struct {}

func (t Tyre) Whoosh() {
	fmt.Println("Whoosh ...")
}

type Wheel struct {
	Tyre
}

type Car struct {
	Wheel
}

func main() {

	t:= new(Tyre)
	w := Wheel{Tyre: *t}
	c := Car{Wheel: w}
	c.Whoosh() //c.Wheel.Tyre.Whoosh()

}