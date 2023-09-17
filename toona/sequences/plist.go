// positional list

package sequences

import (
	"fmt"
	"sync"
)

var (
	ErrWrongCont = fmt.Errorf("position does not belong to the list")
	ErrOutOfBoundOrRemoved = fmt.Errorf(("position out of list boundaries or removed from list"))
)

type Position [T any] struct {
	cont *IList[T]
	node *Node[T]
	lock sync.Mutex
}

func NewPosition[T any](cont *IList[T], node *Node[T]) *Position[T] {
	return &Position[T]{cont: cont, node: node}
}

func (p *Position[T]) Element() T {
	p.lock.Lock()
	defer p.lock.Unlock()
	return p.node.Element
}

type PList [T any] struct {
	list *IList[T]
	lock sync.Mutex
}

func NewPList[T any]() *PList[T] {
	return &PList[T]{list: NewIList[T]()}
}

func (l *PList[T]) validate(pos *Position[T]) (*Node[T], error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if pos.cont != l.list {
		return nil, ErrWrongCont
	}
	if pos.node.Next == nil {
		return nil, ErrOutOfBoundOrRemoved
	}
	return pos.node, nil
}

func (l *PList[T]) position(node *Node[T]) *Position[T] {
	l.lock.Lock()
	defer l.lock.Unlock()
	if node == l.list.tail || node == l.list.head {
		return nil
	}
	return NewPosition[T](l.list, node)
}

func (l *PList[T]) Head() *Position[T] {
	l.lock.Lock()
	defer l.lock.Unlock()
	return NewPosition[T](l.list, l.list.head.Previous)
}

func (l *PList[T]) insert(elt T, prev *Node[T], next *Node[T]) *Position[T] {
	n := l.list.Insert(elt, prev, next)
	return NewPosition[T](l.list, n)
}

func (l *PList[T]) InsertToHead(elt T) *Position[T] {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.insert(elt, l.list.head.Previous, l.list.head)
}