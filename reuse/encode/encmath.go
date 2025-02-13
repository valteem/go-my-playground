package encode

import "math"

func EncodeFloat64Bits(x float64) uint64 {
	return math.Float64bits(x)
}
