package set

type Set[T comparable] map[T]struct{}

func NewWithLen[T comparable](l int) Set[T] {
	return make(Set[T], l)
}

 func New[T comparable](elts ...T) Set[T] {
	s := NewWithLen[T](len(elts))
	s.InsertMany(elts...)
	return s
 }

func (s Set[T]) Insert(elt T) { // maps are actually pointers (https://stackoverflow.com/a/53680008), hence no neet for (s *Set)
	s[elt] = struct{}{}
}

func (s Set[T]) InsertMany(elts ...T) {
	for _, v := range elts {
		s.Insert(v)
	}
}

func (s Set[T]) InsertNew(elt T) Set[T] {
	s[elt] = struct{}{}
	return s
}