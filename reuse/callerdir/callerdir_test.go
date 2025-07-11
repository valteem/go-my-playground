package callerdir

import (
	"runtime"
	"strings"

	"testing"
)

func TestCallerDir(t *testing.T) {

	_, filename, _, _ := runtime.Caller(0)

	fpath := strings.Split(filename, "/")
	fname := fpath[len(fpath)-1]

	if fname != "callerdir_test.go" {
		t.Errorf("file name: get %q, expect %q", fname, "callerdir_test.go")
	}

}
