package isortdecr_test

import (
	isort "clrs.algo/insertion-sort-decr"
	"log"
	"testing"
)

func TestIsortDecr(t *testing.T) {
	intSlice := []int{11, 4, 18, 15}
	log.Println("Slice to sort", intSlice)
	isort.InsertionSortDecrInt(intSlice)
	log.Println("Sorted slice", intSlice)
}