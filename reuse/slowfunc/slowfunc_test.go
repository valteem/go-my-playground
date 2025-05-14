package slowfunc

import (
	"context"
	"testing"
	"time"
)

// Generic anonymous structs not allowed
type testCase[T any] struct {
	spec    string
	f       func(args ...any) T
	timeout time.Duration
	output  T
}

func TestWrapSlowFunc(t *testing.T) {

	tests := []testCase[int]{
		{
			spec: "slow function returns before context timeout",
			f: func(args ...any) int {
				time.Sleep(100 * time.Millisecond)
				return 1
			},
			timeout: 150 * time.Millisecond,
			output:  1,
		},
		{
			spec: "slow function returns after context timeout",
			f: func(args ...any) int {
				time.Sleep(150 * time.Millisecond)
				return 1
			},
			timeout: 100 * time.Millisecond,
			output:  0,
		},
	}

	for _, tc := range tests {
		ctx, cancel := context.WithTimeout(context.Background(), tc.timeout)
		defer cancel()
		output, _ := WrapSlowFunc(ctx, tc.f, nil)
		if output != tc.output {
			t.Errorf("%s: get %d, expect %d", tc.spec, output, tc.output)
		}
	}

}
