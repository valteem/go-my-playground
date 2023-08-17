package reuse

func AppendToSlice[T any](s []T, e T) {
	s = append(s, e)
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

func (sc SliceContainer[T])AppendToSlice(e T) {
	sc.s = append(sc.s, e)
}

func (sc *SliceContainer[T])AppendToSliceP(e T) {
	sc.s = append(sc.s, e)
}