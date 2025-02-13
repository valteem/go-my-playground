// Implement ring buffer with generics
// Mostly follows smallnest/ringbuffer mechanics

package ringbuf

import (
// "sync"
)

const (
	initialSize  = 2
	extendFactor = 2
	shrinkFactor = 2
)

type RingBuf[T any] struct {
	buf []T

	// First version does not allow multiple access
	// readCond sync.Cond // https://victoriametrics.com/blog/go-sync-cond/

	head int
	tail int
	size int
}

func NewRingBuf[T any]() *RingBuf[T] {
	return &RingBuf[T]{buf: make([]T, initialSize), size: initialSize}
}

func (rb *RingBuf[T]) Add(t T) {

	if (rb.head+1)%rb.size == rb.tail {
		rb.extend()
	}
	rb.buf[rb.head] = t
	rb.head = (rb.head + 1) % rb.size

}

func (rb *RingBuf[T]) Buf() []T {

	output := make([]T, rb.size)

	i := rb.tail
	count := 0
	for count < rb.size {
		output[count] = rb.buf[i]
		count++
		i = (i + 1) % rb.size
	}

	return output

}

func (rb *RingBuf[T]) extend() {

	newBuf := make([]T, rb.size*extendFactor)

	i := rb.tail
	count := 0
	for count < rb.size {
		newBuf[count] = rb.buf[i]
		count++
		i = (i + 1) % rb.size
	}

	rb.size = rb.size * extendFactor
	rb.tail = 0
	rb.head = count - 1
	rb.buf = newBuf

}
