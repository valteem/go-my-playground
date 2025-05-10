package main

import (
	"testing"
)

func TestBuildTimeOutput(t *testing.T) {

	t.Errorf("error message to be seen at docker build time")

}
