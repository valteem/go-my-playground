package testutil

import (
	"math"

	"github.com/google/go-cmp/cmp"
)

func EquateInt64Approx(margin int64) cmp.Option {
	return cmp.Comparer(func(i, j int64) bool {
		return math.Abs(float64(i-j)) <= float64(margin)
	})
}
