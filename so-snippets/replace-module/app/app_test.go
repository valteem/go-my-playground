package app

import (
	"testing"

	m "github.com/some-author/some-module"
)

func TestOutput(t *testing.T) {
	output := m.Output()
	if output != m.OutputString {
		t.Errorf("get %s, expect %s", output, m.OutputString)
	}
}
