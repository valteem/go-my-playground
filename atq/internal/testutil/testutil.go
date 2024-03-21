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
