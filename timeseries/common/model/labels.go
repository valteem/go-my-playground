package model

import (
	"encoding/json"
	"fmt"
	"regexp"
	"unicode/utf8"
)

const ()

// IsValid() does the same but faster
var LabelNameRE = regexp.MustCompile("^[a-zA-Z_][a-zA-Z0-9_]*$") // * - zero or more characters

type LabelName string

func (ln LabelName) IsValid() bool {
	if len(ln) == 0 {
		return false
	}
	switch NameValidationScheme {
	case LegacyValidation:
		for i, b := range ln {
			if !((b >= 'a' && b <= 'z') ||
				(b >= 'A' && b <= 'Z') ||
				b == '_' ||
				(b >= '0' && b <= '9' && i > 0)) {
				return false
			}
		}
	case UTF8Validation:
		return utf8.ValidString(string(ln))
	default:
		panic(fmt.Sprintf("Invalid name validation scheme: %d", NameValidationScheme))
	}
	return true
}

func (ln *LabelName) UnmarshalJSON(input []byte) error {
	var s string
	if e := json.Unmarshal(input, &s); e != nil {
		return e
	}
	l := LabelName(s)
	if !l.IsValid() {
		return fmt.Errorf("%q is not a valid label name", s)
	}
	*ln = l
	return nil
}
