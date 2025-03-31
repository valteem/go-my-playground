package logging

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"testing"
)

var (
	msgNotLogged = "this message is not logged"
	msgLogged    = "this message is logged"
)

func TestLoglevel(t *testing.T) {

	level := slog.LevelVar{}
	level.Set(slog.LevelInfo)

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create os.Pipe")
	}

	logger := slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{Level: &level}))

	logger.Debug(msgNotLogged)
	level.Set(slog.LevelDebug)
	logger.Debug(msgLogged)

	w.Close()

	var buf bytes.Buffer
	n, err := io.Copy(&buf, r)
	if n == 0 || err != nil {
		t.Fatalf("failed to copy from os.Pipe to buffer: %v, %d bytes read", err, n)
	}

	output := buf.String()
	if !strings.Contains(output, msgLogged) {
		t.Errorf("log message:\ns\n%s\nshould contain\n%s\n", output, msgLogged)
	}
	if strings.Contains(output, msgNotLogged) {
		t.Errorf("log message:\ns\n%s\nshould not contain\n%s\n", output, msgNotLogged)
	}

}

func TestStderrRedirect(t *testing.T) {

	msg := "just a random message"

	currentStderr := os.Stderr
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create os.Pipe")
	}
	os.Stderr = w

	fmt.Fprintf(os.Stderr, "%s", msg)

	w.Close()
	os.Stderr = currentStderr

	var buf bytes.Buffer
	n, err := io.Copy(&buf, r)
	if n == 0 || err != nil {
		t.Fatalf("failed to copy from os.Pipe to buffer: %v, %d bytes read", err, n)
	}

	if output := buf.String(); output != msg {
		t.Errorf("log message:\nget\n%s\nexpect\n%s\n", output, msg)
	}

}
