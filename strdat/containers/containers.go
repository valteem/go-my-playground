package containers

import (
	"cmp"
	"slices"

	"github.com/valteem/strdat/utils"
)

type Container[T any] interface {
	Empty() bool
	Size() int
	Clear()
	Values() []T
	String() string
}

func GetSortedValues[T cmp.Ordered](c Container[T]) []T {
	values := c.Values()
	if len(values) < 2 {
		return values
	}
	slices.Sort(values)
	return values
}

func GetSortedValuesFunc[T any](c Container[T], comp utils.Comparator[T]) []T {
	values := c.Values()
	if len(values) < 2 {
		return values
	}
	slices.SortFunc(values, comp)
	return values
}
