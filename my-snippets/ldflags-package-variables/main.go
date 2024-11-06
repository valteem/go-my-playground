package main

import (
	"fmt"
)

var (
	outputGlobal = "potatoes"
)

func main() {

	outputLocal := "onions"

	fmt.Println(outputGlobal + " and " + outputLocal)

}
