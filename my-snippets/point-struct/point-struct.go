package main

import "fmt"

type Page struct {
	Title string
	Body []byte
}

func PageInitVal(title string, body string) Page {
	return Page{Title: title, Body: []byte(body) }
}

func PageInitPtr(title string, body string) *Page {
	return &Page{Title: title, Body: []byte(body) }
}

func main() {
	title := "Page title"
	body := "Page body"

	p1 := PageInitVal(title, body)
	fmt.Println(p1)

	p2 := *PageInitPtr(title, body)
	fmt.Println(p2)

	p2Ptr := PageInitPtr(title, body)
	fmt.Println(p2Ptr)

}