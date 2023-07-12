package reuse_test

import (
	"fmt"
	"testing"

	"github.com/valteem/reuse"
)

func TestBinaryStringToByteSlice(t *testing.T) {
	b := make([]byte, 2)
	s := "1010101010101010"
	reuse.BinaryStringToByteSlice(s, b)
	fmt.Println(b)

}

func TestAnd(t *testing.T) {
	b1 := make([]byte, 2)
	s1 := "1010101010101010"
	reuse.BinaryStringToByteSlice(s1, b1)
	fmt.Println(b1)
	b2 := make([]byte, 2)
	s2 := "1111111100000000"
	reuse.BinaryStringToByteSlice(s2, b2)
	fmt.Println(b2)
	b3 := make([]byte, 2)
	reuse.And(b1, b2, b3)
	fmt.Println(b3)
}

func TestAndNot(t *testing.T) {
	b1 := make([]byte, 2)
	s1 := "1010101010101010"
	reuse.BinaryStringToByteSlice(s1, b1)
	fmt.Println(b1)
	b2 := make([]byte, 2)
	s2 := "0000000011111111"
	reuse.BinaryStringToByteSlice(s2, b2)
	fmt.Println(b2)
	b3 := make([]byte, 2)
	reuse.AndNot(b1, b2, b3)
	fmt.Println(b3)
}