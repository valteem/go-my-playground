package nestedstruct

type A struct{}

func (A) Write() string {
	return "A"
}

type B struct{}

func (B) Write() string {
	return "B"
}

type C1 struct {
	A   // A.Write() promoted
	b B // B.Write() shadowed
}

type C2 struct {
	a A // A.Write() shadowed
	B   // B.Write() promoted
}

type C1B struct {
	A
	b B
}

// overload A.Write() promotion
func (c C1B) Write() string {
	return c.b.Write()
}

type C2A struct {
	a A
	B
}

// overload B.Write() promotion
func (c C2A) Write() string {
	return c.a.Write()
}
