// Type assertion with dynamic type definition:
// i, ok := v.(interface{/* type signature here */})

package reuse_test

import (
	"testing"
)

type testCounter struct {
	count int
}

func (tc testCounter) testCount() int {
	return tc.count
}

type testGauge struct {
	gauge int
}

func (tg testGauge) testGauge() int {
	return tg.gauge
}

func getCount(v any) int {
	if c, ok := v.(interface{ testCount() int }); ok {
		return c.testCount()
	}
	return 0
}

func getGauge(v any) int {
	if g, ok := v.(interface{ testGauge() int }); ok {
		return g.testGauge()
	}
	return 0
}

func TestAssertMethodImplementation(t *testing.T) {

	tests := []struct {
		name        string
		input       any
		outputCount int
		outputGauge int
	}{
		{
			name:        "testCount(), no testGauge()",
			input:       testCounter{count: 5},
			outputCount: 5,
			outputGauge: 0,
		},
		{
			name:        "testGauge(), no testCount()",
			input:       testGauge{gauge: 5},
			outputCount: 0,
			outputGauge: 5,
		},
		{
			name:        "no testCount(), no testGauge()",
			input:       "just a string",
			outputCount: 0,
			outputGauge: 0,
		},
	}

	for _, tc := range tests {
		if count, gauge := getCount(tc.input), getGauge(tc.input); count != tc.outputCount || gauge != tc.outputGauge {
			t.Errorf("%s: expect %d for count, %d for gauge, get %d for count, %d for gauge", tc.name, tc.outputCount, tc.outputGauge, count, gauge)
		}
	}

}
