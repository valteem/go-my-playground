package reuse

func IsNum(input rune) bool {
	if input >= 48 && input <= 57 {
		return true
	} else {
		return false
	}
}
