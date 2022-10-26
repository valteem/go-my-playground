package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {

// https://stackoverflow.com/questions/46917331/what-is-the-difference-between-backticks-double-quotes-in-golang
// You don't need to escape newline, tab or other special characters in backticks string

	t := `word1 word2  word3   
	     word4`

	scanner := bufio.NewScanner(strings.NewReader(t))
//	scanner.Split(bufio.ScanLines)
    linenum := 0
	for scanner.Scan() {
		linenum++
		fmt.Printf("%02d %s\n", linenum, scanner.Text() )
	}
}