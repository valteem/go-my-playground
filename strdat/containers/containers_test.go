package containers

import (
	"cmp"
	"fmt"
	"strings"
	"testing"
)

type ContainerTest[T any] struct {
	values []T
}

func (c ContainerTest[T]) Empty() bool {
	return len(c.values) == 0
}

func (c ContainerTest[T]) Size() int {
	return len(c.values)
}

func (c *ContainerTest[T]) Clear() {
	c.values = []T{}
}

func (c ContainerTest[T]) Values() []T {
	return c.values
}

func (c ContainerTest[T]) String() string {
	output := "ContainerTest\n"
	var values []string
	for _, v := range c.values {
		values = append(values, fmt.Sprintf("%v", v))
	}
	output += strings.Join(values, ", ")
	output = strings.TrimSuffix(output, " ")
	return output
}

func TestGetSortedValuesInt(t *testing.T) {
	c := &ContainerTest[int]{} // https://stackoverflow.com/a/33937234
	c.values = []int{4, 2, 1, 5, 3}
	v := GetSortedValues[int](c)
	for i := 0; i < c.Size()-1; i++ {
		if v[i] > v[i+1] {
			t.Errorf("sorting order violation: v[%d] = %d > v[%d] = %d", i, v[i], i+1, v[i+1])
		}
	}
}

type WrapInt struct {
	i int
}

func TestGetSortedValuesWrapInt(t *testing.T) {
	c := &ContainerTest[WrapInt]{}
	c.values = []WrapInt{{4}, {2}, {1}, {5}, {3}}
	v := GetSortedValuesFunc[WrapInt](c, func(a, b WrapInt) int {
		return cmp.Compare[int](a.i, b.i)
	})
	for i := 0; i < c.Size()-1; i++ {
		if v[i].i > v[i+1].i {
			t.Errorf("sorting order violation: v[%d] = %d > v[%d] = %d", i, v[i].i, i+1, v[i+1].i)
		}
	}
}
