package sequences

import (
	"fmt"
	"sync"
)

type CircularQueue [T any] struct {
	arr []T
	cap int
	front int
	size int
	lock sync.Mutex
}

func NewCircularQueue[T any](n int)  *CircularQueue[T] {
	if n < 1 {
		return nil
	}
	return &CircularQueue[T]{
		arr: make([]T, n),
		cap: n,
		front: 0,
		size: 0,
	}
}

func (q *CircularQueue[T]) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.size == 0
}

func (q *CircularQueue[T]) IsFull() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.size == q.cap
}

func (q *CircularQueue[T]) Enqueue(e T) bool {
	if q.IsFull() {
		return false
	}
	q.lock.Lock()
	defer q.lock.Unlock()
	available := (q.front + q.size) % q.cap
	q.arr[available] = e
	q.size++
	return true
}

func (q *CircularQueue[T]) Dequeue() T {
	if q.IsEmpty() {
		var t T
		return t 
	}
	q.lock.Lock()
	defer q.lock.Unlock()
	t := q.arr[q.front]
	q.size--
	q.front = (q.front + 1) % q.cap
	return t
}

func (q *CircularQueue[T]) String() string {
	if q.IsEmpty() {
		return "empty queue"
	}
	q.lock.Lock()
	defer q.lock.Unlock()
	i := q.front
	output := fmt.Sprintf("head<-%v", q.arr[i])
	for {
		i = (i + 1) % q.cap
		if i == q.front {
			break
		}
		output += fmt.Sprintf("<-%+v", q.arr[i])
	}
	output += "<-tail"
	return output
}