package reuse_test

import (
	"fmt"
	"testing"

	"github.com/valteem/reuse"
)

func TestDoubleLinkedList(t *testing.T) {

	m := reuse.Create()
	e1 := m.PushBack("s1")
	e2 := m.PushFront(11)
	m.InsertBefore(true, e1)
	m.InsertAfter(true, e2)

	fmt.Println(m)

}