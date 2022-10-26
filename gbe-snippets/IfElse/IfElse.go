package main

import (
	"fmt"
	"strconv"
)

func main() {

	var str1 string
	var str2 string 
	fmt.Println("Enter dividend")
	fmt.Scanf("%s", &str1)
	fmt.Printf("%s\n", "Dividend" + str1)
	fmt.Println("Enter divider")
	fmt.Scanf("%s", &str2)
    
	var p1 int64
	var p2 int64
	p1, _ = strconv.ParseInt(str1, 10, 32)
	fmt.Println("Dividend converted to int is " + strconv.FormatInt(p1,10))
	p2, _ = strconv.ParseInt(str2, 10, 32)
	fmt.Println("Divider converted to int is  " + strconv.FormatInt(p2,10))

	if p1%p2 == 0 {
		fmt.Println("Fully divisible")
	} else {
		fmt.Println("Not fully divisible")
	}
}