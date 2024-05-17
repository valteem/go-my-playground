package doublylinkedlist

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRemoveFirstLast(t *testing.T) {
	l := New[string]("a", "b", "c", "x")
	l.Remove(0)
	if l.first.prev != nil {
		t.Errorf("New list.first.prev should be nil")
	}
	l.Remove(2)
	if l.last.next != nil {
		t.Errorf("New list.last.next should be nil")
	}
	l.Remove(5)
	if l.size != 2 {
		t.Errorf("Trying to remove from index-out-of-bounds should have no effect")
	}
	l.Remove(0)
	if l.first.value != "c" || l.last.value != "c" || l.size != 1 {
		t.Errorf("List structure after three Remove() ops differs from expected")
	}
}

func TestListNew(t *testing.T) {
	tests := []struct {
		input         []int
		expectedEmpty bool
		expectedSize  int
	}{
		{
			input:         []int{},
			expectedEmpty: true,
			expectedSize:  0,
		},
		{
			input:         []int{1},
			expectedEmpty: false,
			expectedSize:  1,
		},
		{
			input:         []int{1, 2},
			expectedEmpty: false,
			expectedSize:  2,
		},
		{
			input:         []int{1, 2, 3},
			expectedEmpty: false,
			expectedSize:  3,
		},
	}
	for _, tst := range tests {
		l := New[int](tst.input...)

		if empty := l.Empty(); empty != tst.expectedEmpty {
			t.Errorf("%v - Empty(): get %t, expect %t", tst.input, empty, tst.expectedEmpty)
		}
		if size := l.Size(); size != tst.expectedSize {
			t.Errorf("%v - Size(): get %d, expect %d", tst.input, size, tst.expectedSize)
		}
		if v := l.Values(); !cmp.Equal(v, tst.input) {
			t.Errorf("Values(): get %v, expect %v", v, tst.input)
		}
	}
}

func TestListAdd(t *testing.T) {
	tests := []struct {
		input []string
		empty bool
		size  int
	}{
		{
			input: []string{},
			empty: true,
			size:  0,
		},
		{
			input: []string{"a"},
			empty: false,
			size:  1,
		},
		{
			input: []string{"a", "b"},
			empty: false,
			size:  2,
		},
		{
			input: []string{"a", "b", "c"},
			empty: false,
			size:  3,
		},
	}
	for _, tst := range tests {
		l := New[string]()
		l.Add(tst.input...)
		if empty := l.Empty(); empty != tst.empty {
			t.Errorf("%v - Empty(): get %t, expect %t", tst.input, empty, tst.empty)
		}
		if size := l.Size(); size != tst.size {
			t.Errorf("%v - Size(): get %d, expect %d", tst.input, size, tst.size)
		}
		if values := l.Values(); !cmp.Equal(values, tst.input) {
			t.Errorf("Values(): get %v, expect %v", values, tst.input)
		}
	}
}

func TestListAppendPrepend(t *testing.T) {
	tests := []struct {
		init      []string
		toPrepend []string
		toAppend  []string
	}{
		{
			init:      []string{},
			toPrepend: []string{"a", "b", "c"},
			toAppend:  []string{},
		},
		{
			init:      []string{"a", "b", "c"},
			toPrepend: []string{},
			toAppend:  []string{},
		},
		{
			init:      []string{},
			toPrepend: []string{},
			toAppend:  []string{"a", "b", "c"},
		},
		{
			init:      []string{"k", "l", "m"},
			toPrepend: []string{"a", "b", "c"},
			toAppend:  []string{},
		},
		{
			init:      []string{"a", "b", "c"},
			toPrepend: []string{},
			toAppend:  []string{"k", "l", "m"},
		},
		{
			init:      []string{"k", "l", "m"},
			toPrepend: []string{"a", "b", "c"},
			toAppend:  []string{"u", "v", "w"},
		},
	}
	for _, tst := range tests {
		l := New[string](tst.init...)
		l.Prepend(tst.toPrepend...)
		l.Append(tst.toAppend...)
		expected := append(append(append([]string{}, tst.toPrepend...), tst.init...), tst.toAppend...)
		values := l.Values()
		if !cmp.Equal(values, expected) {
			t.Errorf("Prepend, Append: get %v, expect %v", values, expected)
		}
	}
}

func TestListRemove(t *testing.T) {
	tests := []struct {
		input  []string
		idx    []int
		output []string
	}{
		{
			input:  []string{"a", "b", "c", "u", "v", "w"},
			idx:    []int{1, 3, 2, 5},
			output: []string{"a", "c", "w"},
		},
		{
			input:  []string{"a", "b", "c", "u", "v", "w"},
			idx:    []int{0, 4, 3, 1, 5},
			output: []string{"b", "u"},
		},
	}
	for _, tst := range tests {
		l := New[string](tst.input...)
		for _, i := range tst.idx {
			l.Remove(i)
		}
		if values := l.Values(); !cmp.Equal(values, tst.output) {
			t.Errorf("Remove(): get %v, expect%v", values, tst.output)
		}
	}
}

func TestListGet(t *testing.T) {
	tests := []struct {
		input   []string
		indices []int
		output  []string
	}{
		{
			input:   []string{"a", "b", "c", "k", "l", "m", "n", "u", "v", "w"},
			indices: []int{1, 3, 5, 7, 11},
			output:  []string{"b", "k", "m", "u", ""},
		},
		{
			input:   []string{"a", "b", "c", "k", "l", "m", "n", "u", "v", "w"},
			indices: []int{-1, 11, 3, 5, 7},
			output:  []string{"", "", "k", "m", "u"},
		},
	}
	for _, tst := range tests {
		l := New[string](tst.input...)
		var output []string
		for _, i := range tst.indices {
			v, _ := l.Get(i)
			output = append(output, v)
		}
		if !cmp.Equal(output, tst.output) {
			t.Errorf("Get(): get %v, expect %v", output, tst.output)
		}
	}
}

func TestListSwap(t *testing.T) {
	tests := []struct {
		input  []string
		idx1   []int
		idx2   []int
		output []string
	}{
		{
			input:  []string{"a", "b", "c", "k", "l", "m", "n"},
			idx1:   []int{0, 1, 8, 4, 1, 5, 6},
			idx2:   []int{1, 0, 7, 8, 3, 4, 1},
			output: []string{"a", "n", "c", "b", "m", "l", "k"},
		},
		{
			input:  []string{"a", "b", "c"},
			idx1:   []int{0, 0, 2, 2, 1, 4},
			idx2:   []int{2, 1, 0, 1, 4, 1},
			output: []string{"a", "b", "c"},
		},
	}
	for _, tst := range tests {
		l := New[string](tst.input...)
		for ri, i := range tst.idx1 {
			j := tst.idx2[ri]
			l.Swap(i, j)
		}
		if values := l.Values(); !cmp.Equal(values, tst.output) {
			t.Errorf("Swap(): get %v, expect %v", values, tst.output)
		}
	}
}

func TestListSort(t *testing.T) {
	tests := []struct {
		input  []string
		output []string
	}{
		{
			input:  []string{"b", "a", "x", "w"},
			output: []string{"a", "b", "w", "x"},
		},
		{
			input:  []string{"b", "a", "11", "1"},
			output: []string{"1", "11", "a", "b"},
		},
	}
	for _, tst := range tests {
		l := New[string](tst.input...)
		l.Sort(func(a, b string) int {
			if a > b {
				return 1
			} else if a < b {
				return -1
			} else {
				return 0
			}
		})
		if values := l.Values(); !cmp.Equal(values, tst.output) {
			t.Errorf("Sort(): get %v, expect %v", values, tst.output)
		}
	}
}

func TestListClear(t *testing.T) {
	tests := []struct {
		input []int
	}{
		{
			input: []int{},
		},
		{
			input: []int{1, 2, 4},
		},
	}
	for _, tst := range tests {
		l := New[int](tst.input...)
		l.Clear()
		if !l.Empty() {
			t.Errorf("Clear(): failed to clear %v", tst.input)
		}
	}
}

func TestListContains(t *testing.T) {
	tests := []struct {
		input  []string
		lookup []string
		output bool
	}{
		{
			input:  []string{"a", "b", "c"},
			lookup: []string{"c"},
			output: true,
		},
		{
			input:  []string{"a", "b", "c"},
			lookup: []string{"c", "x"},
			output: false,
		},
		{
			input:  []string{"a", "b", "c"},
			lookup: []string{},
			output: true,
		},
	}
	for _, tst := range tests {
		l := New[string](tst.input...)
		if output := l.Contains(tst.lookup...); output != tst.output {
			t.Errorf("Contains(): input %v, lookup %v, get %t, expect %t", tst.input, tst.lookup, output, tst.output)
		}
	}
}

func TestListInsert(t *testing.T) {
	tests := []struct {
		input  []string
		insert []string
		pos    int
		output []string
	}{
		{
			input:  []string{"a", "b", "c"},
			insert: []string{"u", "v"},
			pos:    1,
			output: []string{"a", "u", "v", "b", "c"},
		},
		{
			input:  []string{"a", "b", "c"},
			insert: []string{"u", "v"},
			pos:    3,
			output: []string{"a", "b", "c", "u", "v"},
		},
		{
			input:  []string{"a", "b", "c"},
			insert: []string{"u", "v"},
			pos:    0,
			output: []string{"u", "v", "a", "b", "c"},
		},
		{
			input:  []string{"a", "b", "c"},
			insert: []string{"u", "v"},
			pos:    4,
			output: []string{"a", "b", "c"},
		},
		{
			input:  []string{"a", "b", "c"},
			insert: []string{},
			pos:    1,
			output: []string{"a", "b", "c"},
		},
	}
	for _, tst := range tests {
		l := New[string](tst.input...)
		l.Insert(tst.pos, tst.insert...)
		if output := l.Values(); !cmp.Equal(output, tst.output) {
			t.Errorf("Insert(): get %v, expect %v", output, tst.output)
		}
	}
}

func TestListSet(t *testing.T) {
	tests := []struct {
		input     []string
		indices   []int
		setValues []string
		output    []string
	}{
		{
			input:     []string{"a", "b", "c"},
			indices:   []int{0, 1, 0, 3, 2, 1, 5},
			setValues: []string{"u", "v", "w", "f", "g", "h", "x"},
			output:    []string{"w", "h", "g", "f"},
		},
		{
			input:     []string{"a", "b", "c"},
			indices:   []int{0, 1, 2, 5, 6, 7},
			setValues: []string{"u", "v", "w", "", "", ""},
			output:    []string{"u", "v", "w"},
		},
	}
	for _, tst := range tests {
		l := New[string](tst.input...)
		for i, idx := range tst.indices {
			l.Set(idx, tst.setValues[i])
		}
		if output := l.Values(); !cmp.Equal(output, tst.output) {
			t.Errorf("Set(): get %v, expect %v", output, tst.output)
		}
	}
}

func TestListEach(t *testing.T) {
	tests := []struct {
		input []int
		f     func(i, v int)
	}{
		{
			input: []int{0, 1, 2, 3, 4},
			f: func(i, v int) {
				if i != v {
					t.Errorf("Each(): get %d != %d", i, v)
				}
			},
		},
		{
			input: []int{1, 2, 3, 4, 5},
			f: func(i, v int) {
				if (i + 1) != v {
					t.Errorf("Each(): get %d != %d", (i + 1), v)
				}
			},
		},
		{
			input: []int{1, 4, 9, 16, 25},
			f: func(i, v int) {
				if a := (i + 1) * (i + 1); a != v {
					t.Errorf("Each(): get %d != %d", a, v)
				}
			},
		},
	}
	for _, tst := range tests {
		l := New[int](tst.input...)
		l.Each(tst.f)
	}
}

func TestListMap(t *testing.T) {
	tests := []struct {
		input  []int
		f      func(i, v int) int
		output []int
	}{
		{
			input:  []int{1, 2, 3, 4, 5},
			f:      func(i, v int) int { return i * i },
			output: []int{0, 1, 4, 9, 16},
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			f:      func(i, v int) int { return v * v },
			output: []int{1, 4, 9, 16, 25},
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			f:      func(i, v int) int { return 2 * i },
			output: []int{0, 2, 4, 6, 8},
		}, {
			input:  []int{1, 2, 3, 4, 5},
			f:      func(i, v int) int { return 2 * v },
			output: []int{2, 4, 6, 8, 10},
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			f:      func(i, v int) int { return i * v },
			output: []int{0, 2, 6, 12, 20},
		},
	}
	for _, tst := range tests {
		l := New[int](tst.input...)
		m := l.Map(tst.f)
		if output := m.Values(); !cmp.Equal(output, tst.output) {
			t.Errorf("Map(): get %v, expect %v", output, tst.output)
		}
	}
}

func TestListSelect(t *testing.T) {
	tests := []struct {
		input  []int
		f      func(i, v int) bool
		output []int
	}{
		{
			input: []int{1, 2, 3, 4, 5, 6, 7, 8},
			f: func(i, v int) bool {
				return (v/2)*2 == v
			},
			output: []int{2, 4, 6, 8},
		},
		{
			input: []int{1, 2, 3, 4, 5, 6, 7, 8},
			f: func(i, v int) bool {
				return (v/2)*2 != v
			},
			output: []int{1, 3, 5, 7},
		},
		{
			input: []int{1, 2, 4, 5, 16, 17, 36, 37},
			f: func(i, v int) bool {
				return i*i == v
			},
			output: []int{4, 16, 36},
		},
	}
	for _, tst := range tests {
		l := New[int](tst.input...)
		s := l.Select(tst.f)
		if output := s.Values(); !cmp.Equal(output, tst.output) {
			t.Errorf("Select(): get %v, expect %v", output, tst.output)
		}
	}
}

func TestListAny(t *testing.T) {
	tests := []struct {
		input  []int
		f      func(i, v int) bool
		output bool
	}{
		{
			input: []int{0, 1, 4, 5, 7},
			f: func(i, v int) bool {
				return i*i == v
			},
			output: true,
		},
		{
			input: []int{1, 2, 3, 4, 5},
			f: func(i, v int) bool {
				return i*i == v
			},
			output: false,
		},
	}
	for _, tst := range tests {
		l := New[int](tst.input...)
		if output := l.Any(tst.f); output != tst.output {
			t.Errorf("Any() on %v: get %t, expect %t", tst.input, output, tst.output)
		}
	}
}

func TestListAll(t *testing.T) {
	tests := []struct {
		input  []int
		f      func(i, v int) bool
		output bool
	}{
		{
			input: []int{0, 1, 4, 9, 16},
			f: func(i, v int) bool {
				return i*i == v
			},
			output: true,
		},
		{
			input: []int{0, 1, 4, 9, 17},
			f: func(i, v int) bool {
				return i*i == v
			},
			output: false,
		},
	}
	for _, tst := range tests {
		l := New[int](tst.input...)
		if output := l.All(tst.f); output != tst.output {
			t.Errorf("Any() on %v: get %t, expect %t", tst.input, output, tst.output)
		}
	}
}

func TestListFind(t *testing.T) {
	tests := []struct {
		input       []int
		f           func(i, v int) bool
		outputIndex int
		outputValue int
	}{
		{
			input: []int{1, 2, 4, 9, 16},
			f: func(i, v int) bool {
				return i*i == v
			},
			outputIndex: 2,
			outputValue: 4,
		},
		{
			input: []int{1, 2, 3, 4, 5},
			f: func(i, v int) bool {
				return i*i == v
			},
			outputIndex: -1,
			outputValue: 0,
		},
	}
	for _, tst := range tests {
		l := New[int](tst.input...)
		outputIndex, outputValue := l.Find(tst.f)
		if outputIndex != tst.outputIndex || outputValue != tst.outputValue {
			t.Errorf("Find() : get (%d, %d), expect (%d, %d)", outputIndex, outputValue, tst.outputIndex, tst.outputValue)
		}
	}
}

func TestListIteratorNextonEmpty(t *testing.T) {
	l := New[string]()
	it := l.Iterator()
	for it.Next() {
		t.Errorf("Should not iterate over empty list")
	}
}

func TestListIteratorNext(t *testing.T) {
	tests := []struct {
		input   []string
		indices []int
		values  []string
	}{
		{
			input: []string{"a", "b", "c", "u", "v", "w"},
		},
	}
	for _, tst := range tests {
		l := New[string](tst.input...)
		it := l.Iterator()
		ri := 0
		for it.Next() {
			if i, v := it.Index(), it.Value(); i != ri || v != tst.input[ri] {
				t.Errorf("Next(): get (%d, %s), expect (%d, %s)", i, v, ri, tst.values[ri])
			}
			ri++
		}
	}
}

func TestListIteratorPrevOnEmpty(t *testing.T) {
	l := New[string]()
	it := l.Iterator()
	for it.Prev() {
		t.Errorf("should not iterate over empty list")
	}
}

func TestListIteratorPrev(t *testing.T) {
	tests := []struct {
		input   []string
		indices []int
		values  []string
	}{
		{
			input: []string{"a", "b", "c", "u", "v", "w"},
		},
	}
	for _, tst := range tests {
		l := New[string](tst.input...)
		it := l.Iterator()
		it.End()
		ri := 5
		for it.Prev() {
			if i, v := it.Index(), it.Value(); i != ri || v != tst.input[ri] {
				t.Errorf("Next(): get (%d, %s), expect (%d, %s)", i, v, ri, tst.values[ri])
			}
			ri--
		}
	}
}

func TestListIteratorBeginEndFirstLast(t *testing.T) {
	l := New[string]("apples", "pears", "cucumbers", "onions", "potatoes")
	it := l.Iterator()
	it.Begin()
	if i, v := it.Index(), it.Value(); i != -1 || v != "" {
		t.Errorf("Begin(): get (%d, %s), expect (%d, %s)", i, v, -1, "")
	}
	it.First()
	if i, v := it.Index(), it.Value(); i != 0 || v != "apples" {
		t.Errorf("Begin(): get (%d, %s), expect (%d, %s)", i, v, 0, "apples")
	}
	it.End()
	if i, v := it.Index(), it.Value(); i != 5 || v != "" {
		t.Errorf("Begin(): get (%d, %s), expect (%d, %s)", i, v, 6, "")
	}
	it.Last()
	if i, v := it.Index(), it.Value(); i != 4 || v != "potatoes" {
		t.Errorf("Begin(): get (%d, %s), expect (%d, %s)", i, v, 4, "potatoes")
	}
}

func TestListIteratorNextTo(t *testing.T) {
	tests := []struct {
		input []string
		index int
		value string
		f     func(i int, v string) bool
	}{
		{
			input: []string{"apples", "pears", "cucumbers", "onions", "potatoes"},
			index: 3,
			value: "onions",
			f: func(i int, v string) bool {
				return v == "onions"
			},
		},
		{
			input: []string{"apples", "pears", "cucumbers", "onions", "potatoes"},
			index: 1,
			value: "pears",
			f: func(i int, v string) bool {
				return i == 1
			},
		},
	}
	for _, tst := range tests {
		l := New[string](tst.input...)
		it := l.Iterator()
		it.NextTo(tst.f)
		if i, v := it.Index(), it.Value(); i != tst.index || v != tst.value {
			t.Errorf("NextTo(): get (%d, %s), expect (%d, %s)", i, v, tst.index, tst.value)
		}
	}
}

func TestListIteratorPrevTo(t *testing.T) {
	tests := []struct {
		input []string
		index int
		value string
		f     func(i int, v string) bool
	}{
		{
			input: []string{"apples", "pears", "cucumbers", "onions", "potatoes"},
			index: 3,
			value: "onions",
			f: func(i int, v string) bool {
				return v == "onions"
			},
		},
		{
			input: []string{"apples", "pears", "cucumbers", "onions", "potatoes"},
			index: 1,
			value: "pears",
			f: func(i int, v string) bool {
				return i == 1
			},
		},
	}
	for _, tst := range tests {
		l := New[string](tst.input...)
		it := l.Iterator()
		it.End()
		it.PrevTo(tst.f)
		if i, v := it.Index(), it.Value(); i != tst.index || v != tst.value {
			t.Errorf("NextTo(): get (%d, %s), expect (%d, %s)", i, v, tst.index, tst.value)
		}
	}
}
