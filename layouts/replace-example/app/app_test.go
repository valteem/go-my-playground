package main

import (
	"testing"

	"github.com/SomeFancyAccount/receiver"
	"github.com/SomeFancyAccount/sender"
)

func TestApp(t *testing.T) {

	_ = sender.NewSender("http://localhost:3001/add")
	_ = receiver.NewReceiver(":3001", "/add", nil)

}
