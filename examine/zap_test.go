package examine

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"strconv"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

type logMsg struct {
	Level string `json:"level"`
	Msg   string `json:"msg"`
}

func NewCustomLogger(w io.Writer) zapcore.Core {
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zap.CombineWriteSyncers(os.Stderr, zapcore.AddSync(w)),
		zapcore.InfoLevel,
	)
}

func TestCustomLogger(t *testing.T) {
	buf := &bytes.Buffer{}
	customCore := NewCustomLogger(buf)
	logger := zap.New(customCore)
	logger.Error("some error")
	output := buf.Bytes()
	lm := &logMsg{}
	err := json.Unmarshal(output, &lm)
	if err != nil {
		t.Fatalf("failed to decode logger output: %v", err)
	}
	if lm.Level != "error" {
		t.Errorf("log level: get %q, expect \"error\"", lm.Level)
	}
	if lm.Msg != "some error" {
		t.Errorf("log message: get %q, expect \"some error\"", lm.Msg)
	}
}

func TestOutputUsingObserver(t *testing.T) {

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stderr,
		zapcore.InfoLevel,
	)

	observed, logs := observer.New(zapcore.InfoLevel)

	logger := zap.New(zapcore.NewTee(core, observed))

	for i := 0; i < 5; i++ {
		logger.Error("info " + strconv.Itoa(i))
	}

	for i, output := range logs.All() {
		outputMessageActual := output.Entry.Message
		outputMessageExpected := "info " + strconv.Itoa(i)
		if outputMessageActual != outputMessageExpected {
			t.Errorf("get %s, expect %s", outputMessageActual, outputMessageExpected)
		}
	}

}
