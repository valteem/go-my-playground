package main

import (
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("http://127.0.0.1:8081")
	if err != nil {
		log.Fatalf("failed to get proxy response: %v", err)
	}

	respHeaderProxy, ok := resp.Header[http.CanonicalHeaderKey("x-sidecar-response")]
	if !ok {
		log.Fatalf("missing proxy header")
	}
	log.Printf("%v", respHeaderProxy)

}
