package main

import (
	"lang.rev/call-func/src/auxf" // not necessary too use an alias for package name
    Lib "lang.rev/call-func/src/lib" // package name as per go.mod + relative path
	"lang.rev/call-func/src/lib"
	"lang.rev/call-func/libmod"
)

func main() {
	Lib.FuncLib("Hello") // referencing package by alias
	lib2.FuncAux() // referencing package by name
	lib.FuncNext()
	lmod.Announcer()
}