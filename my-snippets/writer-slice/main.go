// Example of why it is bad 'to retain slice' provided to io.Writer.Write method
// https://unexpected-go.com/write-should-not-retain-the-slice.html 
package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type bufWriter struct {
	output io.Writer
	buffer chan []byte
}

func (bw bufWriter) Start() {
	go func () {
		for b := range bw.buffer {
			bw.output.Write(b)
		}
	}() // using goroutine (probably) because we start our writer before anything is written to the buffer
}

func (bw bufWriter) Write(b []byte) (int, error) {
	bw.buffer <- b
	return len(b), nil
}

func main() {
	writer := bufWriter{output: os.Stdout, buffer: make(chan []byte, 10)}
	writer.Start()

	for i := 0; i < 10; i++ {
		fmt.Fprintf(writer, "%d\n", i)
	}

	time.Sleep(time.Second)
}