package main

import (
	"fmt"
)

func main() {
	LoopVarInside()
	LoopVarOutside()
}

func LoopVarInside() {
	var prints []func()
	for i := 1; i <= 3; i++ { // loop variable bug fixed - i is new int instance at each iteration
		prints = append(prints, func() { fmt.Println(i) })
	}
	for _, print := range prints {
		print()
	}
}

func LoopVarOutside() {
	var prints []func()
	var i int
	for i = 1; i <= 3; i++ { // i is same across all iterations - 'bug' is still there
		prints = append(prints, func() { fmt.Println(i) })
	}
	for _, print := range prints {
		print()
	}
}
