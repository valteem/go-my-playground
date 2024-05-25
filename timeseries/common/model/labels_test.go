package model

import (
	"testing"
)

func TestLabelNameIsValid(t *testing.T) {
	tests := []struct {
		ln          LabelName
		legacyValid bool
		utf8Valid   bool
	}{
		{
			ln:          "Valid_12_name",
			legacyValid: true,
			utf8Valid:   true,
		},
		{
			ln:          "_Valid_12_name",
			legacyValid: true,
			utf8Valid:   true,
		},
		{
			ln:          "1Valid_23_name",
			legacyValid: false,
			utf8Valid:   true,
		},
		{
			ln:          "Valid_12:name",
			legacyValid: false,
			utf8Valid:   true,
		},
		{
			ln:          ":Valid_12_name",
			legacyValid: false,
			utf8Valid:   true,
		},
		{
			ln:          "a\xb1",
			legacyValid: false,
			utf8Valid:   false,
		},
	}
	for _, tst := range tests {
		NameValidationScheme = LegacyValidation
		if v := tst.ln.IsValid(); v != tst.legacyValid {
			t.Errorf("Legacy validation: get %t, expect %t for %q", v, tst.legacyValid, tst.ln)
		}
		NameValidationScheme = UTF8Validation
		if v := tst.ln.IsValid(); v != tst.utf8Valid {
			t.Errorf("UTF8 validation: get %t, expect %t for %q", v, tst.utf8Valid, tst.ln)
		}
	}
}

func TestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		input          []byte
		isErr          bool
		expectedString string
	}{
		{
			input:          []byte(`"custom_label"`),
			isErr:          false,
			expectedString: "custom_label",
		},
		{
			input:          []byte(`"1label"`),
			isErr:          true,
			expectedString: "",
		},
	}
	for _, tst := range tests {
		var ln LabelName
		e := ln.UnmarshalJSON(tst.input)
		if e != nil && !tst.isErr {
			t.Errorf("UnmarshalJSON() on %q: get %e", tst.expectedString, e)
		} else {
			if str := string(ln); str != tst.expectedString {
				t.Errorf("UnmarshalJSON(): get %q, expect %q", str, tst.expectedString)
			}
		}
	}
}
