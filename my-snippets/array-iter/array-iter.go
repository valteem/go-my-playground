package main

import (
	"fmt"
	"strconv"
)

const (
	ROWNUM = 4
	COLNUM = 3
)

func main() {
	
	arr := make([][]string, ROWNUM)
	for row_idx := range arr {
		arr[row_idx] = make([]string, COLNUM) 
	}

	for row_idx, arrRow := range arr {
		for col_idx, _ := range arrRow {
			arr[row_idx][col_idx] = "r" + strconv.Itoa(row_idx ) + "c" + strconv.Itoa(col_idx)
		}
	}

	fmt.Println(arr)
	for row_idx, arrRow := range arr {
		for col_idx, _ := range arrRow {
			fmt.Printf("%s ", arr[row_idx][col_idx])
		}
		fmt.Printf("\n")
	}
}