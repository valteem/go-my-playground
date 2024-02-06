package reuse

type CanCompare struct {
	x int
	s string
}

func NewCanCompare(x int, s string) CanCompare {
	return CanCompare{x: x, s: s}
}

type DoNotCompare struct {
	_ [0]func()
	x int
	s string
}

func NewDoNotCompare(x int, s string) DoNotCompare {
	return DoNotCompare{x: x, s: s}
}