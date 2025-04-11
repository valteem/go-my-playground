package comparator

import (
	"testing"
)

func TestNonZeroInt(t *testing.T) {

	tests := []struct {
		value  int
		alt    int
		output int
	}{
		{0, 42, 42},
		{1, 42, 1},
	}

	for _, tc := range tests {
		if output := NonZero(tc.value, tc.alt); output != tc.output {
			t.Errorf("value/alt (%d/%d)): get %d, expect %d", tc.value, tc.alt, output, tc.output)
		}
	}

}

func TestNonZeroString(t *testing.T) {

	tests := []struct {
		value  string
		alt    string
		output string
	}{
		{"Life", "Universe", "Life"},
		{"", "Universe", "Universe"},
	}

	for _, tc := range tests {
		if output := NonZero(tc.value, tc.alt); output != tc.output {
			t.Errorf("value/alt (%q/%q)): get %q, expect %q", tc.value, tc.alt, output, tc.output)
		}
	}

}

func TestNonZeroStruct(t *testing.T) {

	type person struct {
		name string
		age  uint
	}

	p0 := person{}
	p1 := person{"Douglas Adams", 42}
	p2 := person{"Peter Pan", 12}

	tests := []struct {
		value  person
		alt    person
		output person
	}{
		{p1, p2, p1},
		{p0, p2, p2},
	}

	for _, tc := range tests {
		// no need for reflect.DeepEqual as we compare pre-defined instances of person
		if output := NonZero(tc.value, tc.alt); output != tc.output {
			t.Errorf("value/alt (%v/%v)): get %v, expect %v", tc.value, tc.alt, output, tc.output)
		}
	}

}

func TestNonZeroArray(t *testing.T) {

	type seq [4]int

	s0 := seq{}
	s1 := seq{1, 1, 1, 1}
	s2 := seq{2, 2, 2, 2}

	tests := []struct {
		value  seq
		alt    seq
		output seq
	}{
		{s1, s2, s1},
		{s0, s2, s2},
	}

	for _, tc := range tests {
		// we do not compare in greater than or lower then way, hence using == and !=
		if output := NonZero(tc.value, tc.alt); output != tc.output {
			t.Errorf("value/alt (%v/%v)): get %v, expect %v", tc.value, tc.alt, output, tc.output)
		}
	}

}
