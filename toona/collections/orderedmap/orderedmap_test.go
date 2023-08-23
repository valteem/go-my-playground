package orderedmap

import (
	// "fmt"
	"reflect"
	"testing"
)

func TestOrderedMapInsert(t *testing.T) {
	om := New[string, string]()
	om.Insert("key1", "value1")
	om.Insert("key2", "value2")
	om.Insert("key3", "value3")

	result := om.Keys()
	expected := []string{"key1", "key2", "key3"}
	if !reflect.DeepEqual(result, expected) {
		t.Error("wrong insert result - keys")
	}

	resultValues := om.Values()
	expectedValues := []string{"value1", "value2", "value3"}
	if !reflect.DeepEqual(resultValues, expectedValues) {
		t.Error("wrong insert result - values")
	}

	om.Insert("key1", "value1new")
	resultValues = om.Values()
	expectedValues = []string{"value1new", "value2", "value3"}
	if !reflect.DeepEqual(resultValues, expectedValues) {
		t.Error("wrong insert result - values after inserting new value for existing key")
	}
}