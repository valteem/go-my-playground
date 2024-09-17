package main

import (
	//	"database/sql"
	"encoding/base64"
	"encoding/json"
	//	_ "github.com/lib/pq"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func registerUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		log.Println("Error decoding request body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Hash password using bcrypt or another suitable hashing algorithm
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		log.Println("Error generating password hash", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Encode the hashed password for storage in the database
	encodedPassword := base64.StdEncoding.EncodeToString(hashedPassword)

	/*
		db, err := sql.Open("postgres", "user=your_user password=your_password dbname=your_database sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		_, err = db.Exec(`INSERT INTO users (login, password) VALUES ($1, $2)`, newUser.Login, encodedPassword)
		if err != nil {
			log.Println("Error inserting record into DB", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	*/

	err = addUserCredentials(newUser.Login, encodedPassword)
	if err != nil {
		log.Printf("Error adding new user %s: %v", newUser.Login, err)
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}

// TODO: get user credentials from Authorization request header insted of request body
func authenticateUser(w http.ResponseWriter, r *http.Request) {
	var loginUser User
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		log.Println("Error decoding request body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	/*
		db, err := sql.Open("postgres", "user=your_user password=your_password dbname=your_database sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		var storedUser User
		row := db.QueryRow(`SELECT id, login, password FROM users WHERE login=$1`, loginUser.Login)
		err = row.Scan(&storedUser.ID, &storedUser.Login, &storedUser.Password)
		if err == sql.ErrNoRows {
			log.Println("User not found")
			w.WriteHeader(http.StatusNotFound)
			return
		} else if err != nil {
			log.Println("Error fetching user from DB", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	*/

	storedPassword, err := findUserCredentials(loginUser.Login)
	if err != nil {
		log.Println("Error retrieving password", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Decode and compare passwords
	storedPasswordBytes, err := base64.StdEncoding.DecodeString(string(storedPassword))
	if err != nil {
		log.Println("Error decoding password from DB", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword(storedPasswordBytes, []byte(loginUser.Password))
	if err != nil {
		log.Println("Invalid credentials")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Authentication successful"))
}
