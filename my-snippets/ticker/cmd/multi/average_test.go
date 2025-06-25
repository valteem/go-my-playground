package main

import (
	"math"
	"testing"
)

func TestRecurringAverage(t *testing.T) {

	tests := []struct {
		rng    []float64
		output float64
	}{
		{
			rng:    []float64{1., 1., 1.},
			output: 1.,
		},
		{
			rng:    []float64{1., 2., 4., 5.},
			output: 3.,
		},
	}

	var kfloat float64
	equalThreshhold := 1e-3
	for _, tc := range tests {

		av := 0.

		for k := 1; k <= len(tc.rng); k++ {
			kfloat = float64(k)
			av = ((kfloat-1)*av + tc.rng[k-1]) / kfloat
		}

		if variance := math.Abs(av - tc.output); variance > equalThreshhold {
			t.Errorf("output variance is %f for %v", variance, tc.rng)
		}

	}

}
