package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestBindingValidations(t *testing.T) {

	tests := []struct {
		input  *Person
		output string
	}{
		{
			input:  &Person{Name: "name123", Email: "some-invalid-email"},
			output: "\"Key: 'Person.Name' Error:Field validation for 'Name' failed on the 'alpha' tag\\nKey: 'Person.Email' Error:Field validation for 'Email' failed on the 'email' tag\"",
		},
	}

	go runServer(":3001")

	time.Sleep(100 * time.Millisecond) // allow server some time to start properly

	client := http.Client{}

	for _, tc := range tests {

		payload, err := json.Marshal(tc.input)
		if err != nil {
			t.Errorf("failed to encode %v input: %v", tc.input, err)
			continue
		}

		req, err := http.NewRequest(http.MethodPut, "http://localhost:3001/person", bytes.NewReader(payload))
		if err != nil {
			t.Errorf("failed to create a request: %v", err)
			continue
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("failed to fetch response: %v", err)
			continue
		}
		defer resp.Body.Close()

		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("failed to read response body for input %v: %v", tc.input, err)
			continue
		}

		output := string(msg)
		if output != tc.output {
			t.Errorf("response for %v input: get %q, expect %q", tc.input, output, tc.output)
		}

	}

}
