package reuse

import "fmt"

// TODO: filling byte slice with bits of arbitrary size (not 8x), padding to 8, big/little endian
// Refer to roaring.Bitmap / gocroaring.Bitmap 

func BinaryStringToByteSlice(s string, b []byte) error {
	digitCount := 0
	byteCount := 0
	for i := 0; i < len(s); i++ {
		r := s[i]
		switch {
		case r == 0x31:
			b[byteCount] |= (1 << digitCount)
		case r == 0x30:
			b[byteCount] &^= (1 << digitCount)
		default:
			return fmt.Errorf("invalid input string (only 0 and 1 digit values are accepted")
		}
		digitCount++
		if digitCount > 7 {
			digitCount = 0;
			byteCount++
		}
	}
	return nil
}

func And(a, b, c []byte) error {
	if !(len(a) == len(b) || len(a) == len(c)) {
		return fmt.Errorf("slices must be of the same length")
	}
	if len(a) == 0 {
		return fmt.Errorf("empty slices not accepted")
	}
	for i := 0; i < len(a); i++ {
		c[i] = a[i] & b[i]
	}
	return nil
}

func AndNot(a, b, c []byte) error {
	if !(len(a) == len(b) || len(a) == len(c)) {
		return fmt.Errorf("slices must be of the same length")
	}
	if len(a) == 0 {
		return fmt.Errorf("empty slices not accepted")
	}
	for i := 0; i < len(a); i++ {
		c[i] = a[i] &^ b[i]
	}
	return nil
}