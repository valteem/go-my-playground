package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	for j := 10; j <= 15; j++ {
		fmt.Println(j)
	}

	for n := 1; n <= 7; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}