package reuse_test

import (
	"testing"
)

func changeNoEffect(value int) {
	value += 1
}

// change argument value without passing pointer as argument
func changeEffect(value int) int {
	value += 1
	return value
}

func TestReturnArg(t *testing.T) {

	v1 := 1
	changeNoEffect(v1)
	if v1 != 1 {
		t.Errorf("changeNoEffect(): get %d, expect 1", v1)
	}

	v2 := 1
	v2 = changeEffect(v2)
	if v2 != 2 {
		t.Errorf("changeEffect(): get %d, expect 2", v2)
	}

}
