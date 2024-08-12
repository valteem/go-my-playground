package reuse_test

import (
	"encoding/json"
	"io"
	"strings"
	"testing"
)

type Item struct {
	ID   int
	Desc string
}

func TestJSONDecode(t *testing.T) {
	//	input := "POST\r\nHTTP 1.1\r\nContent-Type: application/json\r\n\r\n{\"ID\":1001,\"Desc\":\"Some Name\"}"
	input := "{\"ID\":1001,\"Desc\":\"Some Name\"}"
	var it Item
	decoder := json.NewDecoder(strings.NewReader(input))
	err := decoder.Decode(&it)
	if err != nil {
		t.Errorf("error decoding input string: %v", err)
	}
}

func TestBufRead(t *testing.T) {
	input := "POST\r\nHTTP 1.1\r\nContent-Type: application/json\r\n\r\n{\"ID\":1001,\"Desc\":\"Some Name\"}"
	reader := strings.NewReader(input) // need bufio.Reader to reproduce httt.Request.Body()
	buf := new(strings.Builder)
	io.Copy(buf, reader)
	outputActual := buf.String()
	outputExpected := "{\"ID\":1001,\"Desc\":\"Some Name\"}"
	if outputActual != outputExpected {
		t.Errorf("get %s, expect %s", outputActual, outputExpected)
	}
}
