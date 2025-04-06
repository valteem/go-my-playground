// Extract part of JSON using json.RawMessage
//https://stackoverflow.com/a/79557507

package json

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestRawMessage(t *testing.T) {

	input := `{"landplot":{"location":"somewhere","area":42.0, "terrain":"other"}, "price":"100000","seller":"someone"}`

	var rawdata RawData

	err := json.NewDecoder(strings.NewReader(input)).Decode(&rawdata)
	if err != nil {
		t.Fatalf("failed to decode input: %v", err)
	}

	var lpGet LandPlot
	lpRawData := rawdata.LP

	err = json.NewDecoder(strings.NewReader(string(lpRawData))).Decode(&lpGet)
	if err != nil {
		t.Fatalf("failed to decode raw data: %v", err)
	}

	lpWant := LandPlot{Location: "somewhere", Area: 42.0, Terrain: Other}

	if !reflect.DeepEqual(lpGet, lpWant) {
		t.Errorf("decoding raw data:\nget\n%v\nwant\n%v\n", lpGet, lpWant)
	}

}
