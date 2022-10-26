package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return arg, errors.New("cannot work with it")
	}
	return arg, nil
}

type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s",e.arg, e.prob)
}

func f2(arg int) (int,error) {
	if arg == 42 {
		return -1, &argError{arg, "cannot work with it"}
	} else {
		return arg, nil
	}
}

func main() {

	for _, i := range []int{41,42} {
		if r, e := f1(i); e != nil {
			fmt.Println(e)
		} else {
			fmt.Println(r)}	
	}

	for _, i := range []int{41,42} {
		if r, e := f2(i); e != nil {
			fmt.Println(e)
		} else {
			fmt.Println(r)}	
	}
	
	_, e := f2(42)
	 
	ae, ok := e.(*argError)
	if ok {
		fmt.Println(ae.arg)
	    fmt.Println(ae.prob)
	}
}
