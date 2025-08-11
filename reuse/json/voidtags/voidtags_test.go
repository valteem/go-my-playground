package voidtags

import (
	"bytes"
	"encoding/json"
	"reflect"

	"testing"
)

func TestDecodeVoidTags(t *testing.T) {

	input := `{
		"name": "authz",
		"source":{
			"principals": ["source1", "source2", "source3"] 
		},
		"request":{
			"paths":["some/path","some/other/path"],
			"headers":[
				{
					"key":"Allow-Origin",
					"values":["some-origin","some-other-origin"]
				},
				{
					"key":"Custom-Header",
					"values":["some-custom-header-value","some-other-custom-header-value"]
				}
			]
		}	
	}
`

	expected := NewRule(
		RuleName("authz"),
		RuleSource(
			Peer{
				Principals: []string{"source1", "source2", "source3"},
			},
		),
		RuleRequest(
			Request{
				Paths: []string{"some/path", "some/other/path"},
				Headers: []Header{
					{"Allow-Origin", []string{"some-origin", "some-other-origin"}},
					{"Custom-Header", []string{"some-custom-header-value", "some-other-custom-header-value"}},
				},
			},
		),
	)

	actual := Rule{}
	decoder := json.NewDecoder(bytes.NewReader([]byte(input)))
	err := decoder.Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode input: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("get\n%v\nexpect\n%v\n", actual, expected)
	}

}
