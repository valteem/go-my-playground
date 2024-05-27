package main

import (
	"fmt"
	"strings"
)

const (
	rmin, rmid, rmax = 1, 512 * 1024, 1024 * 1024
)

func String(nmin, nmax int) (string, string) {

	var a [2]string
	n1 := a[:0]
	n2 := make([]string, 0, nmax-nmin+1) // escapes to heap (only if capacity is explicitly set)
	for i := nmin; i <= nmax; i++ {
		v := fmt.Sprintf("%d", i)
		n1 = append(n1, v)
		n2 = append(n2, v)
	}
	s1 := strings.Join(n1, ",")
	s2 := strings.Join(n2, ",")

	return s1, s2

}

func main() {
	s1, s2 := String(rmin, rmax)
	s1a, s2a := String(rmin, rmid)
	s1b, s2b := String(rmid+1, rmax)
	s1x, s2x := s1a+","+s1b, s2a+","+s2b
	fmt.Println(s1 == s1x, s2 == s2x)
}
