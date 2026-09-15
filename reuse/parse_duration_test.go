package reuse_test

import (
	"testing"

	"github.com/valteem/reuse"
)

func TestParseDuration(t *testing.T) {

	tests := []struct {
		inp  float64
		fail bool
	}{
		{(float64(1<<63) + 1000.0), false},
		{(float64(1<<63) + 10000.0), true},
	}

	for _, tc := range tests {
		_, err := reuse.ParseDuration(tc.inp)
		if (tc.fail && err == nil) || (!tc.fail && err != nil) {
			t.Errorf("%f should fail", tc.inp)
		}
	}
}
