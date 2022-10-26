package insertionsort

import (
//	isort "clrs.algo/insertion-sort"
	"flag"
	"log"
	"strconv"
	"testing"
)

// https://stackoverflow.com/a/62510653
func getSliceInt(s *[]string) []int {
	slcStr := *s
	ret := []int{}
	for _, str := range slcStr {
		elt_int, _ := strconv.Atoi(str)
		ret = append(ret, elt_int)
	}
	return ret
}

func TestIsort(t *testing.T) {
 	flag.Parse()
	strSlice := flag.Args()
	intSlice := getSliceInt(&strSlice)
	intSlice = nil
	intSlice = append(intSlice, 11, 4, 18, 15)
	log.Println("Slice to sort", intSlice)
	InsertionSortInt(intSlice)
	log.Println("Sorted slice", intSlice)

	intSliceGH := []int{11, 4, 18, 15}
	insertionSort(intSliceGH)
}