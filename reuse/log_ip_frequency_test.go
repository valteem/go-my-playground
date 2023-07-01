package reuse_test

import (
	"fmt"
	"testing"
	"github.com/valteem/reuse"
)

func TestLogIPFrequence(t *testing.T) {

	s := `[111.12.14.1] ping
	[999.12.12.1] more
	[111.12.14.1] another ping
	[10.10.10.255] attempt
	16.14.100.199`

	fmt.Println(reuse.ParseLogForIP(s))
}