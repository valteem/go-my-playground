package main

import (
	"fmt"
	"strings"
)

func match(path string, pattern string, vars []string) bool {

	for ; len(path) != 0 && len(pattern) != 0; pattern = pattern[1:] {
		switch pattern[0] {
		case '+':
			slash := strings.IndexByte(path, '/')
			if slash < 0 {
				slash = len(path)
			}
			segment := path[:slash]
			path = path[slash:]
			p := &vars[0]
			*p = segment
			vars = vars[1:]
		case path[0]:
			path = path[1:]
		default:
			return false
		}
	}
	return len(pattern) == 0 && len(path) == 0
}

func main() {

	path := "/p1/mycal/p2/myevent"
	pattern := "/p1/+/p2/+"
	
	params := make([]string, 2)

	result := match(path, pattern, params)
	if result {
		fmt.Println(params)
	} else {
		fmt.Println("no match found")
	}	
}