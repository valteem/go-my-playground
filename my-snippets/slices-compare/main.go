package main

import (
	"log"
	"reflect"
)

func main() {

	s1 := []int{3, 7, 11}
	s2 := []int{3, 7, 11}
	log.Println("DeepEqual:", reflect.DeepEqual(s1,s2))
//	log.Println("Equal:", (s1 == s2)) // does not work, says 'slices can be compared only to nil'
}