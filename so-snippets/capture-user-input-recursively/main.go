// https://stackoverflow.com/questions/60649456/how-to-recursively-capture-user-input

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInt(rdr *bufio.Reader, n []int) []int {
    line, err := rdr.ReadString('\n') // func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
	if err != nil {
		log.Fatal(err)
	}
	line = strings.TrimSpace(line)
	if i, err := strconv.Atoi(line); err == nil {
		n = append(n, i)
	}
	if err == io.EOF || strings.ToLower(line) == "end" {
		return n
	}
    return readInt(rdr, n)
}

func readInts() []int {
	fmt.Print("enter integers:\n")
	var n []int
	rdr := bufio.NewReader(os.Stdin)
	return readInt(rdr, n)
}

func main() {
	n := readInts()
	fmt.Println(n)
}