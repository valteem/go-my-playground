package main

import (
	"lang.rev/nested-main/f"
	nestedmain "lang.rev/nested-main"
)

func main() {
    ff.Main()          // call from sibling folder
	nestedmain.Ring() // call from parent (root) folder (package name and folder name are different!)
}