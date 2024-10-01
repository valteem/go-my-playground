// https://stackoverflow.com/a/8769095

package reuse_test

import (
	"testing"
)

type testCarEngine struct {
	engineType string
}

type testCar struct {
	engine *testCarEngine
}

func replaceEngineToHydrogen(tce *testCarEngine) {
	tce = &testCarEngine{"hydrogen"} // re-assign local copy of tce, no effect on outer struct
}

func replaceEngineToElectric(tce **testCarEngine) {
	*tce = &testCarEngine{"electric"} // reassign value of what tce is pointed to
}

func TestPointerCopy(t *testing.T) {

	c := testCar{&testCarEngine{"fossil"}}

	replaceEngineToHydrogen(c.engine)
	if actual, expect := c.engine.engineType, "fossil"; actual != expect {
		t.Errorf("get %q, expect %q", actual, expect)
	}

	replaceEngineToElectric(&c.engine)
	if actual, expect := c.engine.engineType, "electric"; actual != expect {
		t.Errorf("get %q, expect %q", actual, expect)
	}

}
