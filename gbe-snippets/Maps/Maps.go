package main

import "fmt"

func main() {

	mapStrInt := make(map[string]int)
	mapStrInt["key1"] = 11
	mapStrInt["key2"] = 21
	fmt.Println("Map of string (key) and int (value): ", mapStrInt)
	fmt.Println("has length ",len(mapStrInt))

	value1 := mapStrInt["key1"]
	fmt.Println("Value for a key in map ", value1)

	delete(mapStrInt,"key2")
	_,sts := mapStrInt["key2"]
	fmt.Println("Status of deleted key-value pair ", sts)

// Explicit map declaration with values
   mapIntStr := map[int]string{1:"val1",2:"val2"}
   fmt.Println("Explicitly declared map  of int (key) and string (value) ",mapIntStr)	

}