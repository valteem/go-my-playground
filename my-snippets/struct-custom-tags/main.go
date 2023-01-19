package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	ID int `ctag1:"ID" ctag2:"AltID"`
	Name string `ctag1:"Name" ctag2:"AltName"`
	Email string `ctag1:"Email" ctag2:"AltEmail"`
}

func main() {

	p := Person{ID: 1, Name: "pname", Email: "noone@nowhere.com"}
	t := reflect.TypeOf(p)

	for _, fname := range []string{"ID", "Name", "Email"} {
		field, found := t.FieldByName(fname)
		if found {
			fmt.Println(fname, field.Name, field.Tag, field.Tag.Get("ctag1"), field.Tag.Get("ctag2"))
		}
	}
}