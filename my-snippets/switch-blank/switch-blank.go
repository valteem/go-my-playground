package main

import (
	"flag"
	"log"
)

func main() {

	ptr1 := flag.Bool("add", false, "add message")
	ptr2 := flag.Bool("get", false, "get message")

	flag.Parse()

	log.Println(flag.Arg(0), flag.Arg(1), flag.Arg(2))

	arg1 := flag.Args()[0]

	switch {
	case *ptr1:
		log.Println("add: ", arg1, "ptr: ", *ptr1)
	case *ptr2:
		log.Println("get: ", arg1)
	default:
		log.Println("nothing")
	}
}