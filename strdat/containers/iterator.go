package containers

type IteratorWithIndex[T any] interface {
	// Moves to the next element
	// Returns `true` if there exists next element in container, otherwise `false`
	// Points to the first element when called for the first time (if the container has any elements at all)
	// Modifies the state of the iterator
	Next()

	// Returns value of the current element
	// Doesn't modify the state of the iterator
	Value() T

	// Returns index of the current element
	// Doesn't modify the state of the iterator
	Index() int

	// Resets the iterator to its initial state (before first call)
	Begin()

	// Moves the iterator to the first element (if the container has any elements at all)
	// Returns `true` if there exists the first element, otherwise `false`
	First() bool

	// Moves the iterator from current position to the first element for which argument function returns true
	// Returns `true` if such element is found, otherwise `false`
	NextTo(func(index int, value T) bool) bool
}

type IteratorWithKey[K, V any] interface {
	// Moves to the next element
	// Returns `true` if there exists next element in container, otherwise `false`
	// Points to the first element when called for the first time (if the container has any elements at all)
	// Modifies the state of the iterator
	Next()
	// Good question: Which element is `first` in key/value iterator?

	// Returns the value of the current element
	// Doesn't modify the state of the iterator
	Value() V

	// Returns the key of the current element
	// Doesn't modify the state of the iterator
	Key() K

	// Resets the iterator to its initial state (before first call)
	Begin()

	// Moves the iterator to the first element (if the container has any elements at all)
	// Returns `true` if there exists the first element, otherwise `false`
	First() bool

	// Moves the iterator from current position to the first element for which argument function returns true
	// Returns `true` if such element is found, otherwise `false`
	NextTo(func(key K, value V) bool) bool
}

type ReverseIteratorWithIndex[T any] interface {
	// Moves the iterator to the previous element
	// Returns `true` if there exists a previous element, otherwise `false`
	// Modifies the state of the iterator
	Prev() bool

	// Moves the iterator past the last element, calling Prev() after fetches the last element
	End()

	// Moves the iterator to the last element
	// Returns `true` if there exists a last element, otherwise `false`
	// Modifies the state of the iterator
	Last() bool

	// Moves back the iterator from current to the first preceding element for which the argument function returns `true`
	// Returns `true` if there exists such an element, otherwise `false`
	// Modifies the state of the iterator
	PrevTo(func(index int, value T) bool) bool

	IteratorWithIndex[T]
}

type ReverseIteratorWithKey[K, V any] interface {
	// Moves the iterator to the previous element
	// Returns `true` if there exists a previous element, otherwise `false`
	// Modifies the state of the iterator
	Prev() bool

	// Moves the iterator past the last element, calling Prev() after fetches the last element
	End()

	// Moves the iterator to the last element
	// Returns `true` if there exists a last element, otherwise `false`
	// Modifies the state of the iterator
	Last() bool

	// Moves back the iterator from current to the first preceding element for which the argument function returns `true`
	// Returns `true` if there exists such an element, otherwise `false`
	// Modifies the state of the iterator
	PrevTo(func(key K, value V) bool) bool

	IteratorWithKey[K, V]
}
