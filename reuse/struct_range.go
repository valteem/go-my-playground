package reuse

import (
	"fmt"
)

func RangeOverStructDefinedInline() {

	for i, tt := range []struct{
		inpValue any
		expected []int
	}{{"input", []int{1 ,2, 3}}, {"output", []int{2, 4, 6}}} {
		fmt.Println(i, tt)
	}
}