// https://stackoverflow.com/a/75347496

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(s *bufio.Scanner) ([][]string, error) {

	var input [][]string

	if !s.Scan() {
		return nil, s.Err()
	}

	nline, err := strconv.Atoi(strings.TrimSpace(s.Text()))
	if err != nil {
		return nil, err
	}

	for ; nline > 0 && s.Scan(); nline-- {
		input = append(input, strings.Fields(s.Text()))
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	if nline != 0 {
		err := fmt.Errorf("missing %d lines of data", nline)
		return nil, err
	}

	return input, nil

}

func main() {

	s := bufio.NewScanner(os.Stdin)

	input, err := readInput(s)

	fmt.Println(input, err)
	
}