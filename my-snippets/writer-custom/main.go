package main

import "fmt"

type CustomWriter struct {
	status bool
}

func (cw *CustomWriter) Write(p []byte) (int, error) {
	cw.status = true
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	cw := CustomWriter{status:false}
	l, e := cw.Write([]byte("some text"))
	fmt.Printf("%d bytes written to who knows where with error %v %s\n", l, e, e)
	fmt.Fprintf(&cw, "%s", "another text")
}