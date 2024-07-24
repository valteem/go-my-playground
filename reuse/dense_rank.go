package reuse

func ClimbingLeaderboard(ranked []int32, player []int32) []int32 {
	rankedList := []int32{} //
	rankedList = append(rankedList, ranked[0])
	last := ranked[0]
	for _, r := range ranked[1:] {
		if r < last {
			rankedList = append(rankedList, r)
			last = r
		}
	}

	result := []int32{}
	for _, p := range player {
		if p > rankedList[0] {
			result = append(result, 1)
			continue
		}
		if p < rankedList[len(rankedList)-1] {
			result = append(result, int32(len(rankedList)+1))
			continue
		}
		top := 0
		bottom := len(rankedList) - 1
		for {
			pos := (top + bottom) / 2
			if rankedList[pos] == p {
				result = append(result, int32(pos+1))
				break
			} else if rankedList[pos] > p && rankedList[pos+1] <= p {
				result = append(result, int32(pos+2))
				break
			} else {
				if rankedList[pos] > p {
					top = pos
				} else {
					bottom = pos
				}
			}
		}
	}
	return result
}
