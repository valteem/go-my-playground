package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	fcopy := func(srcStr string, srcNumSymCopy int, dst []byte) (int, error) {

		if len(srcStr) < srcNumSymCopy {
			errSrcNum := errors.New("number of copied symbols exceeds length of source string")
			log.Println(errSrcNum)
			return 0, errSrcNum
		}

		if srcNumSymCopy > len(dst) {
			errSrcLen := errors.New("source cannot be longer than destination")
			log.Println(errSrcLen)
			return 0, errSrcLen
		}

		return copy(dst, srcStr[:srcNumSymCopy]), nil
	}

	b1 := []byte("1234567890")
	n, e := fcopy("new", 2, b1)
	fmt.Println(n, e, string(b1))
	n, e = fcopy("more", 4, b1)
	fmt.Println(n, e, string(b1))
	n, e = fcopy("moremore", 10, b1)
	fmt.Println(n, e, string(b1))
	n, e = fcopy("moremoremore", 12, b1)
	fmt.Println(n, e, string(b1))	
	
}