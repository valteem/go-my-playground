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
	buf := bytes.NewBuffer(nil)
	msgWrite, msgRead := "abc", "abcabc"
	var wg sync.WaitGroup
	wg.Add(2)

	go func(wg *sync.WaitGroup) {
		fmt.Fprint(w, msgWrite)
		wg.Done()
	}(&wg)
	go func(wg *sync.WaitGroup) {
		fmt.Fprint(w, msgWrite)
		wg.Done()
	}(&wg)

	go func() {
		_, e := io.Copy(buf, r)
		if e != nil {
			t.Errorf("Error copying to buffer: %s", e)
		}
		if output := buf.String(); output != msgRead {
			t.Errorf("Reading from buffer: get %s, expect %s", output, msgRead)
		}
	}()

	wg.Wait()
	go func() { w.Close() }()

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
