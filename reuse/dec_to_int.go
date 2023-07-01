package reuse

func DecToInt(s string) (int, int, bool) {
	
	n := 0
	i := 0
	for i = 0; i < len(s) && '0' <= s[i] && s[i] <= '9'; i++ {
		n = n*10 + int(s[i] - '0')
	}

	if i == 0 {
		return 0, 0, false
	} else {
		return n, i, true
	}
	
}