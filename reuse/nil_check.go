package reuse

import (
	"reflect"
)

func NilCheck(input any) bool {
	return input == nil || (reflect.ValueOf(input).Kind() == reflect.Ptr && reflect.ValueOf(input).IsNil())
}
