package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {

	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
	
}

func loadPage(title string) (*Page, error) {

	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file")
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil

}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)

}

func editHandler(w http.ResponseWriter, r *http.Request) {

	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)

}

func main() {

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", viewHandler)
	http.HandleFunc("/save/", viewHandler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}