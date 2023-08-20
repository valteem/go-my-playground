package reuse_test

import (
	"fmt"
	"testing"
)

type A struct {
	name string
	count int
}

func NewA() *A {
	return &A{}
}

func TestStructReassign(t *testing.T) {
	a1 := &A{"item1", 1}
	a2 := &A{"item2", 2}
	a3 := &A{}
	a4 := A{}

	*a3 = *a1
	fmt.Println(a1.name, a1.count, a3.name, a3.count)

	*a1 = *a2
	fmt.Println(a1.name, a1.count, a2.name, a2.count)

	a2 = a3
	fmt.Println(a2.name, a2.count, a3.name, a3.count)

	*a1 = *NewA()
	fmt.Println(a1.name, a1.count)

	a4 = *a2
	fmt.Println(a3.name, a4.count)

}