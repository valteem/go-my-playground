package somelogic_test

import (
	sl "example.com/somelogic"
	"flag"
	"testing"
)

func TestSomeLogic(t *testing.T) {
 	flag.Parse()
	strSlice := flag.Args()
	sl.SomeLogic(strSlice)
}