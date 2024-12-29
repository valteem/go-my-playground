package main

import (
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
