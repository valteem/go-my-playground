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
	fmt.Println((*b1).Cover, b1.Cover) // pointer to struct is automatically dereferenced
	// https://go.dev/ref/spec#Selectors
	// As an exception, if the type of x is a defined pointer type and (*x).f is a valid selector expression
	// denoting a field (but not a method), x.f is shorthand for (*x).f.
	b2 := *b1
	b2.Cover = "skin"
	b1.Cover = "nothing" // pointer to a struct is automatically dereferenced (see above)
	b3 := NewBookValue()
	fmt.Println(b1, b2, b3)
	fmt.Println(reflect.TypeOf(b1), reflect.TypeOf(b2), reflect.TypeOf(b3))

}
