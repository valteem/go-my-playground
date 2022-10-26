package main

import (
	"net/http"
	"path/filepath"
)

func main() {

	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {

	fp := filepath.Join("dipole_field_calc.jpg")
	http.ServeFile(w, r, fp)
	
}