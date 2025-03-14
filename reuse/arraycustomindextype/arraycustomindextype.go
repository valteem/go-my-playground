package arraycustomindextype

// Indexing array using custom type
type ArrayIndex int

const (
	ArrayIndexZero ArrayIndex = iota
	ArrayIndexOne
	ArrayIndexTwo
	ArrayIndexThree
	ArrayIndexFour
)

var squares = [...]int{
	ArrayIndexZero:  0,
	ArrayIndexOne:   1,
	ArrayIndexTwo:   4,
	ArrayIndexThree: 9,
	ArrayIndexFour:  16,
}
