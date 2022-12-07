package main

import "fmt"

type T struct {
	id int
	txt string
}

func reassign(sender **T, receiver **T) {

	*receiver = *sender

}

func reassignP(sender *T, receiver *T) {

	*receiver = *sender

}

func reassignV(sender T, receiver T) {

	receiver = sender

}

func main() {

	t1 := &T{id:1, txt:"text1"}
	t2 := &T{}
	t3 := &T{}
	var t4 T

	reassign(&t1, &t2)
	fmt.Println(t1, t2) // t1 and t2 have same value 
	fmt.Printf("%p %p\n", t1, t2) // t1 and t2 point to the same memory address

	reassignP(t1, t3)
	fmt.Println(t1, t3, *t1, *t3) // t1 and t3, *t1 and *t3 have same values
	fmt.Printf("%p %p\n", t1, t3) // t1 and t3 point to different memory addresses

	reassignV(*t1, t4)
	fmt.Println(t1, t4)
	fmt.Printf("%p %p\n", t1, &t4)

}