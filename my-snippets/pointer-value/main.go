package main

import (
	"fmt"
//	"unsafe"
)

func receiver(av int, ap *int) {

	fmt.Println(av, &av, ap, &ap)
	fmt.Println((* int)(ap))  // this just returns address of reference
	fmt.Println(*(* int)(ap)) // this retrieves value of variable
	fmt.Println(av == (*(* int)(ap)))

}

func main() {

	var x int = 1
	fmt.Println(x, &x)
	receiver(x, &x)

}