package reuse

import (
	"fmt"
)

const (
	maxDepth = 3
)

func nestedCall(s []string) {
	s = append(s, "more")
	if len(s) > maxDepth {
		fmt.Println("Panic")
		panic("full!")
	}
	defer fmt.Println("Deferring ", s)
	fmt.Println("Printing", s)
	nestedCall(s)
}

func CallStackPanicRecover() (err error) {
	defer func() {
		if r := recover(); r != nil { // defer parameters evaluated at call time, no parameters - no evaluation :)
			err = fmt.Errorf("%s", r)
		}
	}()

	s := []string{}
	fmt.Println("Deferring on top level")
	fmt.Println("Initial call")
	nestedCall(s)
	err = nil
	return err
}
