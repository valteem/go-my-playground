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
	/*
		https://pkg.go.dev/encoding/json#Unmarshal
		To unmarshal JSON into an interface value, Unmarshal stores one of these in the interface value:
			bool,                   for JSON booleans
			float64,                for JSON numbers
			string,                 for JSON strings
			[]interface{},          for JSON arrays
			map[string]interface{}, for JSON objects
			nil,                    for JSON null
	*/
	if int(p["age"].(float64)) != 42 || p["name"] != "somename" {
		t.Errorf("get %v, expect {\"somename\", 42}", p)
	}

}
