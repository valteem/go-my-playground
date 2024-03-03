package log

import (
	"bytes"
	"fmt"
	"regexp"
	"testing"
)

const (
	regexPID = `[0-9]+` // `+` - at least one occurence is required as opposed to `*`
	// TODO: add more sofisticated regexp patterns for time and date
	// https://stackoverflow.com/questions/1711727/regular-expression-to-match-dates-in-yyyy-mm-dd-format
	regexDate  = `[0-9]{4}/[0-9]{2}/[0-9]{2}`
	regexTime  = `[0-9]{2}:[0-9]{2}:[0-9]{2}`
	regexMicro = `\.[0-9]{6}` // unescaped `.` matches any character
)

type tester struct {
	description string
	message     string
	pattern     string // regexp to match
}

func TestLoggerDebug(t *testing.T) {
	tests := []tester{
		{
			description: "no trailing newline, logger adds newline",
			message:     "some message",
			pattern:     fmt.Sprintf("^atq: pid=%s %s %s%s DEBUG: some message\n$", regexPID, regexDate, regexTime, regexMicro),
		},
		{
			description: "trailing newline, logger adds newline",
			message:     "some message\n",
			pattern:     fmt.Sprintf("^atq: pid=%s %s %s%s DEBUG: some message\n$", regexPID, regexDate, regexTime, regexMicro),
		},
	}

	for _, tst := range tests {
		var buf bytes.Buffer
		logger := NewLogger(newBase(&buf))

		logger.Debug(tst.message)

		result := buf.String()
		match, e := regexp.MatchString(tst.pattern, result)
		if e != nil {
			t.Fatal("pattern does not compile", e) // not actually what we want to test, hence Fatal(), not Error()
		}
		if !match {
			t.Errorf("logger.Debug(%q outputs %q, must match pattern %q", tst.message, result, tst.pattern)
		}
	}
}

type formatTester struct {
	description string
	format      string
	args        []any
	pattern     string // regexp to match
}

func TestLoggerDebugf(t *testing.T) {
	tests := []formatTester{
		{
			description: "no trailing newline, logger adds newline",
			format:      "some %s",
			args:        []any{"message"},
			pattern:     fmt.Sprintf("^atq: pid=%s %s %s%s DEBUG: some message\n$", regexPID, regexDate, regexTime, regexMicro),
		},
		{
			description: "trailing newline, logger adds newline",
			format:      "some %s\n",
			args:        []any{"message"},
			pattern:     fmt.Sprintf("^atq: pid=%s %s %s%s DEBUG: some message\n$", regexPID, regexDate, regexTime, regexMicro),
		},
	}

	for _, tst := range tests {
		var buf bytes.Buffer
		logger := NewLogger(newBase(&buf))

		logger.Debugf(tst.format, tst.args...)

		result := buf.String()
		match, e := regexp.MatchString(tst.pattern, result)
		if e != nil {
			t.Fatal("pattern does not compile", e) // not actually what we want to test, hence Fatal(), not Error()
		}
		if !match {
			t.Errorf("logger.Debug(%q, %v) outputs %q, must match pattern %q", tst.format, tst.args, result, tst.pattern)
		}
	}
}
