package main

import "fmt"

func main() {
	var intSl = make([]int,5)
	fmt.Println("Empty Slice: ", intSl )
	intSl[0] = 1
	intSl[1] = 11
	intSl[2] = 21
	fmt.Println("Half empty slice ",intSl)
	fmt.Println("of length ",len(intSl))
	intSl = append(intSl,71)
	fmt.Println("Slice with appended element ",intSl)
	fmt.Println("and new length ",len(intSl))

	var strSl = make([]string,5)
	fmt.Println("Empty slice of string ",strSl)
	strSl[0] = "sl0"
	strSl[1] = "sl1"
	strSl[2] = "sl2"
	fmt.Println("Half empty slice of string ",strSl)
	strSl = append(strSl, "sl6")
	fmt.Println("Slice of strings with appended element: ", strSl)

	strSlC := make([]string, 3)
	copy(strSlC,strSl)
	fmt.Println("Copied (partially) slice of string ", strSlC)

	fmt.Println("Sliced: ", strSl[:2])

	twoDimSl := make([][]int, 3)
	for i:=0;i<3;i++ {
		innerLen := i+1
		twoDimSl[i] = make([]int, innerLen)
		for j:=0;j<innerLen;j++ {
			twoDimSl[i][j] = i+j
		}
	}
	fmt.Println("MD array with varying slice length: ",twoDimSl)

}