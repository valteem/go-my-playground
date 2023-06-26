package reuse

import (
	"fmt"
	"strconv"
)

// https://github.com/recoilme/sniper/blob/master/chunk.go
func NextPowerOf2(v uint32) uint32 {
	v--
	present(v, 1)
	v |= v >> 1 // bitwise OR
	present(v, 2)
	v |= v >> 2
	present(v, 4)
	v |= v >> 4
	present(v, 8)
	v |= v >> 8
	present(v, 16)
	v |= v >> 16
	v++

	return v
}

func bitsPresent(i uint32) string {
	return (fmt.Sprintf("%03s",strconv.FormatUint(uint64(i),10)) + " bits are " + fmt.Sprintf("%010s",strconv.FormatUint(uint64(i), 2)))
}

func present(i uint32, j uint32) {
	fmt.Println(i, j, i>>j)
	fmt.Println(i, bitsPresent(i))
	fmt.Println(i>>j, bitsPresent(i>>j))
	fmt.Println(i|i>>j, bitsPresent(i|i>>j))

}