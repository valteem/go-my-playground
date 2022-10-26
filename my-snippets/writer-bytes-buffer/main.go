// https://yourbasic.org/golang/io-writer-interface-explained/

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var buf bytes.Buffer // in-memory buffer
	n, _ := fmt.Fprintf(&buf, "Some text including %d number", 7)
	s := buf.String()
	fmt.Printf("%s is %d bytes long\n", s, n)
	io.WriteString(os.Stdout, s)
}
