package reuse_test

import (
	"fmt"
	"testing"

	"github.com/valteem/reuse"
)

func TestConvertParamToSliceOfString(t *testing.T) {

	p := make([]interface{}, 3)

	p[0] = struct {
		Name string
		Age int
		} {Name: "name", Age: 1,}
	
	p[1] = 99
	
	p[2] = "text"

	fmt.Println(reuse.ConvertParamToSliceOfString(p))
}