package main

import (
	"fmt"
	"net/http"
	"testing"
)

type PtrToHeader *http.Header

func TestPtrToHeader(t *testing.T) {

	key, value := "Custom-Header", "some custom header value"

	// type conversion (right-hand-side) + assignment (left-hand side)
	p := PtrToHeader(&http.Header{})

	// does not compile: p.Add() undefined
	//p.Add(key, value)

	if fmt.Sprintf("%p\n", p) != fmt.Sprintf("%p\n", (*http.Header)(p)) {
		t.Errorf("expect two variables share the same address, get\n%s - p\n%s - (*http.Header)(p)", fmt.Sprintf("%p\n", p), fmt.Sprintf("%p\n", (*http.Header)(p)))
	}

	// Type conversion (no assignment)
	(*http.Header)(p).Add(key, value)

	if fmt.Sprintf("%p\n", p) != fmt.Sprintf("%p\n", (*http.Header)(p)) {
		t.Errorf("expect two variables share the same address, get\n%s - p\n%s - (*http.Header)(p)", fmt.Sprintf("%p\n", p), fmt.Sprintf("%p\n", (*http.Header)(p)))
	}

	// getting actual header value - type conversion (right-hand-side) + assignment (left-hand side)
	if actual, expected := (*http.Header)(p).Values(key)[0], value; actual != expected {
		t.Errorf("get %q, expect %q", actual, expected)
	}

}
