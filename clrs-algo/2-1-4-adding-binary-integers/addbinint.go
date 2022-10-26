// 2.1-4 Adding Binary Integers

package addbinint

func Max(x int, y int) int {
	if x > y {return x}
	return y
}

// slice.Contains as of Go 1.18
func Contains(lv int, s []int) bool {
	for _, v := range s {
		if lv == v {return true}
	}
	return false
}

func AddBinaryIntegers (b1 []int, b2[]int) []int {
	var sumcurr int
	var sumprev int = 0
	bits := []int{0,1}
	l1 := len(b1)
	l2 := len(b2)
	lmax := Max(l1, l2)
	slc := make([]int, lmax + 1)
	if l1 != l2 {
		// Adding leading zeroes
		for i := 0; i < lmax; i++ {
			if i < (lmax - l1) {
				b1 = append([]int{0}, b1...)
			}
			if i < (lmax - l2) {
				b2 = append([]int{0}, b2...)
			}
		}
	}

	for j := (lmax - 1); j >= 0; j-- {
		sumcurr = b1[j] + b2[j] + sumprev
		if Contains(sumcurr, bits) {
			slc[j] = sumcurr
			sumprev = 0
			} else {
				slc[j] = 0
				sumprev = 1
			}
	}
	slc[0] = sumprev
	return slc
}