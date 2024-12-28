package main

import (
	"net/http"
	"testing"
)

type PtrToHeader *http.Header

func TestPtrToHeader(t *testing.T) {

	key, value := "Custom-Header", "some custom header value"

	// right-hand-side conversion
	p := PtrToHeader(&http.Header{})

	// does not compile: p.Add() undefined
	//p.Add(key, value)

	// Left-hand-side conversion
	(*http.Header)(p).Add(key, value)

	// right-hand-side conversion to get actual header value
	if actual, expected := (*http.Header)(p).Values(key)[0], value; actual != expected {
		t.Errorf("get %q, expect %q", actual, expected)
	}

}
