package sequences

import (
	"fmt"
	"sync"
)

var (
	errDequeEmpty = fmt.Errorf("empty deque")
)

type CircularDeque [T any] struct {
	arr []T
	cap int
	size int
	front int
	lock sync.Mutex
}

func NewCircularDeque[T any](n int) *CircularDeque[T] {
	if n < 1 {
		return nil
	}
	return &CircularDeque[T]{
		arr: make([]T, n),
		cap: n,
		size: 0,
		front: 0,
	}
}

func (q *CircularDeque[T]) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.size == 0
}

func (q *CircularDeque[T]) IsFull() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.size == q.cap
}

func (q *CircularDeque[T]) resize() {
	n := make([]T, 2 * q.cap)
	i := q.front
	j := 0
	count := 0
	for {
		count++
		if count > q.size {
			break
		}
		n[j] = q.arr[i]
		i = (i + 1) % q.cap
		j ++
	}
	q.cap = q.cap * 2
	q.arr = n
	q.front = 0
}

func (q *CircularDeque[T]) AddFirst(e T) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.size == q.cap {
		q.resize()
	}
	q.front = (q.front - 1) % q.cap
	if q.front < 0 {
		q.front = q.cap + q.front // wrap around
	}
	q.arr[q.front] = e
	q.size++
}

func (q *CircularDeque[T]) DeleteFirst() (T, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.size == 0 {
		var t T
		return t, fmt.Errorf("deque empty")
	}
	r := q.arr[q.front]
	q.size--
	q.front = (q.front + 1) % q.cap
	return r, nil
}

func (q *CircularDeque[T]) AddLast(e T) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.size == q.cap {
		q.resize()
	}
	i := (q.front + q.size) % q.cap
	q.arr[i] = e
	q.size++
}

func (q *CircularDeque[T]) DeleteLast() (T, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.size == 0 {
		var t T
		return t, fmt.Errorf("deque empty")
	}
	i := (q.front + q.size - 1) % q.cap
	r := q.arr[i]
	q.size--
	return r, nil
}

func (q *CircularDeque[T]) String() string {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.size == 0 {
		return "empty deque"
	}
	i := q.front
	output := fmt.Sprintf("first<-%v", q.arr[i])
	for {
		i = (i + 1) % q.cap
		if i == q.front {
			break
		}
		output += fmt.Sprintf("<-%+v", q.arr[i])
	}
	output += "<-last"
	return output
}