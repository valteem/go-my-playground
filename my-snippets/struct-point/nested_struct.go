package main

import "fmt"

type person struct {
	name string
	age int
}

type requester struct {
	pers person
	email string
}

type requesterSmart struct {
	pers *person
	email string
}

func main() {

	p := person{name:"Vasya", age: 30}

	rv := requester{pers:p, email:"a1@yahoo.com"} // nested structure can be assigned by value, not pointer
	fmt.Println("Nested structure assigned by value", rv)

	rp := requesterSmart{pers: &p, email: "a1@yahoo.com"}
	fmt.Println("Nested structure assigned by referencw", rp)
	fmt.Println("Components(fields) of the nested structure assigned by reference",rp.pers, rp.email)

	p.name = "Petya"
	p.age = 31
	fmt.Println(rv) // changes in nested structure included by value are not reflected by outer structure
	fmt.Println(rv.pers) // this is technically not p, it is another person object
	fmt.Printf("%p %p %p %p\n", &p, &rv.pers, rp.pers, &rp.pers) // this shows different memory addresses for p and rv.pers, and same for p and rp.person
	fmt.Println(*rp.pers, rp.pers, &rp, &rp.pers) // changes is nested structure are taken by outer structure only if nested structure is included as pointer
}