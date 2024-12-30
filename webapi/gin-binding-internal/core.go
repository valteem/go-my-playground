package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// HeadTail() divides input string into two parts - head and tail, separated by first instance of separator string
func HeadTail(input, separator string) (string, string) {
	index := strings.Index(input, separator)
	if index < 0 {
		return input, ""
	}
	return input[:index], input[index+len(separator):]
}

func SetIntField(input string, bitSize int, field reflect.Value) error {
	if input == "" {
		input = "0"
	}
	intVal, err := strconv.ParseInt(input, 10, bitSize)
	if err != nil {
		return err
	}
	if !field.CanSet() {
		return fmt.Errorf("cannot set value")
	}
	field.SetInt(intVal)
	return nil
}

func SetWithProperType(input string, value reflect.Value, field reflect.StructField) error {
	switch value.Kind() {
	case reflect.Int:
		return SetIntField(input, 0, value)
	case reflect.String:
		value.SetString(input)
	}
	return nil
}
