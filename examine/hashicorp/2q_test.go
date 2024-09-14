package hashicorp

import (
	"testing"

	"github.com/hashicorp/golang-lru/v2"
)

func Test2Q(t *testing.T) {

	c, err := lru.New2Q[int, int](3)
	if err != nil {
		t.Fatalf("error creating 2Q cache: %v", err)
	}

	c.Add(1, 1)
	c.Add(2, 2)
	v, ok := c.Get(1)
	if v != 1 || !ok {
		t.Errorf("Get(1): fetch (%d, %t), expect (1, true)", v, ok)
	}
	v, ok = c.Get(2)
	if v != 2 || !ok {
		t.Errorf("Get(2): fetch (%d, %t), expect (2, true)", v, ok)
	}

	c.Add(3, 3)
	c.Add(3, 33)
	v, ok = c.Get(3)
	if v != 33 || !ok {
		t.Errorf("Get(2): fetch (%d, %t), expect (33, true)", v, ok)
	}

	// Both Add(key, value) and Get(key) move the key/value pair from `recent` to `frequent` queue
	// after just one more call except initial,
	// so that `frequent` actually means 'referred to (Add()/Get()) more than just once'
}
