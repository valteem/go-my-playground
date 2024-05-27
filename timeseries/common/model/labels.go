package model

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
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

type LabelNames []LabelName

func (lns LabelNames) Len() int {
	return len(lns)
}

func (lns LabelNames) Less(i, j int) bool {
	return lns[i] < lns[j]
}

func (lns LabelNames) Swap(i, j int) {
	lns[i], lns[j] = lns[j], lns[i]
}

func (lns LabelNames) String() string {
	labelStrings := make([]string, 0, len(lns))
	for _, label := range lns {
		labelStrings = append(labelStrings, string(label))
	}
	return strings.Join(labelStrings, ", ")
}

type LabelValue string

func (lv LabelValue) IsValid() bool {
	return utf8.ValidString(string(lv))
}

type LabelValues []LabelValue

func (lvs LabelValues) Len() int {
	return len(lvs)
}

func (lvs LabelValues) Less(i, j int) bool {
	return lvs[i] < lvs[j]
}

func (lvs LabelValues) Swap(i, j int) {
	lvs[i], lvs[j] = lvs[j], lvs[i]
}

type LabelPair struct {
	Name  LabelName
	Value LabelValue
}

type LabelPairs []*LabelPair

func (lps LabelPairs) Len() int {
	return len(lps)
}

func (lps LabelPairs) Less(i, j int) bool {
	switch {
	case lps[i].Name > lps[j].Name:
		return false
	case lps[i].Name < lps[j].Name:
		return true
		// compare values if names are equal
	case lps[i].Value > lps[j].Value:
		return false
	case lps[i].Value < lps[j].Value:
		return true
	default:
		return false // if both names and values are equal
	}
}

func (lps LabelPairs) Swap(i, j int) {
	lps[i], lps[j] = lps[j], lps[i]
}
