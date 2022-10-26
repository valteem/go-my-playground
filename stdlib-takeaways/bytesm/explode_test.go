package bytesm_test

import (
	. "lang.rev/stdlib-takeaways/bytesm" // dot allows for calling functions without package qualifier
	"fmt"
	"testing"
	"unicode/utf8"
)


func TestExplode(t *testing.T) {

	s := "s\n\bometext"
	b := []byte(s)
//  bSl := explode.Explode(b, 5) // package qualifier is needed if we don't add dot to import package statement
	bsl := Explode(b, 5)         // after having added dot before package name in import statement we dont't need package qualifier
	fmt.Println(bsl)

	s = "\n"
	r, size := utf8.DecodeRune([]byte(s)) // rune and 1 (one rune read, probably)
	fmt.Println(r, size)

	s = "slice"
	bt := []byte(s)
	fmt.Println(bt)
	rec := make([]byte, 5)
	fmt.Println(rec)
	rec = bt[0:3:5]
	fmt.Println(rec)

}