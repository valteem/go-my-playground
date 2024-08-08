package examine

import (
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

type addr struct {
	//lint:ignore U1000 for test purposes only
	zipcode int
	//lint:ignore U1000 for test purposes only
	place string
}

type Person struct {
	Name  string
	Age   int
	addrs []*addr
}

func TestSpewDump(t *testing.T) {
	p := Person{Name: "Some Name", Age: 35, addrs: []*addr{{123456, "some city, street and house number"}}}
	s := spew.Sdump(p)
	lines := strings.Split(s, "\n")
	expected := []string{
		"(examine.Person) {",
	}
	if lines[0] != expected[0] {
		t.Errorf("get %q, expect %q", lines[0], expected[0])
	}

}
