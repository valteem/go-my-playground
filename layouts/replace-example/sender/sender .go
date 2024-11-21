package sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Sender struct {
	URI string
}

func NewSender(u string) *Sender {
	return &Sender{URI: u}
}

func (s *Sender) Send(payload any) error {

	json, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload %v to JSON: %v", payload, err)
	}

	var b bytes.Buffer
	b.Write(json)

	req, err := http.NewRequest(http.MethodPost, s.URI, &b)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}

	return nil

}
