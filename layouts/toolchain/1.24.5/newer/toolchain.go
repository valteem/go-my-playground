package newer

import (
	"runtime"
)

// Imported from another module with older Go version
func WrapToolchain() string {
	return runtime.Version()
}
