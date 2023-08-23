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

func (s Set[T]) Insert(elt T) { // maps are actually pointers (https://stackoverflow.com/a/53680008), hence no need for (s *Set)
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

func (s Set[T]) Remove(elt T) {
	delete(s, elt)
}

func (s Set[T]) RemoveMany(elts... T) {
	for _, v := range elts {
		s.Remove(v)
	}
}

func (s Set[T]) Merge(s2 Set[T]) {
	for k := range s2 {
		s[k] = struct{}{}
	}
}

func (s Set[T]) ExtractSlice() []T {
	l := []T{}
	for k := range s {
		l = append(l, k)
	}
	return l
}

func (s Set[T]) Contains(elt T) bool  {
	_, in := s[elt]
	return in
}

func (s Set[T]) IsSuperSetOf(s2 Set[T]) bool {

	if s2 == nil {
		return true
	}

	if len(s2) > len(s) {
		return false
	}

	for k := range s2 {
		if !s.Contains(k) {
			return false
		}
	}

	return true

}