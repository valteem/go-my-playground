package bag

import (
	"fmt"
)

var (
	err = fmt.Errorf("nothing to remove from bag")
)

type Bag[T comparable] struct {
	size int
	contents map[T]int // how many items of a certain T value
}

func New[T comparable]()  *Bag[T] {
	return &Bag[T]{0,make(map[T]int) }
}

func (b *Bag[T]) Insert(elt T) {
	b.contents[elt]++ // magic
	b.size++
}

func (b *Bag[T]) Remove(elt T) error {
	if count, ok := b.contents[elt]; ok {
		if count > 1 {
			b.contents[elt]--
		} else {
			delete(b.contents, elt)
		}
		b.size--
		return nil
	} else {
		return err
	}
}

func (b *Bag[T]) Count(elt T) int {
	return b.contents[elt]
}

func (b *Bag[T]) Reset() {
	*b = *New[T]() // magic again
}