package singlylinkedlist

import (
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
