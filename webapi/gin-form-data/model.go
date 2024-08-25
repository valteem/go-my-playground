package main

type Person struct {
	Name string `form:"name"`
	Age  int    `form:"age"`
	// ContactInfo ContactInfo `form:"contactinfo"`
}

/*
type ContactInfo struct {
	PhoneNumber string  `form:"phonenumber"`
	Address     Address `form:"address"`
}

type Address struct {
	ZIPCode     string `form:"zipcode"`
	City        string `form:"city"`
	Street      string `form:"street"`
	HouseNumber string `form:"housenumber"`
}
*/
