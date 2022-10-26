package main

import "fmt"

type mtype struct {i int; s string; next *mtype}

func (t *mtype) Initialize(i_init int, s_init string, m_init *mtype) {t.i = i_init; t.s = s_init; t.next = m_init}

func main() {

	var mt1, mt2 mtype
	mt1.Initialize(1,"text1", &mt2)
	mt2.Initialize(2,"text2", &mt1)
	fmt.Printf("%d %s %p %d %s %p\n", mt1.i, mt1.s, mt1.next, mt1.next.i, mt1.next.s, mt1.next.next)
	fmt.Printf("%d %s %p %p\n", mt2.i, mt2.s, mt2.next, &mt1)

}