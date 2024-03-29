package utils

import (
	"time"
)

type Comparator[T any] func(x, y T) int // x > y: 1, x < y: -1, x = y: 0

func TimeComparator(x, y time.Time) int {
	switch {
	case x.After(y):
		return 1 // x > y
	case y.After(x):
		return -1 // x < y
	default:
		return 0
	}
}
