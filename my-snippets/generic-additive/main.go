package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func Scale[S ~[]E, E int](s S, sc E) S{
    r := make(S, len(s))
    for i, v := range s {
        r[i] = v * sc
    }

    return r
}

type additive interface {
	int|uint32|uint64|float32|float64
}

func add[T additive](a T, b T) T {
	return a + b
}

func addConstraints[T constraints.Ordered](a T, b T) T {
	return a + b
}

func main() {

	a_int, b_int := 1, 2
	fmt.Println(add(a_int, b_int), addConstraints(a_int, b_int))
	a_uint32, b_uint32 := 1, 2
	fmt.Println(add(a_uint32, b_uint32), addConstraints(a_uint32, b_uint32))
	a_uint64, b_uint64 := 1, 2
	fmt.Println(add(a_uint64, b_uint64), addConstraints(a_uint64, b_uint64))
	a_float32, b_float32 := 1.0, 2.0
	fmt.Println(add(a_float32, b_float32), addConstraints(a_float32, b_float32))
	a_float64, b_float64 := 1.0, 2.0
	fmt.Println(add(a_float64, b_float64), addConstraints(a_float64, b_float64))

	sl := []int{1, 2, 4, 8}
	sl = Scale(sl, 2)
	fmt.Println(sl)

}