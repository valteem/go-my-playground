package model

type Feature struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}

type Product struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Features    []Feature `json:"features"`
}
