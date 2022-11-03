package main

import (
	"fmt"
	"lang.rev/init-package-var/initvar"
)

func main() {

	fmt.Println("Variable A:", initvar.VarA)
	fmt.Println("Variable B:", initvar.VarB)
	fmt.Println("Variable C:", initvar.VarC)
	fmt.Println("Reference to variable A:", initvar.RefA)
	fmt.Println("Initializing variables from another source file of the same package:", initvar.WarmUp)

}