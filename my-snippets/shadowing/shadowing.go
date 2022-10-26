package main

import "fmt"

func main() {

	va1 := "this"
	va2 := "this"
	pad_l := 50
	pad_r := 5
	cond := true

	fmt.Printf("Example without in-block declaration (assignment only)\n")
	fmt.Printf("%-*s%-*s\n\r", pad_l, "Main body - before block", pad_r, va1) // %-*s provides for padding
	if cond {
		va1 = "that" // declared in outers scope and re-assigned in-block - changes value in outer scope
		fmt.Printf("%-*s%-*s\n\r", pad_l, "Block", pad_r, va1)
	}
	fmt.Printf("%-*s%-*s\n\r", pad_l, "Main body - after block", pad_r, va1)

    fmt.Printf("\n")

	fmt.Printf("Example with in-block declaration and assignment\n")
	fmt.Printf("%-*s%-*s\n\r", pad_l, "Main body - before block", pad_r, va2)
	if cond {
		va2 := "that" // decclared and assigned in-block - no effect on outer scope same name variable
		fmt.Printf("%-*s%-*s\n\r", pad_l, "Block", pad_r, va2)
	}
	fmt.Printf("%-*s%-*s\n\r", pad_l, "Main body - after block", pad_r, va2)

}