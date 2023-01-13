// https://www.reddit.com/r/golang/comments/106hi38/comment/j3iwogn/?utm_source=share&utm_medium=web2x&context=3
package main

import (
	"fmt"
	"sync"
)

func main() {
	foo := 0

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			foo += 1
		}()
	}

	wg.Wait()

	fmt.Printf("%d", foo)
}