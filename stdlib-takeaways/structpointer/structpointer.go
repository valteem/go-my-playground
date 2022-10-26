package structpointer

import (
	"fmt"
)

type BufStr struct {
	buf []string
}

func (b *BufStr) Reset(buf []string) {
	fmt.Printf("Resetting BufStr at address %p\n", b)
	*b = BufStr{buf: buf}
}

func (b BufStr) ResetByVal (buf []string) {
	fmt.Printf("Resetting BufStr at address %p\n", &b)
	b = BufStr{buf: buf} // does not change anything, cause struct variable is passed by value
}

func NewBufStrPtr (buf []string) *BufStr {
	b := BufStr{buf: buf}
	fmt.Printf("Created BufStr at address %p\n", &b)
	return &b
}

func NewBufStrVal (buf []string) BufStr {
	b := BufStr{buf: buf}
	fmt.Printf("Created BufStr at address %p\n", &b)
	return b
}