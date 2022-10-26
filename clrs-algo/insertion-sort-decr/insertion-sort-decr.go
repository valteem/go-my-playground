package isortdecr

import "log"

func InsertionSortDecrInt(slc []int) {

	length := len(slc)

	if length <= 1 {
		log.Println("Slice to sort must have length > 1")
		return
	}

	for j := 1; j <= length-1; j++ {
		key := slc[j]
		i := j - 1
		for 
		{
			if !((i > -1) && (slc[i] < key)) {break}
			slc[i+1] = slc[i]
			i--
		}
		slc[i+1] = key
	}
}