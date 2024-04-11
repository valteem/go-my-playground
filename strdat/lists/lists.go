package lists

import (
	"github.com/valteem/strdat/containers"
	"github.com/valteem/strdat/utils"
)

type List[T comparable] interface {
	Get(index int) (T, bool)
	Remove(index int) // does not return error or bool
	Add(values ...T)
	Contains(values ...T) bool // intersection (?)
	Sort(comparator utils.Comparator[T])
	Swap(index1, index2 int)
	Insert(index int, values ...T) // after or before index?
	Set(index int, value T)

	containers.Container[T]
	// Empty()  bool
	// Size     int
	// Clear()
	// Values() []T
	// String   string
}
