package app

import (
	"testing"
)

func TestApp(t *testing.T) {

	app := NewApp()

	if actual, expected := app.url, "username@host:port/resource"; actual != expected {
		t.Errorf("food name: get %q, expect %q", actual, expected)
	}

}
