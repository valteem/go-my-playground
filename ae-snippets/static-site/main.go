package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", serveTemplate)

	log.Println("Listening on :3000 ...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {

	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

	// Return a 404 if the template does not exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
		}
	}

	// Return a 404 if the request path is for a folder
	if info.IsDir() {
		http.NotFound(w, r)
	}

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}
	err = tmpl.ExecuteTemplate(w,"layout", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)		
	}

}