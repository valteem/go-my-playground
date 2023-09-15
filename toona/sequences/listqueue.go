package sequences

type ListQueue [T any] struct {
	list *DoubleLinkedList[T]
}

func NewListQueue[T any]() *ListQueue[T] {
	return &ListQueue[T]{list: NewDoubleLinkedList[T]()}
}

func (q *ListQueue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *ListQueue[T]) Enqueue(e T) {
	q.list.AddToHead(e)
}

func (q *ListQueue[T]) Dequeue() (T, error) {
	r, e := q.list.RemoveFromTail()
	if e != nil {
		e = errQueueEmpty
	}
	return r, e
}