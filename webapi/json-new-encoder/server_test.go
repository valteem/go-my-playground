package main

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestHandlerWithJSONNewEncoder(t *testing.T) {

	p := &Person{
		GivenName:  "Some-Given-Name",
		FamilyName: "Some-Family0Name",
		Age:        42,
		Address: &Address{
			City:    "Some-City",
			Street:  "Some-Street",
			ZipCode: "12345",
		},
	}

	go Run(p)

	// Allow server some time to start properly
	time.Sleep(100 * time.Millisecond)

	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, "http://localhost:3001/person", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("failed to fetch server response: %v", err)
	}
	defer resp.Body.Close()

	pResp := &Person{}
	err = json.NewDecoder(resp.Body).Decode(pResp)
	if err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if !reflect.DeepEqual(pResp, p) {
		t.Errorf("decoding server response:\nget\n%v\nexpect\n%v\n", pResp, p)
	}

}
