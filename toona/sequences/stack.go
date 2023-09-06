package sequences

import (
	"errors"
	"sync"
)

type Stack [T any] struct {
	lock sync.Mutex
	arr []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{sync.Mutex{}, make([]T, 0)}
}

func (s *Stack[T]) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return len(s.arr) == 0
}

func (s *Stack[T]) Push(e T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.arr = append(s.arr, e)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var t T // returning 'zero' T value
		return t, errors.New("stack empty")
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	r := s.arr[len(s.arr) - 1]
	s.arr = s.arr[:len(s.arr) - 1]
	return r, nil
}