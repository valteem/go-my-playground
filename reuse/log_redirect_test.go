package reuse_test

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
)

func TestLogRedirectToScanner(t *testing.T) {

	logMsg := "some custom log message"

	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create os pipe: %v", err)
	}
	defer reader.Close()
	defer writer.Close()

	log.SetOutput(writer)

	buf := bufio.NewScanner(reader)

	log.Println(logMsg)

	buf.Scan()
	output := buf.Text()
	// trim date-time
	output = output[(len(output) - len(logMsg)):]

	if output != logMsg {
		t.Errorf("log output: get %q, expect %q", output, logMsg)
	}

	log.SetOutput(os.Stderr)

}

func TestLogRedirectToReader(t *testing.T) {

	logMsg := "some custom log message"

	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create os pipe: %v", err)
	}
	defer reader.Close()
	defer writer.Close()

	log.SetOutput(writer)

	buf := make([]byte, 4096)

	log.Printf("%s", logMsg)

	n, err := reader.Read(buf)
	if err != nil {
		t.Fatalf("failed to read: %v", err)
	}
	// output from reader.Read() contains trailing "\n" (contrary to bufio.Scanner())
	// This is because "the split function defaults to [ScanLines]" for a new Scanner
	output := strings.TrimRight(string(buf[:n]), "\n")
	// trim date-time
	output = output[(len(output) - len(logMsg)):]

	if output != logMsg {
		t.Errorf("log output: get %q, expect %q", output, logMsg)
	}

	log.SetOutput(os.Stderr)

}
