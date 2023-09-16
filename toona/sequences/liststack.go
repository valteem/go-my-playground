package sequences

import (
	"sync"
)

type ListStack [T any] struct {
	list *DoubleLinkedList[T]
	lock sync.Mutex
}

func NewListStack[T any]() *ListStack[T] {
	return &ListStack[T]{list: NewDoubleLinkedList[T]()}
}

func (s *ListStack[T]) IsEmpty() bool {
	return s.list.IsEmpty()
}

func (s *ListStack[T]) Push(e T) {
	s.list.AddToHead(e)
}

func (s *ListStack[T]) Pop() (T, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	r, e := s.list.RemoveFromHead()
	if e != nil {
		e = errStackEmpty
	} 
	return r, e
}

func(s *ListStack[T]) String() string {
	return s.list.String()
}