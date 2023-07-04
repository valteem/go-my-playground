package reuse_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/valteem/reuse"
)

func TestAssertType(t *testing.T) {

	p := reuse.Person{Name: "name", Age: 30}
	s := reuse.StockItem{ID: 1, Description: "description", SupplierID: 1}
	f := struct{
		        input string
			    output string
			}{input: "input",
			  output: "output",
			}

	fmt.Println(reuse.AssertType(p))
	fmt.Println(reuse.AssertType(s))
	fmt.Println(reuse.AssertType(f))

	fmt.Println(reflect.TypeOf(f))

}