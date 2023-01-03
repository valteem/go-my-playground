package benchmarkexample

func LoopRange(input []int) int {

	i := 0

	for _, v := range input {
		i += v
	}

	return i

}

func LoopFor(input []int) int {

	i := 0

	for j := 0; j <len(input); j++ {
		i += j
	}

	return i

}