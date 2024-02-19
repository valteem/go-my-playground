package reuse

import (
	"time"
)

// as opposed to Unix 

func ConvertUnixTimeZero() (time.Time, time.Time) {
	return time.Time{},     // std  zero time ((time.Time(0001-01-01T00:00:00Z))) https://pkg.go.dev/time#Time
	       time.Unix(0, 0)  // Unix zero time ((time.Time(1970-01-01T00:00:00Z))) https://pkg.go.dev/time#Time.Unix
}