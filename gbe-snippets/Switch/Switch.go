package main

import (
	"fmt"
	"time"
)

func main() {

	tm := time.Now().Weekday()
	fmt.Println("Switch with default")
	switch tm {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}
	fmt.Println("Switch without defaault")
	switch tm {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	}

	th := time.Now()
	switch {
	case th.Hour() < 12:
		fmt.Println("Before noon")
	default:
		fmt.Println("After noon")
	}

	whatAmI := func( i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am Bool")
		case int:
			fmt.Println("I am Int")
		default:
			fmt.Printf("I am type %T\n", t)
		}
	}

	whatAmI(false)
	whatAmI(11)
	whatAmI(1.1)
	whatAmI("Hello")

}