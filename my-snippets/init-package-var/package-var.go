package main

import (
	"fmt"
	"lang.rev/init-package-var/initvar"
)

func main() {

	fmt.Println("Variable A:", initvar.VarA)                // variables from start.go
	fmt.Println("Variable B:", initvar.VarB)
	fmt.Println("Variable C:", initvar.VarC)
	fmt.Println("Reference to variable A:", initvar.RefA)
	fmt.Println("Warming up:", initvar.WarmUp)              // variables from warmup.go

}