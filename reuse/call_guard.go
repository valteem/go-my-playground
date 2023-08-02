package reuse

import (
	"log"
)

var (
	guard bool
)

func Init() {

	if guard {
		log.Fatal("Init() has already been called")
	}
	guard = true
}

func Stop() {

	if !guard {
		log.Fatal("nothing to stop")
	}

	guard = false
	
}