package main

import (
	"fmt"
	"context"
)

func main() {

	fmt.Println("Beginning of the example")
	defer fmt.Println("End of the example")

	ctx1, cancel1 := context.WithCancel(context.Background())
	defer cancel1() // cancel call is deferred for ctx1

	ctx2, cancel2 := context.WithCancel(context.Background())
	cancel2() // cancel call executed immediately

	select {
	case <-ctx1.Done():
		fmt.Println("If cancel call is deferred ...")
	default:
		fmt.Println("... then default clause is executed")
	}

	select {
	case <-ctx2.Done():
		fmt.Println("If cancel call is made immediately ...")
	default:
		fmt.Println("... then default clause has no chance to execute")
	}

 }