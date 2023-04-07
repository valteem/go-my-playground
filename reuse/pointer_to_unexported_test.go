package reuse_test

import (
	"testing"

	reuse "github.com/valteem/reuse"
)

func TestNewMux(t *testing.T) {

	path := []string{"url1", "url2", "url3"}
	handler := []func(){func() {}, func() {}, func() {}}

	mux := reuse.ExportMux(path, handler)

	for idx, out := range mux.Path {
		if out != path[idx] {
			t.Errorf("%v should be equal %v", out, path[idx])
		}
	}

}
