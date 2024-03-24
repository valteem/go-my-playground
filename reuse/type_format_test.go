package reuse_test

import (
	"fmt"
	"testing"
)

type entryPerson struct {
	age  int
	name string
}

type entryCar struct {
	carmodel string
	owner    entryPerson
}

var (
	someInt    = 1
	someFloat  = 1.0
	someString = "txt"
)

func TestTypeFormat(t *testing.T) {
	// local struct definition
	type localPerson struct {
		name string
		age  int
	}

	tests := []struct {
		description       string
		typedVar          any
		typeFormattedName string
	}{
		{
			description:       "package variable type struct (entryPerson)",
			typedVar:          entryPerson{age: 30, name: "some name"},
			typeFormattedName: "reuse_test.entryPerson",
		},
		{
			description:       "package variable type struct (entryCar)",
			typedVar:          entryCar{carmodel: "some model", owner: entryPerson{age: 30, name: "some name"}},
			typeFormattedName: "reuse_test.entryCar",
		},
		{
			description:       "package variable type int",
			typedVar:          someInt,
			typeFormattedName: "int", // global type, no package name prefix
		},
		{
			description:       "package variable type float",
			typedVar:          someFloat,
			typeFormattedName: "float64", // global type, no package name prefix
		},
		{
			description:       "package variable type string",
			typedVar:          someString,
			typeFormattedName: "string", // global type, no package name prefix
		},
		{
			description: "composit literal",
			typedVar: struct {
				name  string
				count int
			}{
				name:  "some name",
				count: 1,
			},
			typeFormattedName: "struct { name string; count int }",
		},
		{
			description:       "local variable type local struct",
			typedVar:          localPerson{name: "some name", age: 31},
			typeFormattedName: "reuse_test.localPerson", // package name prefix despite variable being local
		},
	}

	for _, tst := range tests {
		if fmt.Sprintf("%T", tst.typedVar) != tst.typeFormattedName {
			t.Errorf("%s: formatted name must be %s, receive %T", tst.description, tst.typeFormattedName, tst.typedVar)
		}
	}
}
