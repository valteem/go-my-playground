package model

import "time"

type Feature struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}

type Product struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Features    []Feature `json:"features"`
}

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
