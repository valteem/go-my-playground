package replacenested

import (
	"testing"

	"github.com/some-prominent-account/some-popular-module" // silently imported as `nested`
)

func TestFoodName(t *testing.T) {

	tests := []struct {
		input  nested.FoodCode
		output string
	}{
		{nested.Apples, "apples"},
		{nested.Cherries, "cherries"},
		{nested.Onions, "onions"},
		{nested.FoodCode(42), "not found"},
	}

	for _, tc := range tests {
		if actual, expected := nested.FoodName(tc.input), tc.output; actual != expected {
			t.Errorf("get %q, expect %q", actual, expected)
		}
	}
}
