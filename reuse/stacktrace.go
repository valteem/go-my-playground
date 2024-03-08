package reuse

import (
	"fmt"
	"runtime"
	"strings"
)

func StackTrace() string {
	var s strings.Builder
	var f string
	pc := make([]uintptr, 32)
	n := runtime.Callers(2, pc) // 2 - Callers() + StackTrace()
	frames := runtime.CallersFrames(pc[:n])
	for n > 0 {
		frame, _ := frames.Next()
		f = fmt.Sprintf("%s:%d %s\n", frame.File, frame.Line, frame.Func.Name())
		s.WriteString(f)
		n -= 1
	}
	return s.String()
}
