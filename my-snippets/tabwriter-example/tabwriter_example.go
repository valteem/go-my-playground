package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

func main() {
	var tbl = [][]string{
		{"fruit", "apple", "sweet"},
		{"veg","tomato","round"},
		{"tree","fir", "scentic"},
	}

	for arg1, arg2 := range tbl {
		fmt.Println(arg1, arg2)
	}

	writer := tabwriter.NewWriter(os.Stdout, 15, 8, 0, '\t', 0)
	for _, line := range tbl {
		fmt.Fprintln(writer, strings.Join(line, "\t"))
	}
	writer.Flush()
}