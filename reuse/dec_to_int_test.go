package reuse_test

import (
	"fmt"
	"testing"

	"github.com/valteem/reuse"
)

func TestDecToInt(t *testing.T) {

	s := `abc01589.:`
	for _, b := range s {
		fmt.Println(b, '0' <= b, b <= '9', int(b), int(b - '0')) // int(b) returns ASCII code, int(b - '0') returns int number for integers
	}

	fmt.Println(reuse.DecToInt(`a001`));
	fmt.Println(reuse.DecToInt(`0001`));
	fmt.Println(reuse.DecToInt(`001a`));
	
}