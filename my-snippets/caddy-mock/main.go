package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {

    tc := struct {
	    	input      string
	     	expect     map[string]string
	    	shouldErr  bool
	        }{
		         
	    	    	input: `
	    	            KEY=value
		                OTHER_KEY=Some Value
		                `,
		            expect: map[string]string{
		        	    "KEY":       "value",
		        	    "OTHER_KEY": "Some Value",	            
	                },
	        }

    nr := strings.NewReader(tc.input)

	scanner := bufio.NewScanner(nr)

	var lineNum int
    var args []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			args = append(args, line)
		    lineNum++
		}
	}

	fmt.Println(args)

}