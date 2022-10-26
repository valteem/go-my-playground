package main

import (
	"fmt"
	"reflect"
)

type Book struct {
	Title  string
	Volume int
}

type Books []Book

func main() {

	b := Books{{Title: "T1", Volume: 101}, {Title: "T2", Volume: 201}}
	fmt.Println(b)
	t := b.GetFieldValuesAsSlice("Title")
	fmt.Printf("%-75s %s\n", "Book titles read with FieldByName, object method:", t)
	ts := b.GetTitles()
	fmt.Printf("%-75s %s\n", "Book titles read as struct fields:", ts)

	bb := []Book{{Title: "T1", Volume: 101}, {Title: "T2", Volume: 201}}
	fmt.Printf("%-75s %s\n", "Type of books slice:", reflect.TypeOf(bb))
	tbb := GetFieldValuesAsSlice(bb, "Title")
	fmt.Printf("%-75s %s\n", "Book titles read with FieldByName, function with book slice as argument:", tbb)

}

func (b Books) GetFieldValuesAsSlice(fieldname string) []string {

	var s []string

	for _, book := range b {
		r := reflect.ValueOf(book)
		f := reflect.Indirect(r).FieldByName(fieldname)
		s = append(s, f.String())
	}

	return s
}

func (b Books) GetTitles() []string {

	var s []string

	for _, book := range b {
		t := book.Title
		s = append(s, t)
	}

	return s
}

func GetFieldValuesAsSlice(b []Book, fieldname string) []string {

	var s []string

	for _, book := range b {
		r := reflect.ValueOf(book)
		f := reflect.Indirect(r).FieldByName(fieldname)
		s = append(s, f.String())
	}

	return s
}
