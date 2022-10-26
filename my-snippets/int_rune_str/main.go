package main

import "fmt"

func main() {

	var i int32 = 0

	fmt.Printf("Start reading integers ... \n")

	for i >= 0 {
		fmt.Printf("Input integer (negative to quit)\n")
		fmt.Scan(&i)
		if i < 0 {
			return
		}
		fmt.Println("i = ", i, "rune(i) = ", rune(i), " string[i] = ", string(i))
	}

}
