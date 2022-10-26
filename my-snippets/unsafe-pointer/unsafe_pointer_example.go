package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := "sometext"
	b := []byte(s)

	fmt.Println(b)
	fmt.Println(&b)
	fmt.Println(unsafe.Pointer(&b))
	fmt.Println((* string)(unsafe.Pointer(&b)))
	fmt.Println(*(* string)(unsafe.Pointer(&b)))

	fmt.Println(s)
	fmt.Println(&s)                 // this ...
	fmt.Println(unsafe.Pointer(&s)) // ... and this show exactly the same ...
	fmt.Println((* string)(&s))     // ... and this too
	
}