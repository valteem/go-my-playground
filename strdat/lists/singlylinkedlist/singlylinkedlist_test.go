package singlylinkedlist

import (
	stdcmp "cmp"
	"encoding/json"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRemove(t *testing.T) {

	list := New[int](0, 1, 2, 3, 4)

	list.Remove(0)
	if v, b := list.Get(0); v != 1 || b != true {
		t.Errorf("Get (%v, %t) at the head, expect (1, true", v, b)
	}

}

func TestValues(t *testing.T) {

	list := New[int]()
	expect := []int{}
	for i := 0; i < 1<<16; i++ {
		list.Add(i)
		expect = append(expect, i)
	}
	if v := list.Values(); !cmp.Equal(v, expect) {
		t.Errorf("Get %v, expect %v", v, expect)
	}
}

func TestListNew(t *testing.T) {

	list := New[int]()
	if actualOutput, expectedOutput := list.Empty(), true; actualOutput != expectedOutput {
		t.Errorf("list.Empty(): get %t expect %t", actualOutput, expectedOutput)
	}

	list = New[int](1, 2, 3)
	if actualOutput, expectedOutput := list.Size(), 3; actualOutput != expectedOutput {
		t.Errorf("list.Size(): get %d, expect %d", actualOutput, expectedOutput)
	}
	tests := []struct {
		index          int
		expectedOutput int
		expectedOk     bool
	}{
		{
			index:          0,
			expectedOutput: 1,
			expectedOk:     true,
		},
		{
			index:          1,
			expectedOutput: 2,
			expectedOk:     true,
		},
		{
			index:          2,
			expectedOutput: 3,
			expectedOk:     true,
		},
	}
	for _, tst := range tests {
		actualOutput, actualOk := list.Get(tst.index)
		if actualOutput != tst.expectedOutput || actualOk != tst.expectedOk {
			t.Errorf("list.Get(%d): get (%d, %t), expect (%d, %t)", tst.index, actualOutput, actualOk, tst.expectedOutput, tst.expectedOk)
		}
	}

}

func TestListAdd(t *testing.T) {

	list := New[string]("a")
	list.Add("b", "c")
	list.Add("x")

	if actualOutput, expectedOutput := list.Empty(), false; actualOutput != expectedOutput {
		t.Errorf("list.Empty(): get %t, expect %t", actualOutput, expectedOutput)
	}

	if actualOutput, expectedOutput := list.Size(), 4; actualOutput != expectedOutput {
		t.Errorf("list.Size(): get %d, expect %d", actualOutput, expectedOutput)
	}

	tests := []struct {
		index          int
		expectedOutput string
		expectedOk     bool
	}{
		{
			index:          0,
			expectedOutput: "a",
			expectedOk:     true,
		},
		{
			index:          1,
			expectedOutput: "b",
			expectedOk:     true,
		},
		{
			index:          2,
			expectedOutput: "c",
			expectedOk:     true,
		},
		{
			index:          3,
			expectedOutput: "x",
			expectedOk:     true,
		},
		{
			index:          4,
			expectedOutput: "",
			expectedOk:     false,
		},
	}
	for _, tst := range tests {
		actualOutput, actualOk := list.Get(tst.index)
		if actualOutput != tst.expectedOutput || actualOk != tst.expectedOk {
			t.Errorf("list.Get(%d): get (%s, %t), expect (%s, %t)", tst.index, actualOutput, actualOk, tst.expectedOutput, tst.expectedOk)
		}
	}
}

func TestListAppendPrepend(t *testing.T) {

	list := New[string]("j")
	list.Append("k")
	list.Prepend("i")

	if actualOutput, expectedOutput := list.Empty(), false; actualOutput != expectedOutput {
		t.Errorf("list.Empty(): get %t, expect %t", actualOutput, expectedOutput)
	}

	if actualOutput, expectedOutput := list.Size(), 3; actualOutput != expectedOutput {
		t.Errorf("list.Size(): get %d, expect %d", actualOutput, expectedOutput)
	}

	tests := []struct {
		index          int
		expectedOutput string
		expectedOk     bool
	}{
		{
			index:          0,
			expectedOutput: "i",
			expectedOk:     true,
		},
		{
			index:          1,
			expectedOutput: "j",
			expectedOk:     true,
		},
		{
			index:          2,
			expectedOutput: "k",
			expectedOk:     true,
		},
		{
			index:          3,
			expectedOutput: "",
			expectedOk:     false,
		},
	}
	for _, tst := range tests {
		actualOutput, actualOk := list.Get(tst.index)
		if actualOutput != tst.expectedOutput || actualOk != tst.expectedOk {
			t.Errorf("list.Get(%d): get (%s, %t), expect (%s, %t)", tst.index, actualOutput, actualOk, tst.expectedOutput, tst.expectedOk)
		}
	}

}

func TestListRemove(t *testing.T) {

	list := New[string]("i", "j", "k")

	list.Remove(1)
	if actualOutput, expectedOutput := list.Empty(), false; actualOutput != expectedOutput {
		t.Errorf("list.Empty(): get %t, expect %t", actualOutput, expectedOutput)
	}
	if actualOutput, expectedOutput := list.Size(), 2; actualOutput != expectedOutput {
		t.Errorf("list.Size(): get %d, expect %d", actualOutput, expectedOutput)
	}
	tests := []struct {
		index          int
		expectedOutput string
	}{
		{
			index:          0,
			expectedOutput: "i",
		},
		{
			index:          1,
			expectedOutput: "k",
		},
	}
	for _, tst := range tests {
		actualOutput, _ := list.Get(tst.index)
		if actualOutput != tst.expectedOutput {
			t.Errorf("list.Get(%d): get %s, expect %s", tst.index, actualOutput, tst.expectedOutput)
		}
	}

	list.Remove(0)
	if actualOutput, expectedOutput := list.Empty(), false; actualOutput != expectedOutput {
		t.Errorf("list.Empty(): get %t, expect %t", actualOutput, expectedOutput)
	}
	if actualOutput, expectedOutput := list.Size(), 1; actualOutput != expectedOutput {
		t.Errorf("list.Empty(): get %d, expect %d", actualOutput, expectedOutput)
	}
	actualOutput, _ := list.Get(0)
	index, expectedOutput := 0, "k"
	if actualOutput != expectedOutput {
		t.Errorf("list.GetO(%d): get %s, expect %s", index, actualOutput, expectedOutput)
	}

	list.Remove(1) // out-of bounds, nothing happens
	if actualOutput, expectedOutput := list.Empty(), false; actualOutput != expectedOutput {
		t.Errorf("list.Empty(): get %t, expect %t", actualOutput, expectedOutput)
	}
	if actualOutput, expectedOutput := list.Size(), 1; actualOutput != expectedOutput {
		t.Errorf("list.Empty(): get %d, expect %d", actualOutput, expectedOutput)
	}
	actualOutput, _ = list.Get(0)
	index, expectedOutput = 0, "k"
	if actualOutput != expectedOutput {
		t.Errorf("list.GetO(%d): get %s, expect %s", index, actualOutput, expectedOutput)
	}

	list.Remove(0)
	if actualOutput, expectedOutput := list.Empty(), true; actualOutput != expectedOutput {
		t.Errorf("list.Empty(): get %t, expect %t", actualOutput, expectedOutput)
	}
	if actualOutput, expectedOutput := list.Size(), 0; actualOutput != expectedOutput {
		t.Errorf("list.Empty(): get %d, expect %d", actualOutput, expectedOutput)
	}
	actualOutput, actualOk := list.Get(0)
	index, expectedOutput, expectedOk := 0, "", false
	if actualOutput != expectedOutput || actualOk != expectedOk {
		t.Errorf("list.GetO(%d): get (%s, %t), expect (%s, %t)", index, actualOutput, actualOk, expectedOutput, expectedOk)
	}

	list.Remove(0) // empty list, nothing happens
	if actualOutput, expectedOutput := list.Empty(), true; actualOutput != expectedOutput {
		t.Errorf("list.Empty(): get %t, expect %t", actualOutput, expectedOutput)
	}
	if actualOutput, expectedOutput := list.Size(), 0; actualOutput != expectedOutput {
		t.Errorf("list.Empty(): get %d, expect %d", actualOutput, expectedOutput)
	}
	actualOutput, actualOk = list.Get(0)
	index, expectedOutput, expectedOk = 0, "", false
	if actualOutput != expectedOutput || actualOk != expectedOk {
		t.Errorf("list.GetO(%d): get (%s, %t), expect (%s, %t)", index, actualOutput, actualOk, expectedOutput, expectedOk)
	}
}

func TestListSwap(t *testing.T) {

	list := New[string]("a", "b", "c")

	list.Swap(0, 2)
	if actual, expected := list.Values(), []string{"c", "b", "a"}; !cmp.Equal(actual, expected) {
		t.Errorf("list.Swap(): get %v, expect %v", actual, expected)
	}

	list.Swap(0, 3) // out-of-bounds, nothing happens
	if actual, expected := list.Values(), []string{"c", "b", "a"}; !cmp.Equal(actual, expected) {
		t.Errorf("list.Swap(): get %v, expect %v", actual, expected)
	}

}

func TestListSort(t *testing.T) {
	list := New[string]("x", "k", "v", "a", "y", "m")
	list.Sort(stdcmp.Compare[string])
	if actual, expected := list.Values(), []string{"a", "k", "m", "v", "x", "y"}; !cmp.Equal(actual, expected) {
		t.Errorf("list.Sort(): get %v, expect %v", actual, expected)
	}
}

func TestListClear(t *testing.T) {
	list := New[string]("k", "u", "x")
	list.Clear()
	if actualOutput, expectedOutput := list.Empty(), true; actualOutput != expectedOutput {
		t.Errorf("list.Empty(): get %t, expect %t", actualOutput, expectedOutput)
	}
	if actualOutput, expectedOutput := list.Size(), 0; actualOutput != expectedOutput {
		t.Errorf("list.Size(): get %d, expect %d", actualOutput, expectedOutput)
	}
	index, expectedOutput, expectedOk := 0, "", false
	actualOutput, actualOk := list.Get(index)
	if actualOutput != expectedOutput || actualOk != expectedOk {
		t.Errorf("list.Get(%d): get (%s, %t), expect (%s, %t)", index, actualOutput, actualOk, expectedOutput, expectedOk)
	}
}

func TestListContains(t *testing.T) {

	list := New[string]("a", "b", "c")

	tests := []struct {
		value          string
		expectedOutput bool
	}{
		{
			value:          "a",
			expectedOutput: true,
		},
		{
			value:          "b",
			expectedOutput: true,
		},
		{
			value:          "c",
			expectedOutput: true,
		},
		{
			value:          "x",
			expectedOutput: false,
		},
		{
			value:          "y",
			expectedOutput: false,
		},
	}

	for _, tst := range tests {
		if actualOutput := list.Contains(tst.value); actualOutput != tst.expectedOutput {
			t.Errorf("list.Contains(%s): get %t, expect %t", tst.value, actualOutput, tst.expectedOutput)
		}
	}

}

func TestListValues(t *testing.T) {
	list := New[string]("i", "j", "k")
	if actual, expected := list.Values(), []string{"i", "j", "k"}; !cmp.Equal(actual, expected) {
		t.Errorf("list.Values(): get %v, expect %v", actual, expected)
	}
}

func TestListIndexOf(t *testing.T) {

	list := New[string]("a", "b", "c")

	tests := []struct {
		value          string
		expectedOutput int
	}{
		{
			value:          "a",
			expectedOutput: 0,
		},
		{
			value:          "b",
			expectedOutput: 1,
		},
		{
			value:          "c",
			expectedOutput: 2,
		},
		{
			value:          "x",
			expectedOutput: -1,
		},
		{
			value:          "bb",
			expectedOutput: -1,
		},
		{
			value:          "",
			expectedOutput: -1,
		},
	}

	for _, tst := range tests {
		if actualOutput := list.IndexOf(tst.value); actualOutput != tst.expectedOutput {
			t.Errorf("list.IndexOf(%s): get %d, expect %d", tst.value, actualOutput, tst.expectedOutput)
		}
	}

}

func TestListInsert(t *testing.T) {

	list := New[string]("a", "b", "c")

	list.Insert(0, "") // prepend
	if actual, expected := list.Values(), []string{"", "a", "b", "c"}; !cmp.Equal(actual, expected) {
		t.Errorf("list.Insert(): get %v, expect %v", actual, expected)
	}

	list.Insert(4, "x") // append
	if actual, expected := list.Values(), []string{"", "a", "b", "c", "x"}; !cmp.Equal(actual, expected) {
		t.Errorf("list.Insert(): get %v, expect %v", actual, expected)
	}

	list.Insert(3, "bb")
	if actual, expected := list.Values(), []string{"", "a", "b", "bb", "c", "x"}; !cmp.Equal(actual, expected) {
		t.Errorf("list.Insert(): get %v, expect %v", actual, expected)
	}

	list.Insert(7, "u") // out-of-bounds, nothing happens
	if actual, expected := list.Values(), []string{"", "a", "b", "bb", "c", "x"}; !cmp.Equal(actual, expected) {
		t.Errorf("list.Insert(): get %v, expect %v", actual, expected)
	}

}

func TestListSet(t *testing.T) {

	list := New[string]()

	list.Set(0, "a")
	list.Set(1, "b") // append
	list.Set(2, "c") // append
	if actual, expected := list.Values(), []string{"a", "b", "c"}; !cmp.Equal(actual, expected) {
		t.Errorf("list.Set(): get %v, expect %v", actual, expected)
	}

	list.Set(0, "x")
	list.Set(2, "y")
	if actual, expected := list.Values(), []string{"x", "b", "y"}; !cmp.Equal(actual, expected) {
		t.Errorf("list.Set(): get %v, expect %v", actual, expected)
	}

	list.Set(4, "u") //out-of-bounds, nothing happens
	if actual, expected := list.Values(), []string{"x", "b", "y"}; !cmp.Equal(actual, expected) {
		t.Errorf("list.Set(): get %v, expect %v", actual, expected)
	}

}

func TestListEach(t *testing.T) {

	list := New[string]("a", "b", "c")

	a := []string{"a", "b", "c"}

	list.Each(func(index int, value string) {
		if actual := a[index]; actual != value {
			t.Errorf("list.Each(%d): get %s, expect %s", index, actual, value)
		}
	})

}

func TestListMap(t *testing.T) {

	list := New[string]("a", "b", "c")

	mapped := list.Map(func(index int, value string) string {
		return value + value
	})

	if actual, expected := mapped.Values(), []string{"aa", "bb", "cc"}; !cmp.Equal(actual, expected) {
		t.Errorf("list.Map(): get %v, expect %v", actual, expected)
	}

}

func TestListSelect(t *testing.T) {

	list := New[string]("a", "f", "m", "t", "v", "x")

	selected := list.Select(func(index int, value string) bool {
		return (index >= 2) && (value < "w")
	})

	if actual, expected := selected.Values(), []string{"m", "t", "v"}; !cmp.Equal(actual, expected) {
		t.Errorf("list.Select(): get %v, expect %v", actual, expected)
	}

}

func TestListAny(t *testing.T) {

	list := New[string]("a", "b", "c")

	a := list.Any(func(index int, value string) bool {
		return value <= "c"
	})
	if !a {
		t.Errorf("list.Any(): get %t, expect %t", a, true)
	}

	a = list.Any(func(index int, value string) bool {
		return value >= "x"
	})
	if a {
		t.Errorf("list.Any(): get %t, expect %t", a, false)
	}

}

func TestListAll(t *testing.T) {

	list := New[string]("a", "b", "c")

	all := list.All(func(index int, value string) bool {
		return value >= "a" && value <= "c"
	})
	if !all {
		t.Errorf("list.All(): get %t, expect %t", all, true)
	}

	all = list.All(func(index int, value string) bool {
		return value >= "x" && value <= "y"
	})
	if all {
		t.Errorf("list.All(): get %t, expect %t", all, false)
	}

}

func TestListFind(t *testing.T) {

	list := New[string]("a", "b", "c")

	index, value := list.Find(func(index int, value string) bool {
		return index == 1 && value == "b"
	})
	if index != 1 || value != "b" {
		t.Errorf("list.Find(): get (%d, %s), expect (%d, %s)", index, value, 1, "b")
	}

	index, value = list.Find(func(index int, value string) bool {
		return index > 10 && value > "u"
	})
	if index != -1 || value != "" {
		t.Errorf("list.Find(): get (%d, %s), expect (%d, %s)", index, value, -1, "")
	}
}

// `Chaining`, Enumerable{} (and probably some more stuff) actually come from Ruby (https://stackoverflow.com/q/70818283)
func TestListChaining(t *testing.T) {
	list := New[string]("a", "b", "c", "x", "y")
	chainedList := list.Select(func(index int, value string) bool {
		return value > "a"
	}).Map(func(index int, value string) string {
		return value + value
	})
	if actual, expected := chainedList.Size(), 4; actual != expected {
		t.Errorf("list.Size(): get %d, expect %d", actual, expected)
	}
	tests := []struct {
		index         int
		expectedValue string
		expectedOk    bool
	}{
		{
			index:         0,
			expectedValue: "bb",
			expectedOk:    true,
		},
		{
			index:         1,
			expectedValue: "cc",
			expectedOk:    true,
		},
		{
			index:         2,
			expectedValue: "xx",
			expectedOk:    true,
		},
		{
			index:         3,
			expectedValue: "yy",
			expectedOk:    true,
		},
		{
			index:         4,
			expectedValue: "",
			expectedOk:    false,
		},
	}
	for _, tst := range tests {
		actualValue, actualOk := chainedList.Get(tst.index)
		if actualValue != tst.expectedValue || actualOk != tst.expectedOk {
			t.Errorf("list.Get(%d): get (%s, %t), expect (%s, %t", tst.index, actualValue, actualOk, tst.expectedValue, tst.expectedOk)
		}
	}
}

func TestListIteratorNextOverEmpty(t *testing.T) {
	list := New[string]()
	it := list.Iterator()
	for it.Next() {
		t.Errorf("cannot iterate over empty list")
	}
}
func TestListIteratorNext(t *testing.T) {
	list := New[string]("a", "b", "c")
	it := list.Iterator()
	count := 0
	var index int
	var actualValue string
	elements := []string{"a", "b", "c"}
	length := len(elements)
	for it.Next() {
		count++
		index = it.Index()
		actualValue = it.Value()
		if index >= length {
			t.Errorf("it.Next(): index %d out of bounds (should be less than %d)", index, length)
		}
		if expectedValue := elements[index]; actualValue != expectedValue {
			t.Errorf("it.Next() on count %d (index %d): get %s, expect %s", count, index, actualValue, expectedValue)
		}
	}
}

func TestListIteratorBegin(t *testing.T) {
	list := New[string]()
	it := list.Iterator()
	it.Begin()
	list.Add("a", "b", "c")
	it.Next()
	if actualIndex, expectedIndex, actualValue, expectedValue := it.Index(), 0, it.Value(), "a"; actualIndex != expectedIndex || actualValue != expectedValue {
		t.Errorf("it.Next() after it.Begin(): get %s at %d, expect %s at %d", actualValue, actualIndex, expectedValue, expectedIndex)
	}
}

func TestListIteratorFirst(t *testing.T) {
	list := New[string]()
	it := list.Iterator()
	if actual, expected := it.First(), false; actual != expected {
		t.Errorf("empty list - First() must return `false`")
	}
	list.Add("a", "b", "c")
	if actual, expected := it.First(), true; actual != expected {
		t.Errorf("non-empty list - First() must return `true`")
	}
	if actualIndex, actualValue, expectedIndex, expectedValue := it.Index(), it.Value(), 0, "a"; actualIndex != expectedIndex || actualValue != expectedValue {
		t.Errorf("get %s at %d, expect %s at %d", actualValue, actualIndex, expectedValue, expectedIndex)
	}
}

func TestListIteratorNextTo(t *testing.T) {
	lookup := func(index int, value string) bool {
		return strings.HasPrefix(value, "b")
	}

	{
		list := New[string]()
		it := list.Iterator()
		for it.NextTo(lookup) {
			t.Errorf("cannot iterate over empty list")
		}
	}

	{
		list := New[string]("a", "c", "x")
		it := list.Iterator()
		for it.NextTo(lookup) {
			t.Errorf("it.NextTo() cannot return 'true` - lookup functions should find nothing")
		}
	}

	{
		list := New[string]("a", "b", "c")
		it := list.Iterator()
		it.Begin()
		if !it.NextTo(lookup) {
			t.Errorf("it.NextTo should return `true` - lookup value is in the list")
		}
		if actualIndex, actualValue, expectedIndex, expectedValue := it.Index(), it.Value(), 1, "b"; actualIndex != expectedIndex || actualValue != expectedValue {
			t.Errorf("get %s at %d, expect %s at %d", actualValue, actualIndex, expectedValue, expectedIndex)
		}
	}
}

func TestListSerialization(t *testing.T) {
	list := New[string]("a", "b", "c")
	var err error
	assert := func() {
		if actualValues, expectedValues := list.Values(), []string{"a", "b", "c"}; !cmp.Equal(actualValues, expectedValues) {
			t.Errorf("list.Values(): get %v, expect %v", actualValues, expectedValues)
		}
		if actualSize, expectedSize := list.Size(), 3; actualSize != expectedSize {
			t.Errorf("list.Size(): get %d, expect %d", actualSize, expectedSize)
		}
		if err != nil {
			t.Errorf("%v", err)
		}
	}

	assert()

	bytes, err := list.ToJSON()
	assert()

	err = list.FromJSON(bytes)
	assert()

	bytes, err = json.Marshal(list)
	assert()

	list.Clear()
	err = json.Unmarshal(bytes, list) // `list` is already a pointer to 'singlylinkedlist` struct, not a struct itself (provided by New())
	assert()
}
