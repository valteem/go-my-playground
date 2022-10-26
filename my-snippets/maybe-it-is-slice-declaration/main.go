package main

import (
	"fmt"
)

func main() {

	var empBytes [8]byte
	empBytes[1] = 1
	var str [4]string
	str[1] = "1"
	var i [3]uint32
	sbytes := []byte("abc")

	fmt.Println(empBytes)
	fmt.Println(str)
	fmt.Println(i)
	fmt.Println(sbytes)
}