// Implement ring buffer with generics
// Mostly follows smallnest/ringbuffer mechanics

package ringbuf

import (
	"errors"
)

const (
	initialSize  = 2
	extendFactor = 2
	shrinkFactor = 2
)

var (
	ErrRingBufEmpty = errors.New("ring buffer is empty")
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

func (rb *RingBuf[T]) Write(t T) {
	if (rb.head+1)%rb.size == rb.tail {
		rb.extend()
	}
	rb.buf[rb.head] = t
	rb.head = (rb.head + 1) % rb.size
}

func (rb *RingBuf[T]) Read() (T, error) {
	if rb.tail == rb.head {
		var t T
		return t, ErrRingBufEmpty
	}
	r := rb.buf[rb.tail]
	rb.tail = (rb.tail + 1) % rb.size
	return r, nil
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
