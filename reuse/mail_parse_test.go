package reuse_test

import (
	"io"
	"net/mail"
	"strings"
	"testing"
)

func TestParseMessage(t *testing.T) {

	// No blanks/tabs before headers (From, To, Subject)
	msg := `Date: Sat, 21 Sep 2024 16:11:30 -0300
From: Sender.Name<mail_sender@sender_server.info>
To: mail_receiver@receiver_server.info
Subject: Mail message to be parsed

This is a sample message to parse`

	reader := strings.NewReader(msg)
	m, err := mail.ReadMessage(reader)
	if err != nil {
		t.Fatalf("failed to read message: %v", err)
	}

	headers := m.Header
	if len(headers) == 0 {
		t.Fatalf("failed to parse message headers")
	}

	expectedHeaders := []struct {
		header string
		value  string
	}{
		{"Date", "Sat, 21 Sep 2024 16:11:30 -0300"},
		{"From", "Sender.Name<mail_sender@sender_server.info>"},
		{"To", "mail_receiver@receiver_server.info"},
		{"Subject", "Mail message to be parsed"},
	}

	for _, h := range expectedHeaders {
		if v := headers[h.header][0]; v != h.value {
			t.Errorf("%s: get %s, expect %s", h.header, v, h.value)
		}
	}

	bodyBytes, err := io.ReadAll(m.Body)
	if err != nil {
		t.Errorf("failes to read message body: %v", err)
	}

	body, expectedBody := string(bodyBytes), `This is a sample message to parse`
	if body != expectedBody {
		t.Errorf("message body: get\n%q\nexpect\n%q", body, expectedBody)
	}

	senderAddress, err := mail.ParseAddress(headers["From"][0])
	if err != nil {
		t.Errorf("failed to parse sender address: %v", err)
	} else {
		if actual, expected := senderAddress.Name, "Sender.Name"; actual != expected {
			t.Errorf("sender name:\nget\n%q\nexpect\n%q", actual, expected)
		}
		if actual, expected := senderAddress.Address, "mail_sender@sender_server.info"; actual != expected {
			t.Errorf("sender address:\nget\n%q\nexpect\n%q", actual, expected)
		}
	}

	receiverAddress, err := mail.ParseAddress(headers["To"][0])
	if err != nil {
		t.Errorf("failed to parse receiver address: %v", err)
	} else {
		if actual, expected := receiverAddress.Name, ""; actual != expected {
			t.Errorf("receiver name:\nget\n%q\nexpect\n%q", actual, expected)
		}
		if actual, expected := receiverAddress.Address, "mail_receiver@receiver_server.info"; actual != expected {
			t.Errorf("receiver address:\nget\n%q\nexpect\n%q", actual, expected)
		}
	}

}
