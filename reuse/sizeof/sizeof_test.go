package sizeof

import (
	"math/bits"
	"testing"
	"unsafe"
)

func TestSizeofIntUint(t *testing.T) {

	var a int // size of int in bytes
	sizeOfInt := unsafe.Sizeof(a)

	var b uint // size of uint int bytes
	sizeOfUint := unsafe.Sizeof(b)

	if sizeOfInt != sizeOfUint {
		t.Fatalf("expect same size for int and uint, get %d for int and %d for uint", sizeOfInt, sizeOfUint)
	}

	var sizeBitsExpected int
	if sizeOfUint == 4 {
		sizeBitsExpected = 32
	} else if sizeOfUint == 8 {
		sizeBitsExpected = 64
	} else {
		t.Fatalf("expect size of uint 4 or 8 bytes, get %d", sizeOfInt)
	}

	sizeBitsActual := bits.UintSize
	if sizeBitsActual != sizeBitsExpected {
		t.Errorf("size of uint in bits: get %d, expect %d", sizeBitsActual, sizeBitsExpected)
	}

}
