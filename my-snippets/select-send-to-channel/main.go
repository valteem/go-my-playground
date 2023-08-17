package main

import (
	"fmt"
	"sync"
)

type UBuffer[T any] struct {
	c chan T
	backlog []T
}

var (
	numWrites = 5
)

func (u *UBuffer[T]) Load() {

//	n := new(T)
	if len(u.backlog) > 0 {
		select {
		case u.c <- u.backlog[0]:
//			u.backlog[0] = *n
			u.backlog = u.backlog[1:] // this works because UBuffer as a whole and its backlog are passed by reference
		default:
		}
	}

}

func main() {

	u := UBuffer[int]{c: make(chan int, 1), backlog: []int{1, 2, 3, 4, 5}}

	var wg sync.WaitGroup
	wg.Add(1) 

	go func() {
		defer wg.Done()
		for i := 0; i < numWrites; i++  {
			u.Load()
			v := <-u.c
			fmt.Println(v)
		}
	}()

	wg.Wait()

}