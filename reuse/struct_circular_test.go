package reuse_test

import (
	"testing"

	"github.com/valteem/reuse"
)

func TestStructCircularRef(t *testing.T) {

	apples := reuse.NewFruit("apples")
	oranges := reuse.NewFruit("oranges")

	if apples.Veg != nil {
		t.Errorf("apples.Veg: expect nil, get %v", apples.Veg)
	}
	if oranges.Veg != nil {
		t.Errorf("oranges.Veg: expect nil, get %v", oranges.Veg)
	}

	onions := reuse.NewVeg("onions")
	garlics := reuse.NewVeg("garlics")

	if onions.Fruit != nil {
		t.Errorf("onions.Fruit: expect nil, get %v", onions.Fruit)
	}
	if garlics.Fruit != nil {
		t.Errorf("garlics.Fruit: expect nil, get %v", garlics.Fruit)
	}

	apples.Tweak(onions)
	oranges.Tweak(garlics)

	if actual, expected := apples.Veg.Name, "onions"; actual != expected {
		t.Errorf("apples.Veg.Name: get %q, expect %q", actual, expected)
	}
	if actual, expected := oranges.Veg.Name, "garlics"; actual != expected {
		t.Errorf("oranges.Veg.Name: get %q, expect %q", actual, expected)
	}

	onions.Tweak(oranges)
	garlics.Tweak(apples)

	if actual, expected := apples.Veg.Fruit, oranges; actual != expected {
		t.Errorf("apples-to-oranges chain: get %v, expect %v", actual, expected)
	}
	if actual, expected := oranges.Veg.Fruit, apples; actual != expected {
		t.Errorf("apples-to-oranges chain: get %v, expect %v", actual, expected)
	}

}
