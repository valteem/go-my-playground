package arraylist

import (
	"encoding/json"
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

func TestString(t *testing.T) {
	tests := []struct {
		description string
		lst         *List[int]
		output      string
	}{
		{
			description: "list with values",
			lst:         New[int](0, 1, 2, 3, 4),
			output:      "ArrayList\n0, 1, 2, 3, 4",
		},
		{
			description: "empty list",
			lst:         New[int](),
			output:      "ArrayList\n",
		},
	}
	for _, tst := range tests {
		if r := tst.lst.String(); r != tst.output {
			t.Errorf("%s:, get %v, expect %v", tst.description, r, tst.output)
		}
	}
}

func TestIteratorNextEmpty(t *testing.T) {
	list := New[string]()
	it := list.Iterator()
	for it.Next() { // it.Next() immediately returns `false` on empty list, hence no error
		t.Errorf("Cannot iterate on empty list")
	}
}

func TestIteratorNext(t *testing.T) {
	list := New[string]("a", "b")
	list.Add("c")
	expected := []string{"a", "b", "c"}
	countExpected := 3
	it := list.Iterator()
	count := 0
	for it.Next() {
		count++
		index := it.Index()
		value := it.Value()
		if value != expected[index] {
			t.Errorf("get %v, expect %v", value, expected[index])
		}
	}
	if count != countExpected {
		t.Errorf("get %d, expect %d", count, countExpected)
	}
}

func TestIteratorPrevEmpty(t *testing.T) {
	list := New[string]()
	it := list.Iterator()
	for it.Prev() {
		t.Errorf("cannot iterate over empty list")
	}
}

func TestIteratorPrev(t *testing.T) {
	list := New[string]("a", "c")
	list.Insert(1, "b")
	it := list.Iterator()
	expected := []string{"a", "b", "c"}
	for it.Next() { // move to the end of the list
	}
	count := 0 // it.Index() = 3 at this point, right beyond last list element
	for it.Prev() {
		count++
		index := it.Index()
		value := it.Value()
		if value != expected[index] {
			t.Errorf("get value %v for index %d, expect %v", value, index, expected[index])
		}
	}
	if count != len(expected) {
		t.Errorf("get count %d, expect %d", count, len(expected))
	}
}

func TestIteratorBeginEndFirstLast(t *testing.T) {
	list := New[int](0, 1, 2, 3, 4)
	it := list.Iterator()
	idxBegin, idxEnd := -1, 5
	idxFirst, idxLast := 0, 4
	valueFirst, valueLast := 0, 4
	it.End()
	if it.Index() != idxEnd {
		t.Errorf("iterator index at the end should be %d, get %d", idxEnd, it.Index())
	}
	it.Begin()
	if it.Index() != idxBegin {
		t.Errorf("iterator index at the beginning should be %d, get %d", idxBegin, it.Index())
	}
	if last := it.Last(); last != true {
		t.Errorf("arraylist has last element, it.Last() should return `true`, get %t", last)
	} else {
		if i, v := it.Index(), it.Value(); i != idxLast || v != valueLast {
			t.Errorf("it.Last() moves to %v at position %d, expect %v at position %d", v, i, valueLast, idxLast)
		}
	}
	if first := it.First(); first != true {
		t.Errorf("arraylist has first element, it.First() should return `true`, get %t", first)
	} else {
		if i, v := it.Index(), it.Value(); i != idxFirst || v != valueFirst {
			t.Errorf("it.First moves to %v at position %d, expect %v at position %d", v, i, valueFirst, idxFirst)
		}
	}
}

func TestIteratorBeginEndFirstLastEmpty(t *testing.T) {
	list := New[string]()
	it := list.Iterator()

	// it.First() = 0, it.Last() = -1, looks pretty strange, last index before first
	idxBegin, idxFirst, idxLast, idxEnd := -1, 0, -1, 0

	it.Begin()
	if i := it.Index(); i != idxBegin {
		t.Errorf("it.Index() on empty list should return %d after it.Begin(), get %d", idxBegin, i)
	}

	first := it.First()
	if first != false {
		t.Errorf("it.First() on empty list should return `false`, get %t", first)
	}
	if i := it.Index(); i != idxFirst {
		t.Errorf("it.Index() on empty list should return %d after it.First(), get %d", idxFirst, i)
	}

	last := it.Last()
	if last != false {
		t.Errorf("it.Last() should return `false` on empty list, get %t", last)
	}
	if i := it.Index(); i != idxLast {
		t.Errorf("it.Index() on empty list should return %d after it.Last(), get %d", idxLast, i)
	}

	it.End()
	if i := it.Index(); i != idxEnd {
		t.Errorf("it.Index() on empty list should return %d after it.End(), get %d", idxEnd, i)
	}
}

func TestIteratorBeginEndFirstLastSingleElement(t *testing.T) {
	list := New[string]("a")
	it := list.Iterator()

	idxBegin, idxFirst, idxLast, idxEnd := -1, 0, 0, 1

	it.Begin()
	if i := it.Index(); i != idxBegin {
		t.Errorf("it.Index() on single element list should return %d after it.Begin(), get %d", idxBegin, i)
	}

	it.First()
	if i := it.Index(); i != idxFirst {
		t.Errorf("it.Index() on single element list should return %d after it.First(), get %d", idxFirst, i)
	}

	it.Last()
	if i := it.Index(); i != idxLast {
		t.Errorf("it.Index() on single element list should return %d after it.Last(), get %d", idxLast, i)
	}

	it.End()
	if i := it.Index(); i != idxEnd {
		t.Errorf("it.Index() on single element list should return %d after it.End(), get %d", idxEnd, i)
	}
}

func TestIteratorNextTo(t *testing.T) {

	minValue, idxFound, valueFound := 0, 0, 0
	f := func(i, v int) bool {
		return v >= minValue
	}

	// Empty list
	list := New[int]()
	it := list.Iterator()
	next := it.NextTo(f)
	if next == true {
		t.Errorf("cannot iterate over empty list")
	}

	// Not found
	list.Add(0, 1, 2, 3, 4)
	minValue = 5
	next = it.NextTo(f)
	if next == true {
		t.Errorf("Not found: should return `false`")
	}
	// Found
	minValue, idxFound, valueFound = 2, 2, 2
	it.Begin()
	next = it.NextTo(f)
	if next != true {
		t.Errorf("failed to move to position satisfying given condition ")
	} else {
		if i, v := it.Index(), it.Value(); i != idxFound || v != valueFound {
			t.Errorf("get %v at position %d, expect %v at position %d", v, i, valueFound, idxFound)
		}
	}
}

func TestIteratorPrevTo(t *testing.T) {

	maxValue, idxFound, valueFound := 0, 0, 0
	f := func(i, v int) bool {
		return v <= maxValue
	}

	list := New[int]()
	it := list.Iterator()

	// Empty list
	it.End()
	prev := it.PrevTo(f)
	if prev != false {
		t.Errorf("cannot iterate over empty list")
	}

	// Not found
	list.Add(0, 1, 2, 3, 4)
	maxValue = -1
	it.End()
	prev = it.PrevTo(f)
	if prev != false {
		t.Errorf("Not found: should return false")
	}

	// Found
	maxValue, idxFound, valueFound = 2, 2, 2
	it.End()
	prev = it.PrevTo(f)
	if prev != true {
		t.Errorf("failed to move to position satisfying given condition")
	} else {
		if i, v := it.Index(), it.Value(); i != idxFound || v != valueFound {
			t.Errorf("get %v at position %d, expect %v at position %d", v, i, valueFound, idxFound)
		}
	}

}

func TestListEach(t *testing.T) {
	list := New[int]()
	list.Add(0, 1, 2, 3, 4)
	expected := []int{0, 1, 2, 3, 4}
	list.Each(func(index int, value int) {
		if value != expected[index] {
			t.Errorf("Get %v at position %d, expect %v", value, index, expected[index])
		}
	})
}

func TestListMap(t *testing.T) {
	list := New[int](0, 1, 2, 3, 4)
	newList := list.Map(func(index int, value int) int {
		return value * value
	})
	expected := []int{0, 1, 4, 9, 16}
	if !cmp.Equal(newList.Values(), expected) {
		t.Errorf("Get %v, expect %v", newList.Values(), expected)
	}
}

func TestListSelect(t *testing.T) {
	list := New[int](0, 1, 2, 3, 4)
	newList := list.Select(func(index int, value int) bool {
		return ((value / 2) * 2) == value
	})
	expected := []int{0, 2, 4}
	if !cmp.Equal(newList.Values(), expected) {
		t.Errorf("Get %v, expect %v", newList.Values(), expected)
	}
}

func TestListAny(t *testing.T) {
	tests := []struct {
		description string
		lst         *List[int]
		output      bool
	}{
		{
			description: "Even only",
			lst:         New[int](2, 4, 6, 8),
			output:      true,
		},
		{
			description: "Uneven only",
			lst:         New[int](1, 3, 5, 7),
			output:      false,
		},
		{
			description: "Mix of even and uneven",
			lst:         New[int](1, 2, 3, 4),
			output:      true,
		},
	}
	for _, tst := range tests {
		a := tst.lst.Any(func(index, value int) bool {
			return ((value / 2) * 2) == value
		})
		if a != tst.output {
			t.Errorf("%s: get %t, expect %t", tst.description, a, tst.output)
		}
	}
}

func TestListAll(t *testing.T) {
	tests := []struct {
		description string
		input       *List[int]
		expected    bool
	}{
		{
			description: "Even only",
			input:       New[int](2, 4, 6, 8),
			expected:    true,
		},
		{
			description: "Uneven only",
			input:       New[int](1, 3, 5, 7),
			expected:    false,
		},
		{
			description: "Mix of even and uneven",
			input:       New[int](1, 2, 3, 4),
			expected:    false,
		},
	}
	for _, tst := range tests {
		output := tst.input.All(func(index, value int) bool {
			return ((value / 2) * 2) == value
		})
		if output != tst.expected {
			t.Errorf("%s: get %t, expect %t", tst.description, output, tst.expected)
		}
	}
}

func TestListFind(t *testing.T) {
	tests := []struct {
		description   string
		input         *List[int]
		expectedIndex int
		expectedValue int
	}{
		{
			description:   "First element found",
			input:         New[int](2, 3, 4, 5),
			expectedIndex: 0,
			expectedValue: 2,
		},
		{
			description:   "Last element found",
			input:         New[int](1, 3, 5, 8),
			expectedIndex: 3,
			expectedValue: 8,
		},
		{
			description:   "Element in the middle found",
			input:         New[int](1, 3, 4, 5, 7),
			expectedIndex: 2,
			expectedValue: 4,
		},
		{
			description:   "Not found",
			input:         New[int](1, 3, 5, 7),
			expectedIndex: -1,
			expectedValue: 0,
		},
	}
	for _, tst := range tests {
		outputIndex, outputvalue := tst.input.Find(func(i, v int) bool {
			return ((v / 2) * 2) == v
		})
		if outputIndex != tst.expectedIndex || outputvalue != tst.expectedValue {
			t.Errorf("%s, get %v at position %d, expect %v at position %d", tst.description, outputvalue, outputIndex, tst.expectedValue, tst.expectedIndex)
		}
	}
}

func TestListSerialization(t *testing.T) {

	list := New[int](0, 1, 2, 3, 4)

	var err error

	assert := func() {
		if output, expected := list.Values(), []int{0, 1, 2, 3, 4}; !cmp.Equal(output, expected) {
			t.Errorf("List values: get %v, expect %v", output, expected)
		}
		if output, expected := list.Size(), 5; output != expected {
			t.Errorf("List size: get %d, expect %d", output, expected)
		}
		if err != nil {
			t.Errorf("Get error %v", err)
		}
	}

	assert()

	bytes, err := list.ToJSON()
	assert()

	err = list.FromJSON(bytes)
	assert()

	_, err = json.Marshal([]any{list})
	if err != nil {
		t.Errorf("Error marshaling json: %v", err)
	}

	err = json.Unmarshal([]byte(`[0,1,2,3,4]`), list) // don't need a pointer to list here
	assert()

}
