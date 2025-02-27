package main

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestApp(t *testing.T) {

	go RunServer()

	time.Sleep(100 * time.Millisecond)

	resp, err := http.Get("http://localhost:3001/debug/vars")
	if err != nil {
		t.Fatalf("failed to fetch server response: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	varsActual := Vars{}

	err = json.Unmarshal(respBody, &varsActual)
	if err != nil {
		t.Fatalf("failed to convert response body to struct: %v", err)
	}

	varsExpected := Vars{Counter: 42, Load: 42, Status: "running"}
	if !reflect.DeepEqual(varsActual, varsExpected) {
		t.Errorf("/debug/vars:\nget\n%v\nexpect\n%v\n", varsActual, varsExpected)
	}

}
