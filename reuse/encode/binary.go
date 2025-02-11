package encode

// Set all bits to the right from most significant non-zero bit to 1
func AllSignificantBitsTo1(x uint64) uint64 {

	numOfBits := 0
	bx := x
	for bx > 0 {
		bx /= 2
		numOfBits++
	}

	var output uint64 = 0
	var multiplier uint64 = 1
	for i := 0; i < numOfBits; i++ {
		output += multiplier
		multiplier *= 2
	}

	return output
}
