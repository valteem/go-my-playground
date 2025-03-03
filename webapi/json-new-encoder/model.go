package main

type Address struct {
	City    string `json:"city"`
	Street  string `json:"street"`
	ZipCode string `json:"zipcode"`
}

type Person struct {
	GivenName  string   `json:"givenname"`
	FamilyName string   `json:"familyname"`
	Age        int      `json:"age"`
	Address    *Address `json:"address"`
}
