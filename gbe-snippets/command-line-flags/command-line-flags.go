package main

import (
	"flag"
	"log"
)

func main() {

	// One way is to declare and define a pointer with flag.String, flag.Int, flag.Bool etc

	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 31, "an int")
	boolPtr := flag.Bool("fork", true, "a bool (this time)") // last argument is what shows with -help option

	// Another way is to declare a variable first, and then bind it with flag.StringVar
	var svar string
	flag.StringVar(&svar, "svar", "bar", "another string arg")


	flag.Parse()
	
	log.Println(*wordPtr)
	log.Println(*numbPtr)
	log.Println(*boolPtr)
	log.Println(svar)
	log.Println(flag.Args())

}