package addbinint_test

import (
	abi "clrs.algo/2-1-4-adding-binary-integers"
	"log"
	"testing"
)

func TestAddBinInt(t *testing.T) {
	s1 := []int{1,1,1}
	s2 := []int{1}
	s3 := abi.AddBinaryIntegers(s1, s2)
	log.Println(s3)
}