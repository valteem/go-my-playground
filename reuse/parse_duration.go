package reuse

import (
	"fmt"
	"math"
	"time"
)

func ParseDuration(inp float64) (time.Duration, error) {
	ts := inp * float64(time.Nanosecond)

	if (ts > float64(math.MaxInt64)) || (ts < float64(math.MinInt64)) || (math.IsNaN(ts)) {
		return 0, fmt.Errorf("cannot parse %f to a valid duration, it overflows int64", inp)
	}

	return time.Duration(ts), nil
}
