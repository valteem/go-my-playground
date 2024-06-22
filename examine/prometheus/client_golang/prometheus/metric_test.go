package prometheus

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

func TestBuildFQName(t *testing.T) {
	tests := []struct{ namespace, subsystem, name, output string }{
		{"ns", "sub", "name", "ns_sub_name"},
		{"", "sub", "name", "sub_name"},
		{"ns", "", "name", "ns_name"},
		{"ns", "sub", "", ""},
		{"ns", "", "", ""},
		{"", "sub", "", ""},
		{"", "", "", ""},
	}
	for i, tc := range tests {
		if output := prometheus.BuildFQName(tc.namespace, tc.subsystem, tc.name); output != tc.output {
			t.Errorf("BuildFQName(), test case %d: get %s, expect %s", i, output, tc.output)
		}
	}
}
