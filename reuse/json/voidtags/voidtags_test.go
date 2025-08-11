package voidtags

import (
	"bytes"
	"encoding/json"

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

	r := &rule{}
	decoder := json.NewDecoder(bytes.NewReader([]byte(input)))
	err := decoder.Decode(r)
	if err != nil {
		t.Errorf("failed to decode input: %v", err)
	}

}
