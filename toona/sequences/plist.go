// positional list

package sequences

import (
	"fmt"
	"sync"
)

var (
	ErrWrongCont = fmt.Errorf("position does not belong to the list")
	ErrOutOfBoundOrRemoved = fmt.Errorf(("position out of list boundaries or removed from list"))
	ErrPListEmpty = fmt.Errorf("positional list empty")
	ErrPListEmptyInsert = fmt.Errorf("positional list empty, only InsertToTail()/InsertToHead operations available")
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
// lock/unlock is take care of by the caller
	// l.lock.Lock()
	// defer l.lock.Unlock()
	if pos.cont != l.list {
		return nil, ErrWrongCont
	}
	if pos.node.Next == nil {
		return nil, ErrOutOfBoundOrRemoved
	}
	return pos.node, nil
}

func (l *PList[T]) position(node *Node[T]) *Position[T] {
// lock/unlock is take care of by the caller
	// l.lock.Lock()
	// defer l.lock.Unlock()
	if node == l.list.tail || node == l.list.head {
		return nil
	}
	return NewPosition[T](l.list, node)
}

func (l *PList[T]) Head() (*Position[T], error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.list.size == 0 {
		var p *Position[T]
		return p, ErrPListEmpty
	}
	return NewPosition[T](l.list, l.list.head.Previous), nil
}

func (l *PList[T]) Tail() (*Position[T], error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.list.size == 0 {
		var p *Position[T]
		return p, ErrPListEmpty
	}
	return NewPosition[T](l.list, l.list.tail.Next), nil
}

func (l *PList[T]) Before(pos *Position[T]) (*Position[T], error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.list.size == 0 {
		var p *Position[T]
		return p, ErrPListEmpty
	}
	n, e := l.validate(pos)
	if e != nil {
		var p *Position[T]
		return p, e
	}
	return l.position(n.Previous), nil
}

func (l *PList[T]) After(pos *Position[T]) (*Position[T], error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.list.size == 0 {
		var p *Position[T]
		return p, ErrPListEmpty
	}
	n, e := l.validate(pos)
	if e != nil {
		var p *Position[T]
		return p, e
	}
	return l.position(n.Next), nil
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

func (l *PList[T]) InsertToTail(elt T) *Position[T] {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.insert(elt, l.list.tail, l.list.tail.Next)
}

func (l *PList[T]) InsertBefore(elt T, pos *Position[T]) (*Position[T], error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.list.size == 0 {
		var p *Position[T]
		return p, ErrPListEmptyInsert
	}
	node, err := l.validate(pos)
	if err != nil {
		var p *Position[T]
		return p, err
	}
	return l.insert(elt, node.Previous, node), nil
}

func (l *PList[T]) InsertAfter(elt T, pos *Position[T]) (*Position[T], error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.list.size == 0 {
		var p *Position[T]
		return p, ErrPListEmptyInsert
	}
	node, err := l.validate(pos)
	if err != nil {
		var p *Position[T]
		return p, err
	}
	return l.insert(elt, node, node.Next), nil
}

func (l *PList[T]) Remove(pos *Position[T]) (T, error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	node, err := l.validate(pos)
	if err != nil {
		var t T
		return t, err
	}
	return l.list.Remove(node), nil
}

func (l *PList[T]) Replace(elt T, pos *Position[T]) (T, error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	node, err := l.validate(pos)
	if err != nil {
		var t T
		return t, err
	}
	r := node.Element
	node.Element = elt
	return r, nil
}