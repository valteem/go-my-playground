package examine

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
