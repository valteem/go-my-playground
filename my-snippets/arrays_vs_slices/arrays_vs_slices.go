package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a [5]int
	b := [5]int{10, 20, 30, 40, 50}
	s := make([]int, 5)

	fmt.Println(reflect.ValueOf(a).Kind())
	fmt.Println(reflect.ValueOf(b).Kind())
	fmt.Println(reflect.ValueOf(s).Kind())
}