package testutil

import (
	"math"
	"sort"

	"github.com/google/go-cmp/cmp"
	"github.com/valteem/atq/internal/base"
)

func EquateInt64Approx(margin int64) cmp.Option {
	return cmp.Comparer(func(i, j int64) bool {
		return math.Abs(float64(i-j)) <= float64(margin)
	})
}

// cmp.Option to compare two slices of task messages
var SortMsgOpt = cmp.Transformer("SortTaskMessages", func(in []*base.TaskMessage) []*base.TaskMessage {
	out := append([]*base.TaskMessage(nil), in...) //avoid mutating input
	sort.Slice(out, func(i, j int) bool {
		return out[i].ID < out[j].ID
	})
	return out
})

// cmp.Option to compare slices of base.Z entries
var SortZSetEntryOpt = cmp.Transformer("SortZSetEntries", func(in []base.Z) []base.Z {
	out := append([]base.Z(nil), in...) // The Transformer must not mutate <input> in any way (https://pkg.go.dev/github.com/google/go-cmp/cmp#Transformer)
	sort.Slice(out, func(i, j int) bool {
		return out[i].Message.ID < out[j].Message.ID
	})
	return out
})
