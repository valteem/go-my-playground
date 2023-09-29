// Python-style generator implemented with channels

package main

import (
	"fmt"
)

type ContainerIterable[T any] struct {
	content []T
}

func NewContainerIterable[T any]() *ContainerIterable[T] {
	return &ContainerIterable[T]{content: make([]T, 0)}
}

func (c *ContainerIterable[T]) Iterate() chan T {
	ch := make(chan T)
	go func() {
		for _, v := range c.content {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

func (c *ContainerIterable[T]) Add(e T) {
	c.content = append(c.content, e)
}

func main() {
	c := NewContainerIterable[int]()

	for i := 0; i < 5; i++ {
		c.Add(i)
	}

	r := make([]int, 0, 5)
	for v := range c.Iterate() {
		r = append(r, v)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i, r[i], i == r[i])
	}
}