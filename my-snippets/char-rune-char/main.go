package main

import "fmt"

func main() {

	var s string = "s"

	for s != "stop" {
		fmt.Printf("Input a single character (or 'stop' to quit\n")
		fmt.Scan(&s)
		if s == "stop" {
			return
		}
		b := []byte(s)
		fmt.Println("s = ", s, "rune(s) = ", rune(s[0]), "string(rune(s)) = ", string(rune(s[0])), " byte(s) = ", b)
	}

}
