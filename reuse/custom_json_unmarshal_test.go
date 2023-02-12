package reuse_test

import (
	"fmt"
	"testing"
	"github.com/valteem/reuse"
)

func TestCustomUnmarshal(t *testing.T) {
	b := []byte(`{"ref":"someref", "books":[{"title":"Title1", "size":"medium", "pages":100}, {"title":"Title2", "size":"large", "pages":1000}]}`)
	u := reuse.OrderUnmatshal(b)
	fmt.Println(u)

	s := []byte(`{"ref":"someref", "books":[ ["Title1", "medium", 100], ["Title2", "large", 1000] ] }`)
	v := reuse.OrderUnmatshal(s)
	fmt.Println(v)

	w := reuse.ShortOrderUnmatshal(s)
	fmt.Println(w)
}