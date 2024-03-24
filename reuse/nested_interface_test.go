// TODO: extend example to actual nested interfaces
package reuse_test

import (
	"fmt"
	"testing"
)

type Clappable interface {
	Clap()
}

type Sounding interface {
	Clappable // example of nested interface
	Voice()
	Noise()
	Knock()
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

func TestNestedInterfaces(t *testing.T) {
	tests := []struct {
		description string
		evalObj     any
		evalResult  bool
	}{
		{
			description: "drums comply with Drumming interface",
			evalObj:     Drums{},
			evalResult:  true,
		},
		{
			description: "flute does not comply with Drumming interface",
			evalObj:     Flute{},
			evalResult:  false,
		},
	}
	for _, tst := range tests {
		if EvalDrum(tst.evalObj) != tst.evalResult {
			t.Errorf("%s, get %v, expect %v", tst.description, EvalDrum(tst.evalObj), tst.evalResult)
		}
	}
}
