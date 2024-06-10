package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLabelSetEqual(t *testing.T) {
	tests := []struct {
		descr string
		ls    LabelSet
		ls1   LabelSet
		equal bool
	}{
		{
			descr: "identical label sets",
			ls: LabelSet{"label1": "value1",
				"label2": "value2"},
			ls1: LabelSet{"label1": "value1",
				"label2": "value2"},
			equal: true,
		},
		{
			descr: "different number of labels",
			ls: LabelSet{"label1": "value1",
				"label2": "value2"},
			ls1: LabelSet{"label1": "value1",
				"label2": "value2",
				"label3": "value3",
			},
			equal: false,
		},
		{
			descr: "same length. different values",
			ls: LabelSet{"label1": "value1",
				"label2": "value2"},
			ls1: LabelSet{"label1": "value1",
				"label2": "value22"},
			equal: false,
		},
		{
			descr: "same length, different names",
			ls: LabelSet{"label1": "value1",
				"label2": "value2"},
			ls1: LabelSet{"label11": "value1",
				"label2": "value2"},
			equal: false,
		},
	}
	for _, tc := range tests {
		if eq := tc.ls.Equal(tc.ls1); eq != tc.equal {
			t.Errorf("%s Equal(): get %t, expect %t", tc.descr, eq, tc.equal)
		}
	}
}
func TestLabelSetBefore(t *testing.T) {
	tests := []struct {
		descr  string
		ls     LabelSet
		ls1    LabelSet
		before bool
	}{
		{
			descr: "identical label sets",
			ls: LabelSet{"label1": "value1",
				"label2": "value2"},
			ls1: LabelSet{"label1": "value1",
				"label2": "value2"},
			before: false,
		},
		{
			descr: "ls has fewer labels",
			ls:    LabelSet{"label1": "value1"},
			ls1: LabelSet{"label1": "value1",
				"label2": "value2"},
			before: true,
		},
		{
			descr: "ls has more labels",
			ls: LabelSet{"label1": "value1",
				"label2": "value2"},
			ls1:    LabelSet{"label1": "value1"},
			before: false,
		},
		{
			descr: "same label names, last value of ls is before",
			ls: LabelSet{"label1": "value1",
				"label2": "1value2"},
			ls1: LabelSet{"label1": "value1",
				"label2": "2value2"},
			before: true,
		},
		{
			descr: "same label names, last value of ls is after",
			ls: LabelSet{"label1": "value1",
				"label2": "2value2"},
			ls1: LabelSet{"label1": "value1",
				"label2": "1value2"},
			before: false,
		},
		{
			descr: "same label names, first value of ls is before",
			ls: LabelSet{"label1": "1value1",
				"label2": "value2"},
			ls1: LabelSet{"label1": "2value1",
				"label2": "value2"},
			before: true,
		},
		{
			descr: "same label names, first value of ls is after",
			ls: LabelSet{"label1": "2value1",
				"label2": "2value2"},
			ls1: LabelSet{"label1": "1value1",
				"label2": "1value2"},
			before: false,
		},
	}
	for _, tc := range tests {
		if b := tc.ls.Before(tc.ls1); b != tc.before {
			t.Errorf("%s Before(): get %t, expect %t", tc.descr, b, tc.before)
		}
	}
}

func TestLabelSetClone(t *testing.T) {
	ls := LabelSet{
		"label1": "value1",
		"label2": "value2",
		"label3": "value3",
	}
	ls1 := ls.Clone()
	if !ls.Equal(ls1) {
		t.Errorf("Expect two equal label sets, get %v, %v", ls, ls1)
	}
}

func TestLabelSetMerge(t *testing.T) {
	tests := []struct {
		descr string
		ls1   LabelSet
		ls2   LabelSet
		lsm   LabelSet
	}{
		{
			descr: "unique names and values",
			ls1: LabelSet{"name1": "value1",
				"name2": "value2"},
			ls2: LabelSet{"name3": "value3",
				"name4": "value4"},
			lsm: LabelSet{"name1": "value1",
				"name2": "value2",
				"name3": "value3",
				"name4": "value4"},
		},
		{
			descr: "one same label",
			ls1: LabelSet{"name1": "value1",
				"name2": "value2"},
			ls2: LabelSet{"name2": "value2",
				"name3": "value3"},
			lsm: LabelSet{"name1": "value1",
				"name2": "value2",
				"name3": "value3"},
		},
		{
			descr: "two values for the same name", // ls2 value goes to output label set
			ls1: LabelSet{"name1": "value1",
				"name2": "value2"},
			ls2: LabelSet{"name2": "value22",
				"name3": "value3"},
			lsm: LabelSet{"name1": "value1",
				"name2": "value22",
				"name3": "value3"},
		},
	}
	for _, tc := range tests {
		if lsm := tc.ls1.Merge(tc.ls2); !cmp.Equal(lsm, tc.lsm) {
			t.Errorf("%s Merge(): get %v, expect %v", tc.descr, lsm, tc.lsm)
		}

	}
}
