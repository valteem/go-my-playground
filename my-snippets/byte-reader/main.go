package main

import (
	"bytes"
	"fmt"
)

func main() {

	b := []byte("abc")
	rec := make([]byte, len(b)) // copy(dst, src) function copies min(len(dst), len(src)) bytes
	br := bytes.NewReader(b)
	fmt.Println(br)
	n, err := br.Read(rec)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n, br, rec)
	
}