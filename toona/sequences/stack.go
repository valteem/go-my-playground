package sequences

import (
	"errors"
	"sync"
)

var (
	errStackEmpty = errors.New("stack empty")
	errMaxDepth = errors.New("max stack depth achieved")
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
		return t, errStackEmpty
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	r := s.arr[len(s.arr) - 1]
	s.arr = s.arr[:len(s.arr) - 1]
	return r, nil
}

type AllocStack [T any] struct {
	lock sync.Mutex
	arr []T
	depth int
	maxDepth int
}

func NewAllocStack[T any](maxDepth int) *AllocStack[T] {
	return &AllocStack[T]{sync.Mutex{}, make([]T, maxDepth), 0, maxDepth}
}

func (a *AllocStack[T]) IsEmpty() bool {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.depth == 0
}

func (a *AllocStack[T]) Push(e T) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	if a.depth == a.maxDepth {
		return errMaxDepth
	} else {
		a.depth++
		a.arr[a.depth - 1] = e
		return nil
	}
}

func (a *AllocStack[T]) Pop() (T, error) {
	if a.IsEmpty() {
		var t T // returning 'zero' T value
		return t, errStackEmpty
	}
	a.lock.Lock()
	defer a.lock.Unlock()
	r := a.arr[a.depth - 1]
	a.depth--
	return r, nil
}