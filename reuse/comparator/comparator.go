package comparator

func NonZero[T comparable](v, alt T) T {
	var zero T
	if v != zero {
		return v
	}
	return alt
}
