package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name string) *person {
	var p person
	p.name = name
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{"Bob",30})

	fmt.Println(person{name:"Jon",age:31})

	fmt.Println(person{name:"Fred"})

	fmt.Println(&person{name:"Ann", age:33})

	fmt.Println(newPerson("Jack"))

	s1 := person{"Jim",50}
	fmt.Println(s1.name, s1.age)
	s1.age = 51
	fmt.Println(s1.name,s1.age)
	sptr := &s1
	sptr.age = 52
	fmt.Println(s1.name,s1.age)
}