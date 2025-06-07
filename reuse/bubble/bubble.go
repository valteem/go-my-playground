package bubble

// In-place bubble sort
func Sort(input []int) {
	if len(input) < 2 {
		return
	}
	var v, w int
	for i := range len(input) - 1 {
		for j := i + 1; j < len(input); j++ {
			if v, w = input[i], input[j]; v > w {
				input[i], input[j] = w, v
			}
		}
	}
}
