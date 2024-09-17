package main

import (
	//	"database/sql"
	"log"
	"net/http"
)

//var db *sql.DB

/*
func init() {
	var err error
	db, err = sql.Open("postgres", "user=your_user password=your_password dbname=your_database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
*/

func main() {

	initStorage()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/users/register", registerUser)
	mux.HandleFunc("/api/users/authenticate", authenticateUser)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
