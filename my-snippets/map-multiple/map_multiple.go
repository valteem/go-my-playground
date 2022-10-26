package main

import "fmt"

func main() {

	v := make(map[int][]string)

	s := []string{"value0", "value1"}
	v[0] = s

	s = []string{"value11", "value12", "value51"}
	v[1] = s

	s = []string{"value39"}
	v[2] = s

	fmt.Println(v)

	fmt.Println("Delete a key")
	delete(v, 2)
	fmt.Println(v)

	fmt.Println("Modify a key")
	s = []string{"value1", "value2"}
	v[0] = s
	fmt.Println(v)
	
}