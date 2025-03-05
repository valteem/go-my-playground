package main

type Address struct {
	City    string `yaml:"city"`
	Street  string `yaml:"street"`
	ZipCode string `yaml:"zipcode"`
}

type Person struct {
	Name    string   `yaml:"name"`
	Age     int      `yaml:"age"`
	Address *Address `yaml:"address"`
}

type PersonEmbeddedAddress struct {
	Name    string  `yaml:"name"`
	Age     int     `yaml:"age"`
	Address Address `yaml:",inline"` // panics if Address is a pointer to struct, not struct value
}
