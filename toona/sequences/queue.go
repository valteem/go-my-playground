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

func (q *CircularQueue[T]) Enqueue(e T) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.size == q.cap {
		q.resize()
	}
	available := (q.front + q.size) % q.cap
	q.arr[available] = e
	q.size++
}

func (q *CircularQueue[T]) Dequeue() (T, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.size == 0 {
		var t T
		return t, fmt.Errorf("empty queue")
	}
	t := q.arr[q.front]
	q.size--
	q.front = (q.front + 1) % q.cap
	return t, nil
}

func (q *CircularQueue[T]) String() string {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.size == 0 {
		return "empty queue"
	}
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

func (q *CircularQueue[T]) resize() {
	n := make([]T, 2 * q.cap)
	idxOld := q.front
	idxNew := 0
	count := 0
	for {
		count++
		if count > q.size {
			break
		}
		n[idxNew] = q.arr[idxOld]
		idxOld = (idxOld + 1) % q.cap
		idxNew ++
	}
	q.cap = q.cap * 2
	q.arr = n
	q.front = 0
}