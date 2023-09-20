package sorting_test

import (
	"reflect"
	"testing"

	"my.play.go/toona/sequences"
	"my.play.go/toona/sorting"
)

func TestInsSortPosList(t *testing.T) {

	pl := sequences.NewPList[string]()

	values := []string{"apple", "onion", "cherry", "berry", "plum", "pear"}
	for _, v := range values {
		pl.InsertToHead(v)
	}

	sorting.InsSortPosList[string](pl, func(x string, y string) bool {return x > y})
	result := []string{}
	pos, _ := pl.Tail()
	i := 0
	for i < len(values) {
		elt := pos.Element()
		result = append(result, elt)
		i++
		pos, _ = pl.After(pos)
	}

	expected := []string{"apple", "berry", "cherry", "onion", "pear", "plum"}
	for j := 0; j < len(values); j++ {
		if !reflect.DeepEqual(expected[j], result[j]) {
			t.Errorf("wrong sort result: expect %+v, receive %+v", expected[j], result[j])
		}
	}
}