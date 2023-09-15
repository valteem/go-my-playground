package sequences

import (
	"fmt"
	"sync"
)

var (
	errListEmpty = fmt.Errorf("list empty")
)

type Node [T any] struct {
	Element T
	Previous *Node[T]
	Next *Node[T]
}

func NewNode[T any](elt T, previous *Node[T], next *Node[T]) *Node[T] {
	return &Node[T]{
		Element: elt,
		Previous: previous,
		Next: next,
	}
}

type DoubleLinkedList [T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
	lock sync.Mutex
}

func NewDoubleLinkedList[T any]() *DoubleLinkedList[T] {
	return &DoubleLinkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (l *DoubleLinkedList[T]) IsEmpty() bool {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.size == 0
}

func (l *DoubleLinkedList[T]) AddToHead(element T) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.size == 0 {
		n := NewNode[T](element, nil, nil)
		l.head = n
		l.tail = n		
	} else {
		n := NewNode[T](element, l.head, nil)
		l.head.Next = n
		l.head = n
	}
	l.size++
}

func (l *DoubleLinkedList[T]) AddToTail(element T) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.size == 0 {
		n := NewNode[T](element, nil, nil)
		l.head = n
		l.tail = n		
	} else {
		n := NewNode[T](element, nil, l.tail)
		l.tail.Previous = n
		l.tail = n
	}
	l.size++
}

func (l *DoubleLinkedList[T]) RemoveFromHead() (T, error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.size == 0 {
		var t T
		return t, errListEmpty 
	} else {
		r := l.head.Element
		l.head = l.head.Previous
		if l.head != nil {
			l.head.Next = nil
		}
		l.size--
		return r, nil
	}
}

func (l *DoubleLinkedList[T]) RemoveFromTail() (T, error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.size == 0 {
		var t T
		return t, errListEmpty
	} else {
		r := l.tail.Element
		l.tail = l.tail.Next
		if l.tail != nil {
			l.tail.Previous = nil
		}
		l.size--
		return r, nil
	}
}

func (l *DoubleLinkedList[T]) String() string {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.size == 0 {
		return "empty list"
	}
	r := "head<->"
	i := 0
	n := l.head
	for i < l.size {
		r += fmt.Sprintf("%+v<->", n.Element)
		n = n.Previous
		i++
	}
	r += "tail"
	return r
}