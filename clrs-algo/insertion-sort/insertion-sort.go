package insertionsort

import "log"

func InsertionSortInt(slc []int) {

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
			if !((i > -1) && (slc[i] > key)) {break}
			slc[i+1] = slc[i]
			i--
		}
		slc[i+1] = key
	}
}

// https://github.com/shady831213/algorithms/blob/master/sort/insertionSort.go
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- { // not optimal, will iterate over all j values with no break at arr[j-1] < arr[j]
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}