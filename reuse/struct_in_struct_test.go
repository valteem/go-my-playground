package reuse_test

import (
	"testing"
	"runtime"
	"github.com/stretchr/testify/assert"
	"github.com/valteem/reuse"
)

func TestNewBuilder(t *testing.T) {
	b := reuse.NewBuilder()
	s := string(runtime.GOOS)
	assert.Equal(t, s, b.OS, "should be equal")
}