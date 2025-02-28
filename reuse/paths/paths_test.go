package paths

import (
	"path/filepath"

	"testing"
)

func TestFromSlash(t *testing.T) {

	tests := []struct {
		input []string
	}{
		{[]string{"testdata", "basic.rules"}},
		{[]string{"..", "..", "testdata", "basic.rules"}},
	}

	for _, tc := range tests {

		path := filepath.Join(tc.input...)
		pathFromSlash := filepath.FromSlash(path)
		if path != pathFromSlash {
			t.Errorf("converting %v to filepath:\nJoin() returns\n%s\nFromSlash() returns\n%s\n",
				tc.input,
				path,
				pathFromSlash,
			)
		}
	}

}
