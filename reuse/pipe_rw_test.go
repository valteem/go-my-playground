package reuse_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"sync"
	"testing"
)

func TestPipeRW(t *testing.T) {

	r, w := io.Pipe()
	defer w.Close()

	buf := bytes.NewBuffer(nil)

	msgWrite, msgRead := "abc", ""

	var wg sync.WaitGroup
	wg.Add(100)

	go func() {
		// io.Copy(dst, src) copies from src to dst until EOF is reached
		_, e := io.Copy(buf, r)
		if e != nil {
			t.Errorf("Error copying to buffer: %s", e)
		}
	}()

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Fprint(w, msgWrite)
			wg.Done()
		}()
		msgRead += msgWrite
	}

	wg.Wait()

	if output := buf.String(); output != msgRead {
		t.Errorf("Reading from buffer: get %s, expect %s", output, msgRead)
	}

}

func TestPipeWrite(t *testing.T) {
	r, w := io.Pipe()
	m := "msg"
	go func() {
		fmt.Fprint(w, m)
		w.Close()
	}()
	buf := new(strings.Builder)
	_, e := io.Copy(buf, r)
	if e != nil {
		t.Errorf("Error reading from buffer: %s", e)
	}
	if output := buf.String(); output != m {
		t.Errorf("Read from buffer: get %s, expect %s", output, m)
	}
}

func TestEmptyPipe(t *testing.T) {

	r, w := io.Pipe()
	w.Close()

	buf := &strings.Builder{}

	_, err := io.Copy(buf, r)

	if err != nil {
		t.Fatalf("failed to read from empty PipeReader: %v", err)

	}
}
