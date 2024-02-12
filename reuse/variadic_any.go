package reuse

import (
	"reflect"
)

func VarAny(args ...any) []string {
	var r []string
	for _, a := range args {
		r = append(r, reflect.TypeOf(a).Name()) 
	}
	return r
}