package bufiom

import("fmt")

type SomeSet interface {
	Show() bool
}

type SetOfInt struct {
	Item1 int
	Item2 int
}

func (s SetOfInt) Show() bool {
	fmt.Println("item1 = ", s.Item1, "item2 = ", s.Item2)
	return true
}

type SetOfStr struct {
	Item1 string
	Item2 string
}

func (s SetOfStr) Show() bool {
	fmt.Println("item1 = ", s.Item1, "item2 = ", s.Item2)
	return true
}

func AssertSetType(s SomeSet) {
	fmt.Println("Passed as value: ", s)
	bip, okip := s.(*SetOfInt)
	biv, okiv := s.(SetOfInt)
	bsp, oksp := s.(*SetOfStr)
	bsv, oksv := s.(SetOfStr)
	fmt.Println ("Type assertion - *SetOfInt", bip, okip)
	fmt.Println ("Type assertion - SetOfInt", biv, okiv)
	fmt.Println ("Type assertion - *SetOfStr", bsp, oksp)
	fmt.Println ("Type assertion - SetOfStr", bsv, oksv)
}

type SomeSlice interface {
	Show() bool
}

type SliceOfInt struct {
	S []int
}

func (s SliceOfInt) Show() bool {
	fmt.Println(s.S)
	return true
}

type SliceOfStr struct {
	S []string
}

func (s SliceOfStr) Show() bool {
	fmt.Println(s.S)
	return true
}

func AssertSliceType(s SomeSlice) {
	fmt.Println("Passed as value: ", s)
	bip, okip := s.(*SliceOfInt)
	biv, okiv := s.(SliceOfInt)
	bsp, oksp := s.(*SliceOfStr)
	bsv, oksv := s.(SliceOfStr)
	fmt.Println ("Type assertion - *SliceOfInt", bip, okip)
	fmt.Println ("Type assertion - SliceOfInt", biv, okiv)
	fmt.Println ("Type assertion - *SliceOfStr", bsp, oksp)
	fmt.Println ("Type assertion - SliceOfStr", bsv, oksv)
}