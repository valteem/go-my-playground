package reuse

func SliceEqual(a []int, b []int) bool {
	
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true

}

func SliceContains(s []int, ss []int) int {
	
	if len(ss) > len(s) {
		return 0
	}

	count := 0

	for i := 0; i <= len(s) - len(ss); i++ {
		if SliceEqual(s[i : i + len(ss)], ss) {
			count ++
		} 
	}

	return count
}

