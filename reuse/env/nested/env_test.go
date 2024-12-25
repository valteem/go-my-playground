package nested

import (
	"os"
	"testing"
)

// Top level package init() not invoked before running nested package functions
func TestTopLevelInit(t *testing.T) {

	envsExpected := []struct {
		name  string
		value string
	}{
		// {"ENV_URL1", "http://localhost:3001/segment1/segment2/segment3"},
		// {"ENV_URL2", "https://example.com/pathvar1/value1/pathvar2/value2"},
		{"ENV_URL1", ""},
		{"ENV_URL2", ""},
	}

	for _, env := range envsExpected {
		if actual, expected := os.Getenv(env.name), env.value; actual != expected {
			t.Errorf("%q:\nget\n%q\nexpect\n%q\n", env.name, actual, expected)
		}
	}

}
