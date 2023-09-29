// Python-style generator implemented with channels

package reuse

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