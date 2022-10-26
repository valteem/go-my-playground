package bufiom_test

import (
	"fmt"
	"testing"

	"lang.rev/stdlib-takeaways/bufiom"
)

func TestReaders(t *testing.T) {

	brp := bufiom.NewBufReader(4)
	brv := bufiom.NewBufReaderVal(4)
    var br bufiom.SomeReader = bufiom.BufReader{Buf: []byte("text"), Size: 4}
	var or bufiom.SomeReader = bufiom.OtherReader{Buf: []byte("text"), Size: 4}   
	orp := bufiom.NewOtherReader(4)
	orv := bufiom.NewOtherReaderVal(4)

	fmt.Println("-------- BufReader --------")
	bufiom.AssertReaderType("Defined as iface variable", br)
	bufiom.AssertReaderType("Defined as pointer", brp)
	bufiom.AssertReaderType("Defined as pointer, dereferenced", *brp)
	bufiom.AssertReaderType("Defined as value", brv)
	brv.Overwrite("hello")
	bufiom.AssertReaderType("Defined as value, after overwrite", brv)
	fmt.Println("------- OtherReader -------")
	bufiom.AssertReaderType("Defined as iface variable", or)
	bufiom.AssertReaderType("Defined as pointer", orp)
	bufiom.AssertReaderType("Defined as pointer, dereferenced", *orp)
	bufiom.AssertReaderType("Defined as value", orv)
	orv.Overwrite("hello")
	bufiom.AssertReaderType("Defined as value, after overwrite", orv)
}