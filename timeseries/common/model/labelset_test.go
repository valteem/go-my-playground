package model

import (
	"testing"
)

func TestLabelSetEqual(t *testing.T) {
	tests := []struct {
		descr string
		ls    LabesSet
		ls1   LabesSet
		equal bool
	}{
		{
			descr: "identical label sets",
			ls: LabesSet{"label1": "value1",
				"label2": "value2"},
			ls1: LabesSet{"label1": "value1",
				"label2": "value2"},
			equal: true,
		},
		{
			descr: "different number of labels",
			ls: LabesSet{"label1": "value1",
				"label2": "value2"},
			ls1: LabesSet{"label1": "value1",
				"label2": "value2",
				"label3": "value3",
			},
			equal: false,
		},
		{
			descr: "same length. different values",
			ls: LabesSet{"label1": "value1",
				"label2": "value2"},
			ls1: LabesSet{"label1": "value1",
				"label2": "value22"},
			equal: false,
		},
		{
			descr: "same length, different names",
			ls: LabesSet{"label1": "value1",
				"label2": "value2"},
			ls1: LabesSet{"label11": "value1",
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
		ls     LabesSet
		ls1    LabesSet
		before bool
	}{
		{
			descr: "identical label sets",
			ls: LabesSet{"label1": "value1",
				"label2": "value2"},
			ls1: LabesSet{"label1": "value1",
				"label2": "value2"},
			before: false,
		},
		{
			descr: "ls has fewer labels",
			ls:    LabesSet{"label1": "value1"},
			ls1: LabesSet{"label1": "value1",
				"label2": "value2"},
			before: true,
		},
		{
			descr: "ls has more labels",
			ls: LabesSet{"label1": "value1",
				"label2": "value2"},
			ls1:    LabesSet{"label1": "value1"},
			before: false,
		},
		{
			descr: "same label names, last value of ls is before",
			ls: LabesSet{"label1": "value1",
				"label2": "1value2"},
			ls1: LabesSet{"label1": "value1",
				"label2": "2value2"},
			before: true,
		},
		{
			descr: "same label names, last value of ls is after",
			ls: LabesSet{"label1": "value1",
				"label2": "2value2"},
			ls1: LabesSet{"label1": "value1",
				"label2": "1value2"},
			before: false,
		},
		{
			descr: "same label names, first value of ls is before",
			ls: LabesSet{"label1": "1value1",
				"label2": "value2"},
			ls1: LabesSet{"label1": "2value1",
				"label2": "value2"},
			before: true,
		},
		{
			descr: "same label names, first value of ls is after",
			ls: LabesSet{"label1": "2value1",
				"label2": "2value2"},
			ls1: LabesSet{"label1": "1value1",
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
