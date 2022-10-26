package main

import "fmt"

func main() {

	creatures := []string{"elfs", "hobbits", "orcs", "goblins", "wizards"}

	for i, c := range creatures {
		fmt.Println(i, c)
		switch c{
		case "elfs", "hobbits":
			fmt.Println(c, "are good")
		case "orcs", "goblins":
			fmt.Println(c, "are bad")
			fallthrough
		case "wizards":
			fmt.Println(c, "are both")
			fallthrough
		case "motherfuckers":
			fmt.Println(c, "are no more")
		}

	}
}