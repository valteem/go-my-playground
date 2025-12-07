package metric

import (
	"bytes"

	"testing"

	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
)

func TestNewEncoder(t *testing.T) {

	tests := []struct {
		input  string
		format expfmt.Format
		output expfmt.Encoder
	}{
		{
			"{some fancy metrics}",
			expfmt.NewFormat(expfmt.TypeProtoText),
			nil,
		},
		{
			"some fancy metrics",
			expfmt.NewFormat(expfmt.TypeTextPlain),
			nil,
		},
	}

	for _, tc := range tests {
		var out *bytes.Buffer
		encoder := expfmt.NewEncoder(out, expfmt.Format(tc.format))
		mf := &dto.MetricFamily{}
		encoder.Encode(mf)
	}
}
