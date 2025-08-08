package dto

import (
	"strings"

	"testing"

	//	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"github.com/prometheus/common/model"
)

func TestParse(t *testing.T) {

	input := `
#TYPE average_response_time gauge
#HELP average_response_time average response time in msec
average_response_time {request_path="some_path"} 110.0
average_response_time {request_path="some_other_path"} 220.0
`

	var parser expfmt.TextParser
	parsed, err := parser.TextToMetricFamilies(strings.NewReader(input))
	if err != nil {
		t.Fatalf("failed to parse text input: %v", err)
	}

	if actual, expected := len(parsed), 1; actual != expected {
		t.Errorf("expect %d metric families, get %d", expected, actual)
	}

	vec, err := expfmt.ExtractSamples(&expfmt.DecodeOptions{}, parsed["average_response_time"])
	if err != nil {
		t.Fatalf("failed to extract metric data: %v", err)
	}

	if len(vec) != 2 {
		t.Errorf("expect two samples (for some_path and some_other_path), get %d", len(vec))
	}

	// prometheus/common/model/value_type.go#34
	if actual, expected := vec.Type(), model.ValVector; actual != expected {
		t.Errorf("ValueType: get %v, expect %v", actual, expected)
	}

}
