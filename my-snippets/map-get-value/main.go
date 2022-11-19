package main

import (
	"fmt"
	"regexp"
	"strings"
)

type mp map[string]string

type es struct {
	field string
}

func main() {

	m := mp{}
	m["valid key"] = "valid value"
	fmt.Println(m["valid key"], "\n", m["non-exisiting key"]) // m["non-exisiting key"] returns empty string
	
	mr := make(map[string]*regexp.Regexp)
	regex, ok := mr["non-existing string"] // mr["non-existing string"] return 'zero value', for the mp struct in this case

	fmt.Println(regex, ok, &regex) // shows <nil> (refex.expr ?) and some address | good question 'what is a pointer to nil?'

	e := &es{}
	fmt.Println(e, &e)

	mv := make(map[string] regexp.Regexp)
	rv, okv := mv["non-existing string"]
	fmt.Println(rv, okv) // returns a lot of {<nil> <nil> 00 [] [] false false}, looks like initalized but empty map

	path := "abc/123"
	slash := strings.IndexByte(path, '/')
	fmt.Println(slash)

	vars := []any{"param1", "param2"}
	fmt.Println(ptype(vars))
	s := "param1"
	vars1 := []interface{}{&s}
	fmt.Println(ptype(vars1))

 }

 func ptype (vars ...any) string { // same as (vars []interface{})
	switch p := vars[0].(type) {
	case string:
		fmt.Println(p)
		return "string"
	case *string:
		fmt.Println(p)
		return "*string"
	case int:
		fmt.Println(p)
		return "int"
	default:
		fmt.Println(p)
		return "invalid"
	}
 }