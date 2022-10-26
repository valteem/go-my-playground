package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("Empty Array: ", a)
	a[4] = 99
	fmt.Println("Set: ",a)
	fmt.Println("Get: ",a[4])

	b := [4]int{11,21,31,41}
	fmt.Println("Init with values: ",b)

	var twoDim [3][4]int
	for i:=0;i<3;i=i+1 {
		for j:=0;j<4;j=j+1 {
			twoDim[i][j] = i + j
			fmt.Printf("2D array element row %d col %d %d\n",i,j,twoDim[i][j])
		}
	}
	fmt.Println("2D array: ",twoDim)

}