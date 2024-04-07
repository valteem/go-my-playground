package arraylist

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/valteem/strdat/utils"
)

func TestAddRemove(t *testing.T) {
	l := New[int](0, 1, 2, 3, 4)
	l.Remove(2)
	expected := New[int](0, 1, 3, 4)
	if !cmp.Equal(l.elements[:l.size], expected.elements[:l.size], cmp.AllowUnexported(List[int]{})) {
		t.Errorf("get %v, expect %v", l, expected)
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		lst      *List[int]
		values   []int
		expected bool
	}{
		{
			lst:      New[int](0, 1, 2, 3, 4),
			values:   []int{1, 2},
			expected: true,
		},
		{
			lst:      New[int](0, 1, 2, 3, 4),
			values:   []int{1, 5},
			expected: false,
		},
	}
	for _, tst := range tests {
		if r := tst.lst.Contains(tst.values...); r != tst.expected {
			t.Errorf("get %t, expect %t", r, tst.expected)
		}
	}
}

func TestValues(t *testing.T) {
	l := New[int](0, 1, 2, 4, 3)
	expected := []int{0, 1, 2, 4, 3}
	if v := l.Values(); !cmp.Equal(v, expected) {
		t.Errorf("get %v, expect %v", v, expected)
	}
}

func TestIndexOf(t *testing.T) {
	tests := []struct {
		lst         *List[int]
		searchValue int
		output      int
	}{
		{
			lst:         New[int](0, 1, 2, 3, 4),
			searchValue: 1,
			output:      1,
		},
		{
			lst:         New[int](0, 1, 2, 3, 4),
			searchValue: 2,
			output:      2,
		},
		{
			lst:         New[int](0, 1, 2, 3, 4),
			searchValue: 5,
			output:      -1,
		},
	}
	for _, tst := range tests {
		if result := tst.lst.IndexOf(tst.searchValue); result != tst.output {
			t.Errorf("get %d, expect %d", result, tst.output)
		}
	}
}

func TestEmpty(t *testing.T) {
	tests := []struct {
		lst    *List[int]
		output bool
	}{
		{
			lst:    New[int](),
			output: true,
		},
		{
			lst:    New[int](0),
			output: false,
		},
	}
	for _, tst := range tests {
		if result := tst.lst.Empty(); result != tst.output {
			t.Errorf("get %t, expect %t", result, tst.output)
		}
	}
}

func TestSize(t *testing.T) {
	tests := []struct {
		lst    *List[int]
		output int
	}{
		{
			lst:    New[int](),
			output: 0,
		},
		{
			lst:    New[int](0),
			output: 1,
		},
		{
			lst:    New[int](0, 1),
			output: 2,
		},
		{
			lst:    New[int](0, 1, 2),
			output: 3,
		},
	}
	for _, tst := range tests {
		if result := tst.lst.Size(); result != tst.output {
			t.Errorf("get %d, expect %d", result, tst.output)
		}
	}
}

func TestClear(t *testing.T) {
	l := New[int](0, 1, 2, 3, 4)
	l.Clear()
	if s := l.Size(); s != 0 {
		t.Errorf("List size: get %d, expect 0", s)
	}
	expectedElements := []int{}
	if !cmp.Equal(l.elements, expectedElements) {
		t.Errorf("List elements: get %v, expect empty slice", l.elements)
	}
}

func TestSort(t *testing.T) {
	comp := func(x, y int) int {
		if x == y {
			return 0
		}
		if x > y {
			return 1
		}
		return -1
	}
	tests := []struct {
		lst              *List[int]
		comp             utils.Comparator[int]
		expectedElements []int
	}{
		{
			lst:              New[int](0, 4, 1, 3, 2),
			comp:             comp,
			expectedElements: []int{0, 1, 2, 3, 4},
		},
		{
			lst:              New[int](0, 4, 1, 4, 1),
			comp:             comp,
			expectedElements: []int{0, 1, 1, 4, 4},
		},
	}
	for _, tst := range tests {
		tst.lst.Sort(tst.comp)
		if result := tst.lst.elements[:tst.lst.size]; !cmp.Equal(result, tst.expectedElements) {
			t.Errorf("get %v, expect %v", result, tst.expectedElements)
		}
	}
}

func TestSwap(t *testing.T) {
	tests := []struct {
		lst              *List[int]
		i                int
		j                int
		expectedElements []int
	}{
		{
			lst:              New[int](0, 1, 2, 3, 4),
			i:                1,
			j:                3,
			expectedElements: []int{0, 3, 2, 1, 4},
		},
		{
			lst:              New[int](0, 1, 2, 3, 4),
			i:                0,
			j:                4,
			expectedElements: []int{4, 1, 2, 3, 0},
		},
		{
			lst:              New[int](0, 1, 2, 3, 4),
			i:                1,
			j:                -1,
			expectedElements: []int{0, 1, 2, 3, 4},
		},
		{
			lst:              New[int](0, 1, 2, 3, 4),
			i:                1,
			j:                5,
			expectedElements: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tst := range tests {
		tst.lst.Swap(tst.i, tst.j)
		if result := tst.lst.elements[:tst.lst.size]; !cmp.Equal(result, tst.expectedElements) {
			t.Errorf("get %v, expect %v", result, tst.expectedElements)
		}
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		description string
		lst         *List[int]
		pos         int
		values      []int
		output      []int
	}{
		{
			description: "insert in the middle",
			lst:         New(0, 1, 2, 3, 4),
			pos:         2,
			values:      []int{5, 6},
			output:      []int{0, 1, 5, 6, 2, 3, 4},
		},
		{
			description: "append",
			lst:         New(0, 1, 2, 3, 4),
			pos:         5,
			values:      []int{5, 6},
			output:      []int{0, 1, 2, 3, 4, 5, 6},
		},
		{
			description: "insert at the beginning",
			lst:         New(0, 1, 2, 3, 4),
			pos:         0,
			values:      []int{5, 6},
			output:      []int{5, 6, 0, 1, 2, 3, 4},
		},
		{
			description: "position out of range - negative",
			lst:         New(0, 1, 2, 3, 4),
			pos:         -1,
			values:      []int{5, 6},
			output:      []int{0, 1, 2, 3, 4},
		},
		{
			description: "position out of range - greater than list size",
			lst:         New(0, 1, 2, 3, 4),
			pos:         6,
			values:      []int{5, 6},
			output:      []int{0, 1, 2, 3, 4},
		},
	}
	for _, tst := range tests {
		tst.lst.Insert(tst.pos, tst.values...)
		if r := tst.lst.elements[:tst.lst.size]; !cmp.Equal(r, tst.output) {
			t.Errorf("%s: get %v, expect %v", tst.description, r, tst.output)
		}
	}
}

func TestSet(t *testing.T) {
	tests := []struct {
		description string
		lst         *List[int]
		pos         int
		value       int
		output      []int
	}{
		{
			description: "set at the beginning",
			lst:         New[int](0, 1, 2, 3, 4),
			pos:         0,
			value:       99,
			output:      []int{99, 1, 2, 3, 4},
		},
		{
			description: "set at the end",
			lst:         New[int](0, 1, 2, 3, 4),
			pos:         4,
			value:       99,
			output:      []int{0, 1, 2, 3, 99},
		},
		{
			description: "append",
			lst:         New[int](0, 1, 2, 3, 4),
			pos:         5,
			value:       99,
			output:      []int{0, 1, 2, 3, 4, 99},
		},
		{
			description: "out of bounds - negative index",
			lst:         New[int](0, 1, 2, 3, 4),
			pos:         -1,
			value:       99,
			output:      []int{0, 1, 2, 3, 4},
		},
		{
			description: "out of bounds - greater than list size",
			lst:         New[int](0, 1, 2, 3, 4),
			pos:         6,
			value:       99,
			output:      []int{0, 1, 2, 3, 4},
		},
	}
	for _, tst := range tests {
		tst.lst.Set(tst.pos, tst.value)
		if r := tst.lst.elements[:tst.lst.size]; !cmp.Equal(r, tst.output) {
			t.Errorf("%s: get %v, expect %v", tst.description, r, tst.output)
		}
	}
}
