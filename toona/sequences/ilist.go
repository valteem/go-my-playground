// double linked list
// basic insert/remove operations
// arbitrary sentinels for head and tail

package sequences

import (
	"fmt"
	"sync"
)

type IList [T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
	lock sync.Mutex
}

func NewIList[T any]() *IList[T] {
	var t T
	head := NewNode[T](t, nil, nil)
	tail := NewNode[T](t, nil, nil)
	head.Previous = tail
	tail.Next = head
	return &IList[T]{
		size: 0,
		head: head,
		tail: tail,
	}
}

func (l *IList[T]) IsEmpty() bool {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.size == 0
}

func (l *IList[T]) Head() *Node[T] {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.head
}

func (l *IList[T]) Tail() *Node[T] {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.tail
}

func (l *IList[T]) Insert(elt T, prev *Node[T], next *Node[T]) *Node[T] {
	l.lock.Lock()
	defer l.lock.Unlock()
	n := NewNode[T](elt, prev, next)
	prev.Next = n
	next.Previous = n
	l.size++
	return n
}

func (l *IList[T]) Remove(n *Node[T]) T {
	l.lock.Lock()
	defer l.lock.Unlock()
	e := n.Element
	prev := n.Previous
	next := n.Next
	prev.Next = next
	next.Previous = prev
	l.size--
	return e
}

func (l *IList[T]) String() string {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.size == 0 {
		return "empty list"
	}
	s := "tail<->"
	n := l.tail.Next
	i := 0
	for i < l.size {
		s += fmt.Sprintf("%v<->", n.Element)
		i++
		n = n.Next
	}
	s += "head"
	return s
}