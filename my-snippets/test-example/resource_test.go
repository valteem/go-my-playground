// https://go.dev/blog/examples

package testexample

import (
	"fmt"
)

func ExampleMultAndAddEQ() {
	fmt.Println(MultAndAdd(1, 2, 3) == 5)
	// Output: true
}

func ExampleMultAndAddGT() {
	fmt.Println(MultAndAdd(1, 2, 3) > 5)
	// Output: false
}

func ExampleMultAndAddValue() {
	fmt.Println(MultAndAdd(1, 2, 3))
	// Output: 5
}
