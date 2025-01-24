package main

import (
	"my-snippets/set-minimum-go-version/goversion"
)

func main() {
	// won't compile if Go version is lower than 1.23
	var _ = goversion.EnforceGoVersion1_23
}
