package reuse

func AppendToSlice[T any](s []T, a T) {
	s = append(s, a) // length and capacity of 's' are not affected by 'append' after function is completed
}

func AppendToSliceP[T any](s *[]T, a T) {
	*s = append(*s, a) // (*s) is a slice, since s is a pointer (?)
}

type SliceContainer [T any] struct {
	s []T
}

func NewSliceContainer[T any](s []T) SliceContainer[T] {
	return SliceContainer[T]{s: s}
}

func (sc SliceContainer[T]) Slice() []T {
	return sc.s
}

func (sc SliceContainer[T])AppendToSlice(a T) {
	sc.s = append(sc.s, a)
}

func (sc *SliceContainer[T])AppendToSliceP(a T) {
	sc.s = append(sc.s, a)
}

// utility func for testing
func SlicesEqual[T comparable](s1 []T, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i <  len(s1); i++ {
		if s1[i] != s2[i] {return false}
	}
	return true
}