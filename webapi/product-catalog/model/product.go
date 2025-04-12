package model

type Feature struct {
	Name  string
	Value any
}

type Product struct {
	Id          int
	Description string
	Features    []Feature
}
