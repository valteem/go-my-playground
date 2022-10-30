package main

import (
	"fmt"
	"reflect"
)

type Book struct {
	Cover string
	Pages int
}

func main() {

	b := new(Book)
	fmt.Println(reflect.TypeOf(b)) // new() returns pointer (*main.Book)

}