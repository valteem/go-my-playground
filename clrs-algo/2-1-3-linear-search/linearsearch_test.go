//2.1-3 Linear Search

package linearsearch_test

import (
	ls "clrs.algo/2-1-3-linear-search"
	"log"
	"testing"
	)

func TestLinearSearch(t *testing.T) {

	s := []int{1,7,11,18,27,31}
	log.Println(ls.LinearSearch(s,11))
	log.Println(ls.LinearSearch(s,99))
}