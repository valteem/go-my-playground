package utils

import (
	"testing"
	"time"
)

func TestTimeComparator(t *testing.T) {
	now := time.Now()
	tests := []struct {
		description string
		x           time.Time
		y           time.Time
		output      int
	}{
		{
			description: "x > y, returns +1",
			x:           now,
			y:           now.Add(-60 * time.Second),
			output:      1,
		},
		{
			description: "x < y, returns -1",
			x:           now,
			y:           now.Add(60 * time.Second),
			output:      -1,
		},
		{
			description: "x = y, returns 0",
			x:           now,
			y:           now,
			output:      0,
		},
	}
	for _, tst := range tests {
		if r := TimeComparator(tst.x, tst.y); r != tst.output {
			t.Errorf("%s: get %d, expect %d", tst.description, r, tst.output)
		}
	}
}
