package reuse_test

import (
	"fmt"
	"testing"
)

type Message struct {
	messageID   int
	messageBody string
}

func (m *Message) Set(id int, body string) *Message {
	m.messageID = id
	m.messageBody = body
	return m
}

func TestPointerReceiverReturned(t *testing.T) {
	m1 := &Message{}
	m2 := m1.Set(1, "msg")
	if m1 != m2 {
		t.Errorf("should be equal")
	}
	a1 := fmt.Sprintf("%p", m1)
	a2 := fmt.Sprintf("%p", m2)
	if a1 != a2 {
		t.Errorf("should be equal")
	}
}
