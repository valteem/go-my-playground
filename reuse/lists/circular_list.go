package lists

type ListElement[T any] struct {
	payload T
	next    *ListElement[T]
	prev    *ListElement[T]
}

func newListElement[T any](payload T) *ListElement[T] {
	return &ListElement[T]{payload: payload}
}

type CircularList[T any] struct {
	anchor *ListElement[T]
}

func NewCircularList[T any]() *CircularList[T] {
	var t T
	elt := newListElement(t)
	elt.next = elt
	elt.prev = elt
	return &CircularList[T]{anchor: elt}
}

func (cl *CircularList[T]) Add(t T) {

	elt := newListElement(t)

	elt.prev = cl.anchor
	elt.next = cl.anchor.next
	cl.anchor.next.prev = elt
	cl.anchor.next = elt

}

func (cl *CircularList[T]) Last() *ListElement[T] {
	return cl.anchor.next
}
