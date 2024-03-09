package reuse_test

import (
	"testing"
)

type Broker interface {
	Ping() error
	Close() error
}

type TestBroker struct{}

func (b *TestBroker) Ping() error {
	return nil
}

func (b *TestBroker) Close() error {
	return nil
}

type AnotherBroker struct{}

func (b *AnotherBroker) Ping() error {
	return nil
}

func (b *AnotherBroker) Close() string {
	return "closed"
}

func TestBrokerImplementation(t *testing.T) {

	var _ Broker = (*TestBroker)(nil) // converting 'nil' to type '*TestBroker' (https://stackoverflow.com/a/60440959)

	// Does not compile - wrong implementation
	// var _ Broker = (*AnotherBroker)(nil)
}
