package arraylist

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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
