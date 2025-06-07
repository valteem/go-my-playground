package bubble

// This looks rather like insertion sort
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

// Bubble sort: len(input) full-slice-length passes
func Bubble(input []int) {

	if len(input) < 2 {
		return
	}

	var j int
	for range len(input) - 1 {
		j = 0
		for j < len(input)-1 {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
			j++
		}
	}

}
