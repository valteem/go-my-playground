package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestResponse(t *testing.T) {

	go main()

	requestURL := "http://localhost:3000//static/a.html"
	req, _ := http.NewRequest(http.MethodGet, requestURL, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("cannot execute request")
		os.Exit(1)
	}
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("cannot read response body")
		os.Exit(1)
	}
	fmt.Printf("response body:\n %s\n", respBody)
}