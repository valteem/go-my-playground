// https://stackoverflow.com/a/40951013

package main

import (
	"fmt"
)


type SomeType interface {
	SomeFunc(int, float32, string) (int, string)
}

type SomeTypeImpl struct {
}

func (sti SomeTypeImpl) SomeFunc(p1 int, p2 float32, p3 string) (int, string) {
	return p1 + int(p2), p3 + " suffix"
}

func Wrap(st SomeType, p1 int, p2 float32, p3 string) {
	fmt.Println(st.SomeFunc(p1, p2, p3))
}

func main() {
	s := SomeTypeImpl{}
	Wrap(s, 1, 2.5, "text")
}