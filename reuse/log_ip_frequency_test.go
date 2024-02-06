package reuse_test

import (
	"testing"
	"github.com/valteem/reuse"
)

func TestLogIPFrequence(t *testing.T) {

	s := `[111.12.14.1] ping
	[999.12.12.1] more
	[111.12.14.1] another ping
	[10.10.10.255] attempt
	16.14.100.199`

	result := reuse.ParseLogForIP(s)

	expected := map[string]int{"10.10.10.255":1, "111.12.14.1":2, "16.14.100.199": 1}

	for k, v := range expected {
		if result[k] != v {
			t.Errorf("should be equal: result %+v, expected %+v", result[k], v)
		}
	}
}