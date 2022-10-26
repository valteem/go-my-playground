package main

import "fmt"

func main() {

	var sl []string
	s1 := make([]string, 10)
	s2 := []string{"names", "words", "birds", "things"}
	s := append(s1, s2...)

	f := func(name string, s []string) {
		fmt.Printf("Slice %s has length %d and capacity %d\n", name, len(s), cap(s))
	}

	f("s1", s1)
	f("s2", s2)
	f("s", s)

	sl = append(sl, "str1", "str2")
	f("sl", sl)
	sl = sl[:1]
	f("sl", sl)

}
