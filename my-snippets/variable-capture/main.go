package main

import (
	"fmt"
	"time"
)

func main() {
	values := []string{"a", "b", "c"}
	for _, val := range values {
		fmt.Println("range:", val)
		valInner := val
		go func() {
			fmt.Println("goroutine", val, valInner) // last value ('c') is captured, fixed in Go 1.22, use new 'inner' variable in older versions
		}()
	}
	time.Sleep(1 * time.Second)	
}