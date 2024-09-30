package reuse_test

import (
	"encoding/json"
	"reflect"
	"testing"
)

type bookMetadata struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	VolumePages int    `json:"vol"`
}

func TestUnmarshalJSONArray(t *testing.T) {

	var books []bookMetadata

	input := `[{"author":"author1", "title":"book1", "vol":301},` +
		`{"author":"author2", "title":"book2", "vol":302},` +
		`{"author":"author3", "title":"book3", "vol":303}]`

	err := json.Unmarshal([]byte(input), &books)
	if err != nil {
		t.Fatalf("failed to unmarshal input JSON: %v", err)
	}

	outputExpected := []bookMetadata{
		{"author1", "book1", 301},
		{"author2", "book2", 302},
		{"author3", "book3", 303},
	}

	if !reflect.DeepEqual(books, outputExpected) {
		t.Errorf("Ummarshal JSON array to slice of struct:\nget\n%v\nexpect\n%v", books, outputExpected)
	}

}
