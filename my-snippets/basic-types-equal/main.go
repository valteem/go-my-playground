package main

import "log"

type Person struct {
	Name string
	Age int
}

func main() {
	i1 := 7
	i2 := 7
	log.Println(&i1, &i2, i1==i2)
	s1 := "txt"
	s2 := "txt"
	log.Println(&s1, &s2, s1==s2)

	p1 := Person{Name:"Vasya", Age:30}
	p2 := Person{Name:"Vasya", Age:30}
	log.Println(p1==p2)
}