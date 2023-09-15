package sequences

type ListDeque [T any] struct {
	list *DoubleLinkedList[T]
}

func NewListDeque[T any]() *ListDeque[T] {
	return &ListDeque[T]{list: NewDoubleLinkedList[T]()}
}

func (q *ListDeque[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *ListDeque[T]) AddFirst(e T) {
	q.list.AddToHead(e)
}

func (q *ListDeque[T]) AddLast(e T) {
	q.list.AddToTail(e)
}

func (q *ListDeque[T]) DeleteFirst() (T, error) {
	r, e := q.list.RemoveFromHead()
	if e != nil {
		e = errDequeEmpty
	}
	return r, e
}

func (q *ListDeque[T]) DeleteLast() (T, error) {
	r, e := q.list.RemoveFromTail()
	if e != nil {
		e = errDequeEmpty
	}
	return r, e
}