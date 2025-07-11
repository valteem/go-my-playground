package pathjoin

import (
	"path"

	"testing"
)

func TestPathJoin(t *testing.T) {

	tests := []struct {
		input  []string
		output string
	}{
		{
			[]string{"some", "path", "to", "some", "folder"},
			"some/path/to/some/folder",
		},
		{
			[]string{"some", "path", "to", "some", "folder", ".."},
			"some/path/to/some",
		},
		{
			[]string{"some", "path", "to", "some", "folder", "..", ".."},
			"some/path/to",
		},
	}

	for _, tc := range tests {
		if output := path.Join(tc.input...); output != tc.output {
			t.Errorf("get\n%s\nexpect\n%s\n", output, tc.output)
		}
	}

}
