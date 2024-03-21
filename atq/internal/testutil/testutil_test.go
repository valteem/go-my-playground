package testutil

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/valteem/atq/internal/base"
)

func TestEquateInt64ApproxEqual(t *testing.T) {
	tests := []struct {
		description string
		v1          int64
		v2          int64
		margin      int64
	}{
		{
			description: "exact equality: zero margin",
			v1:          31,
			v2:          31,
			margin:      0,
		}, {
			description: "exact equality: non-zero margin",
			v1:          31,
			v2:          31,
			margin:      8,
		},
		{
			description: "equality within non-zero margin",
			v1:          31,
			v2:          35,
			margin:      8,
		},
	}

	for _, tst := range tests {
		opt := EquateInt64Approx(tst.margin)
		if !cmp.Equal(tst.v1, tst.v2, opt) {
			t.Errorf("%s: %v must be equal %v within margin %v", tst.description, tst.v1, tst.v2, tst.margin)
		}

	}
}

func TestEquateInt64ApproxNotEqual(t *testing.T) {
	tests := []struct {
		description string
		v1          int64
		v2          int64
		margin      int64
	}{
		{
			description: "exact equality: zero margin",
			v1:          31,
			v2:          32,
			margin:      0,
		}, {
			description: "exact equality: non-zero margin",
			v1:          31,
			v2:          33,
			margin:      1,
		},
		{
			description: "equality within non-zero margin",
			v1:          31,
			v2:          35,
			margin:      3,
		},
	}

	for _, tst := range tests {
		opt := EquateInt64Approx(tst.margin)
		if cmp.Equal(tst.v1, tst.v2, opt) {
			t.Errorf("%s: %v cannot be equal %v within margin %v", tst.description, tst.v1, tst.v2, tst.margin)
		}

	}
}

func TestSortMsgOptEqual(t *testing.T) {
	tests := []struct {
		description   string
		taskMessagesA []*base.TaskMessage
		taskMessagesB []*base.TaskMessage
	}{
		{
			description: "sort order ascending vs descending",
			taskMessagesA: []*base.TaskMessage{
				{ID: "11"}, // https://pkg.go.dev/cmd/gofmt#hdr-The_simplify_command
				{ID: "21"},
				{ID: "31"},
			},
			taskMessagesB: []*base.TaskMessage{
				{ID: "31"},
				{ID: "21"},
				{ID: "11"},
			},
		},
		{
			description: "shuffled order",
			taskMessagesA: []*base.TaskMessage{
				{ID: "31"},
				{ID: "21"},
				{ID: "41"},
				{ID: "11"},
			},
			taskMessagesB: []*base.TaskMessage{
				{ID: "21"},
				{ID: "11"},
				{ID: "41"},
				{ID: "31"},
			},
		},
	}
	for _, tst := range tests {
		if !cmp.Equal(tst.taskMessagesA, tst.taskMessagesB, SortMsgOpt) {
			t.Errorf("%s:\n%v\n%v\nshould be equal", tst.description, tst.taskMessagesA, tst.taskMessagesB)
		}
	}
}

func TestSortMsgOptNotEqual(t *testing.T) {
	tests := []struct {
		description   string
		taskMessagesA []*base.TaskMessage
		taskMessagesB []*base.TaskMessage
	}{
		{
			description: "same length, different values",
			taskMessagesA: []*base.TaskMessage{
				{ID: "11"},
				{ID: "21"},
				{ID: "31"},
			},
			taskMessagesB: []*base.TaskMessage{
				{ID: "12"},
				{ID: "22"},
				{ID: "32"},
			},
		},
		{
			description: "different length",
			taskMessagesA: []*base.TaskMessage{
				{ID: "31"},
				{ID: "21"},
				{ID: "41"},
				{ID: "11"},
			},
			taskMessagesB: []*base.TaskMessage{
				{ID: "21"},
				{ID: "11"},
				{ID: "41"},
				{ID: "31"},
				{ID: "51"},
			},
		},
	}
	for _, tst := range tests {
		if cmp.Equal(tst.taskMessagesA, tst.taskMessagesB, SortMsgOpt) {
			t.Errorf("%s:\n%v\n%v\nshould be equal", tst.description, tst.taskMessagesA, tst.taskMessagesB)
		}
	}
}
