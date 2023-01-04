package main

import (
	"fmt"
	"reflect"
	"regexp"
)

func main() {

	r, _ := regexp.Compile("p([a-z]+)ch") // "p[a-z][a-z]ch" works too
	fmt.Println(r, reflect.TypeOf(r)) // (*Regexp) String() provides input for Println()

	s := r.MatchString("peach")
	fmt.Println(s)

	re := regexp.MustCompile("^/contact/([a-zA-Z0-9]+)/event/([0-9]+)")
	fmt.Println(re.FindStringSubmatch("/contact/Name1/event/1234"))

	r1 := regexp.MustCompile("([/]+)")
	fmt.Println(r1.MatchString("//"))

}