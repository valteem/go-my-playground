package structpointer_test

import (
	. "lang.rev/stdlib-takeaways/structpointer"
	"fmt"
	"testing"
)

func TestBufStr(t *testing.T) {

	bv := NewBufStrVal([]string{"words", "count", "nothing"})
	fmt.Printf("Newly created struct variable has address %p\n", &bv)
	bv.Reset([]string{"nothing", "else", "matters"})
	fmt.Printf("After reset by pointer address is %p\n", &bv )
	bv.ResetByVal([]string{"Oh", "Balladonna"}) // does not work
	fmt.Println(bv)

	bp := NewBufStrPtr([]string{"words", "count", "nothing"})
	fmt.Println(bp)
	bp.Reset([]string{"nothing", "else", "matters"})
	fmt.Println(bp)
	bp.ResetByVal([]string{"Oh", "Balladonna"}) // does not work
	fmt.Println(bp)
}