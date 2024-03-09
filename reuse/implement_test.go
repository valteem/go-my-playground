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

type YetAnotherBrocker struct{}

func (b YetAnotherBrocker) Ping() error {
	return nil
}

func (b YetAnotherBrocker) Close() error {
	return nil
}

func TestBrokerImplementation(t *testing.T) {

	var _ Broker = (*TestBroker)(nil) // converting 'nil' to type '*TestBroker' (https://stackoverflow.com/a/60440959)

	var s struct{}
	var _ Broker = (YetAnotherBrocker)(s) // test implementation with 'value' receiver

	// Does not compile - wrong implementation
	// var _ Broker = (*AnotherBroker)(nil)
}
