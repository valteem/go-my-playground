package containers

type EnumerableWithIndex[T any] interface {
	// Calls argument function once for each element
	Each(func(index int, value T))
	// Passes each element to argument function
	// Rseturns `true` if function ever returns `true` for any element
	Any(func(index int, value T) bool) bool
	// Passes each element to argument function
	// Returns `true` if argument function returns `true` for all elements
	All(func(index int, value T) bool) bool
	// Passes each element to argument function
	// Returns first (index, value) pair for which argument function is true
	// or (-1, nil) if argument function returns `false` for all elements
	Find(func(index int, value T) bool) (int, T)
}

type EnumerableWithKey[K, V any] interface {
	// Calls argument function once for each element
	Each(func(key K, value V))
	// Passes each element to argument function
	// Rseturns `true` if function ever returns `true` for any element
	Any(func(key K, value V) bool) bool
	// Passes each element to argument function
	// Returns `true` if argument function returns `true` for all elements
	All(func(key K, value V) bool) bool
	// Passes each element to argument function
	// Returns first (key, value) pair for which argument function is true
	// or (nil, nil) if argument function returns `false` for all elements
	Find(func(key K, value V) bool) (K, V)
}
