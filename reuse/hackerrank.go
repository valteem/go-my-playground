package reuse

func CircularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	l := int32(len(a))
	if k > l {
		k = k % l
	}
	output := []int32{}
	for _, q := range queries {
		index := q - k
		if index < 0 {
			index += l
		}
		output = append(output, a[index])
	}
	return output
}

// n - number of chairs
// m - number of distributed goods
// s - number of chair that starts the distribution
func CircleDistribution(n int32, m int32, s int32) int32 {
	var seatIndex, shift int32

	if m > n {
		shift = m % n
	} else {
		shift = m
	}
	seatIndex = s + shift - 1
	if seatIndex > n {
		seatIndex = seatIndex % n
	}
	if seatIndex == 0 {
		seatIndex = n
	}

	return seatIndex
}

func ReverseDigits(input int) int {
	output := 0
	for input > 0 {
		output *= 10
		output += input % 10
		input /= 10
	}
	return output
}

func myMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func AppendAndDelete(s string, t string, k int32) string {
	rs, rt := []rune(s), []rune(t)
	index := 0
	for index < myMin(len(rs), len(rt)) && rs[index] == rt[index] {
		index++
	}
	count := len(rs) + len(rt) - 2*index
	if count <= int(k) && ((int(k)-count)%2 == 0 || len(rs)+len(rt) <= int(k)) {
		return "Yes"
	} else {
		return "No"
	}
}
