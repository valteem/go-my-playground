package log

import (
	"fmt"
	"io"
	stdlog "log"
	"os"
	"sync"
)

type Level int32

const (
	DebugLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

// fmt.Stringer.String()
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	default:
		return "unknown"
	}
}

type Base interface {
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any) // exit with status 1
}

type baseLogger struct {
	*stdlog.Logger
}

func (l *baseLogger) prefixPrint(prefix string, args ...any) {
	args = append([]any{prefix}, args...)
	l.Print(args...)
}

func (l *baseLogger) Debug(args ...any) {
	l.prefixPrint("DEBUG: ", args...)
}

func (l *baseLogger) Info(args ...any) {
	l.prefixPrint("INFO: ", args...)
}

func (l *baseLogger) Warn(args ...any) {
	l.prefixPrint("WARN: ", args...)
}

func (l *baseLogger) Error(args ...any) {
	l.prefixPrint("ERROR: ", args...)
}

func (l *baseLogger) Fatal(args ...any) {
	l.prefixPrint("FATAL: ", args...)
	os.Exit(1)
}

func newBase(out io.Writer) *baseLogger {
	prefix := fmt.Sprintf("atq: pid=%d ", os.Getegid())
	return &baseLogger{
		stdlog.New(out, prefix, stdlog.Ldate|stdlog.Ltime|stdlog.Lmicroseconds|stdlog.LUTC),
	}
}

type Logger struct {
	base  Base
	mu    sync.Mutex
	level Level // minimum log level
}

func NewLogger(base Base) *Logger {
	if base == nil {
		base = newBase(os.Stderr)
	}
	return &Logger{base: base, level: DebugLevel} // DebulLevel by default
}

func (l *Logger) canLogAt(v Level) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	return v >= l.level
}

func (l *Logger) Debug(args ...any) {
	if !l.canLogAt(DebugLevel) {
		return
	}
	l.base.Debug(args...)
}

func (l *Logger) Info(args ...any) {
	if !l.canLogAt(InfoLevel) {
		return
	}
	l.base.Info(args...)
}

func (l *Logger) Warn(args ...any) {
	if !l.canLogAt(WarnLevel) {
		return
	}
	l.base.Warn(args...)
}

func (l *Logger) Error(args ...any) {
	if !l.canLogAt(ErrorLevel) {
		return
	}
	l.base.Error(args...)
}

func (l *Logger) Fatal(args ...any) {
	if !l.canLogAt(FatalLevel) {
		return
	}
	l.base.Fatal(args...)
}

func (l *Logger) Debugf(format string, args ...any) {
	l.Debug(fmt.Sprintf(format, args...))
}

func (l *Logger) Infof(format string, args ...any) {
	l.Info(fmt.Sprintf(format, args...))
}

func (l *Logger) Warnf(format string, args ...any) {
	l.Warn(fmt.Sprintf(format, args...))
}

func (l *Logger) Errorf(format string, args ...any) {
	l.Error(fmt.Sprintf(format, args...))
}

func (l *Logger) Fatalf(format string, args ...any) {
	l.Fatal(fmt.Sprintf(format, args...))
}

func (l *Logger) SetLevelO(v Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if v < DebugLevel || v > FatalLevel {
		panic("atq log: invalid log level")
	}
	l.level = v
}
