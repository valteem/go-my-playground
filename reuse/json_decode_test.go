// Decode() vs Unmarshal(): https://stackoverflow.com/a/21198571

package reuse_test

import (
	"encoding/json"
	"io"
	"reflect"
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
	var output Item
	expectedOutput := Item{ID: 1001, Desc: "Some Name"}
	decoder := json.NewDecoder(strings.NewReader(input))
	err := decoder.Decode(&output)
	if err != nil {
		t.Errorf("error decoding input string: %v", err)
	}
	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("error decoding JSON:\nget\n%v\nexpect\n%v\n", output, expectedOutput)
	}
}

func TestBufRead(t *testing.T) {
	input := "POST\r\nHTTP 1.1\r\nContent-Type: application/json\r\n\r\n{\"ID\":1001,\"Desc\":\"Some Name\"}"
	reader := strings.NewReader(input) // need bufio.Reader to reproduce httt.Request.Body()
	buf := new(strings.Builder)
	io.Copy(buf, reader)
	outputActual := buf.String()
	outputExpected := "{\"ID\":1001,\"Desc\":\"Some Name\"}"
	// TODO: filter non-JSON part, use bufio.Scanner
	if outputActual != outputExpected {
		t.Errorf("get\n%s\nexpect\n%s", outputActual, outputExpected)
	}
}
