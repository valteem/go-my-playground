package main

import "fmt"

func fctrl (num int) int {
	if num == 0 {
		return	1 } else {
		return num * fctrl(num - 1)
	}
}

func main() {
	fmt.Println(fctrl(5))
}