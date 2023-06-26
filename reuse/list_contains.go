package reuse

type List [T any] []T

func (l *List[T]) Contains(t T, cmp func(T, T) bool) bool {
	for _, s := range *l {
		if cmp(s, t) {
			return true
		}
	}
	return false
}