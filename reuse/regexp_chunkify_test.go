package reuse

import (
	"reflect"
	"regexp"
	"testing"
)

// https://github.com/facette/natsort/blob/2cd4dd1e2dcba4d85d6d3ead4adf4cfd2b70caf2/natsort.go#L24
// \d+ - any (non-zero) number of digits
// \D+ - any (non-zero) number of non-digits
// |   - logical OR
var chunkifyRegexp = regexp.MustCompile(`(\d+|\D+)`)

func chunkify(s string) []string {
	return chunkifyRegexp.FindAllString(s, -1)
}

func TestChunkify(t *testing.T) {

	tests := []struct {
		input  string
		output []string
	}{
		{"abc123xyz", []string{"abc", "123", "xyz"}},
		{"123abc456", []string{"123", "abc", "456"}},
	}

	for _, tc := range tests {
		if output := chunkify(tc.input); !reflect.DeepEqual(output, tc.output) {
			t.Errorf("chunkify string %s:\nget\n%v\nexpect\n%v", tc.input, output, tc.output)
		}
	}

}
