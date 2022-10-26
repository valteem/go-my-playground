// Some explanations here:
// https://stackoverflow.com/questions/19414267/initializing-a-struct-as-a-pointer-or-not

package main

import (
	"fmt"
	"reflect"
)

type Book struct {
	Name  string
	Cover string
}

func NewBookPointer() (b *Book) { return &Book{Name: "new", Cover: "paper"} }

func NewBookValue() (b Book) { return Book{Name: "new", Cover: "paper"} }

func main() {

	b1 := NewBookPointer()
	b2 := *b1
	b2.Cover = "skin"
	b1.Cover = "nothing" // have no idea how this actually works, the concept of pointers in Go still looks pretty obscure
	b3 := NewBookValue()
	fmt.Println(b1, b2, b3)
	fmt.Println(reflect.TypeOf(b1), reflect.TypeOf(b2), reflect.TypeOf(b3))

}
