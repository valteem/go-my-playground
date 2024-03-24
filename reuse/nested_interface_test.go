// TODO: extend example to actual nested interfaces
package reuse_test

import (
	"fmt"
	"testing"
)

type Knocking interface {
	Knock()
}

type Sounding interface {
	Knocking // example of nested interface
	Voice()
	Noise()
	Drum()
}

type Drumming interface {
	Drum()
}

type Drums struct{}

func (m Drums) Drum() {
	fmt.Println("drumming ...")
}

func EvalDrum(v any) bool {
	_, ok := v.(Drumming)
	return ok
}

type Flute struct{}

func (f Flute) Voice() {
	fmt.Println("tootle-too ...")
}

type Hype struct{}

func (h Hype) Knock() {}
func (h Hype) Voice() {}
func (h Hype) Noise() {}
func (h Hype) Drum()  {}

func TestNestedInterfaces(t *testing.T) {
	tests := []struct {
		description string
		evalObj     any
		evalResult  bool
	}{
		{
			description: "drums type complies with Drumming interface",
			evalObj:     Drums{},
			evalResult:  true,
		},
		{
			description: "flute type does not comply with Drumming interface",
			evalObj:     Flute{},
			evalResult:  false,
		},
		{
			description: "hype type complies with Drumming interface",
			evalObj:     Hype{},
			evalResult:  true,
		},
	}
	for _, tst := range tests {
		if EvalDrum(tst.evalObj) != tst.evalResult {
			t.Errorf("%s, get %v, expect %v", tst.description, EvalDrum(tst.evalObj), tst.evalResult)
		}
	}
}
