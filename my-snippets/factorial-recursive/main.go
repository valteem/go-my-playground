package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func factrlRecur(n int) int {
	if n < 0 {
		log.Fatal("wrong input integer (cannot be negative)")
		return 0
	}
	if n == 0 {
		return 1
	}
	return n * factrlRecur(n-1)
}

func main() {
	fmt.Print("enter integer\n")
	rdr := bufio.NewReader(os.Stdin)
	line, err := rdr.ReadString('\n') // "console": "integratedTerminal" - to enable input debugging
	if err != nil {
		log.Fatal(err)
	}
	line = strings.TrimSpace(line)
	if i, err := strconv.Atoi(line); err == nil {
		fmt.Printf("factorial is %d\n", factrlRecur(i))
		return
	}
	log.Fatal(err)
}