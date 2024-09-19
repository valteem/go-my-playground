package reuse_test

import (
	"testing"
	"unicode/utf8"
)

func TestRuneCount(t *testing.T) {

	s := "012345世世丗丈"

	if actualSizeBytes, expectedSizeBytes := len(s), (6 + 3*4); actualSizeBytes != expectedSizeBytes {
		t.Errorf("byte count: get %d, expect %d", actualSizeBytes, expectedSizeBytes)
	}

	if actualSizeRunes, expectedSizeRunes := utf8.RuneCountInString(s), 10; actualSizeRunes != expectedSizeRunes {
		t.Errorf("rune count: get %d, expect %d", actualSizeRunes, expectedSizeRunes)
	}

}
