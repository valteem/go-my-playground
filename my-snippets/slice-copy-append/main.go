package main

import (
	"fmt"
)

func sliceCopyAppend(sr []byte, sc []byte, sa []byte) (int, int, []byte) {

	nc := copy(sr, sc)
	sn := append(sr, sa...)
	na := len(sa)

	return nc, na, sn

}

func main() {

	sr := make([]byte, 3)

	nc, na, sn := sliceCopyAppend(sr, []byte("abc"), []byte("xyz"))

	fmt.Println(nc, na, string(sr), string(sn))
}
