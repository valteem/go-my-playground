package caller

import (
	"encoding/json"
	"os"

	"testing"

	"gotest.tools/v3/env"
)

type config struct {
	URL string `json:"url"`
}

func TestChangeWorkDir(t *testing.T) {

	tests := []struct {
		workDirRelPath string
		url            string
	}{
		{"../assets/apples", "apples"},
		{"../assets/cherries", "cherries"},
	}

	for _, tc := range tests {

		reset := env.ChangeWorkingDir(t, tc.workDirRelPath)

		b, err := os.ReadFile("config.json")
		if err != nil {
			t.Errorf("failed to read file at %q: %v", tc.workDirRelPath, err)
			reset()
			continue
		}

		c := &config{}
		err = json.Unmarshal(b, c)
		if err != nil {
			t.Errorf("failed to unmarshal data at %q: %v", tc.workDirRelPath, err)
			reset()
			continue
		}

		if actual, expected := c.URL, tc.url; actual != expected {
			t.Errorf("config URL: get %q, expect %q", actual, expected)
		}

		reset()

	}
}
