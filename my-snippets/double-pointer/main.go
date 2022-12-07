package main

import "fmt"

type T struct {
	id int
	txt string
}

// *receiver is a pointer to struct variable, as well as *sender
// this is why after 'reassignment' both pointers point to the same memory address
func reassign(sender **T, receiver **T) {

	*receiver = *sender

}

// same as in 'reassignP', can remove one asterisk from all double pointers
func reassignX(sender **T, receiver **T) {

	**receiver = **sender

}

// *receiver here is value of struct variable that pointer points at
func reassignP(sender *T, receiver *T) {

	*receiver = *sender

}

// receiver here is just a copy of the receiver in caller function, it changes inside this function and does not go any futher outside
func reassignV(sender T, receiver T) { 

	receiver = sender

}

func printvaluesPointers(v1, v2 int, p1, p2 *int, pp1, pp2 **int) {
	fmt.Printf("%d %d %t %p %p %t %d %d %t %p %p %t\n", v1, v2, v1==v2, p1, p2, p1==p2, *p1, *p2, *p1==*p2, pp1, pp2, pp1==pp2)
}

func main() {

	t1 := &T{id:1, txt:"text1"}
	t2 := &T{}
	t3 := &T{}
	var t4 T
	t5 := &T{}

	reassign(&t1, &t2)
	fmt.Println(t1, t2) // t1 and t2 have same value 
	fmt.Printf("%p %p\n", t1, t2) // t1 and t2 point to the same memory address

	reassignX(&t1, &t5)
	fmt.Println(t1, t5) // t1 and t5 have same value 
	fmt.Printf("%p %p\n", t1, t5) // t1 and t5 point to different memory addresses

	reassignP(t1, t3)
	fmt.Println(t1, t3, *t1, *t3) // t1 and t3, *t1 and *t3 have same values
	fmt.Printf("%p %p\n", t1, t3) // t1 and t3 point to different memory addresses

	reassignV(*t1, t4)
	fmt.Println(t1, t4)
	fmt.Printf("%p %p\n", t1, &t4)

	t1.id = 11
	t1.txt = "text11"
	fmt.Println(t1, t2)

	v1, v2 := 1, 2
	p1, p2 := &v1, &v2
	pp1, pp2 := &p1, &p2
	**pp1 = **pp2
	printvaluesPointers(v1, v2, p1, p2, pp1, pp2)

	q1, q2 := 1001, 1002
	pq1, pq2 := &q1, &q2
	ppq1, ppq2 := &pq1, &pq2
	*ppq1 = *ppq2 //pq1 now points at the same address as pq2, but underlying q1 variable remains unchanged
	printvaluesPointers(q1, q2, pq1, pq2, ppq1, pp2)

	u1, u2 := 11, 12
	pu1, pu2 := &u1, &u2
	*pu1 = *pu2
	printvaluesPointers(u1, u2, pu1, pu2, &pu1, &pu2)

	w1, w2 := 101, 102
	pw1, pw2 := &w1, &w2
	pw1 = pw2
	printvaluesPointers(w1, w2, pw1, pw2, &pw1, &pw2)

}