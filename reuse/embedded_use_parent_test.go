package reuse_test

import (
	"testing"
)

type embeddedObject struct {
	status    string
	setStatus func() string
}

func (e *embeddedObject) embedStatus() {
	e.status = e.setStatus()
}

type parentObject struct {
	status string
	e      *embeddedObject
}

func (p parentObject) translateStatus() string {
	return p.status
}

func (p parentObject) getEmbeddedStatus() string {
	return p.e.status
}

func TestEmbeddedCall(t *testing.T) {
	p := &parentObject{status: "created", e: &embeddedObject{}}
	p.e.setStatus = p.translateStatus

	p.e.embedStatus()
	if s := p.getEmbeddedStatus(); s != "created" {
		t.Errorf("expect \"created\", get %s", p.e.status)
	}
}
