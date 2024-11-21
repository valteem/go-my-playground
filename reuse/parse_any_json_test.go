package reuse_test

import (
	"testing"

	"github.com/valteem/reuse"
)

func TestParseAnyJSON(t *testing.T) {

	s := `{
	"name": "somename",
	"age": 42
	}`

	output, err := reuse.ParseAnyJSON(s)
	if err != nil {
		t.Fatalf("failed to parse input: %v", err)
	}

	p := output.(map[string]any)
	if p["age"] != float64(42) || p["name"] != "somename" {
		t.Errorf("get %v, expect {\"somename\", 42}", p)
	}

}
