package env

import (
	"os"
	"testing"
)

func TestEnvData(t *testing.T) {

	if len(pairs) == 0 {
		t.Fatalf("failed to fetch any env data")
	}

	envsExpected := []struct {
		name  string
		value string
	}{
		{"ENV_URL1", "http://localhost:3001/segment1/segment2/segment3"},
		{"ENV_URL2", "https://example.com/pathvar1/value1/pathvar2/value2"},
	}

	for _, env := range envsExpected {
		if actual, expected := os.Getenv(env.name), env.value; actual != expected {
			t.Errorf("%q:\nget\n%q\nexpect\n%q\n", env.name, actual, expected)
		}
	}

}

// Nested package init() not invoked before running top level package functions
func TestNestedPackageInit(t *testing.T) {

	envsExpected := []struct {
		name  string
		value string
	}{
		// {"ENV_NESTED1", "ENV_NESTED1_VALUE"},
		// {"ENV_NESTED2", "ENV_NESTED2_VALUE"},
		{"ENV_NESTED1", ""},
		{"ENV_NESTED2", ""},
	}

	for _, env := range envsExpected {
		if actual, expected := os.Getenv(env.name), env.value; actual != expected {
			t.Errorf("%q:\nget\n%q\nexpect\n%q\n", env.name, actual, expected)
		}
	}
}
