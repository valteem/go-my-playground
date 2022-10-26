package main

import (
	"fmt"
	"strconv"
)

func Bits(i uint64) string {
	return strconv.FormatUint(i, 2)
}

func BitsPres(i uint64) string {
	return (fmt.Sprintf("%03s",strconv.FormatUint(i,10)) + " bits are " + fmt.Sprintf("%010s",strconv.FormatUint(i, 2)))
}

func main() {

	var i uint64 = 62
	var j uint64 = 39
	fmt.Println(BitsPres(i))
	fmt.Println(BitsPres(j))
	fmt.Println("AND (i & j)")
    fmt.Println(BitsPres(i & j))
	fmt.Println("OR (i | j)")
	fmt.Println(BitsPres(i | j))
    fmt.Println("XOR (i ^ j)")
	fmt.Println(BitsPres(i ^ j))
	fmt.Println("AND NOT (i &^ j)") // who knows what it is
	fmt.Println(BitsPres(i &^ j))
}