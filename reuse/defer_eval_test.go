package reuse

import (
	"fmt"
	"testing"
	"time"
)

type profile struct {
	balance int
}

func (p *profile) getBalance() (balance int) {
	defer fmt.Println("no closure:", balance)
	defer func() {
		fmt.Println("with closure:", balance)
	}()
	return p.balance
}

func TestDeferEval(t *testing.T) {
	input := 1
	p := &profile{balance: input}
	if output := p.getBalance(); output != input {
		t.Errorf("getValue(): get %d, expect %d", output, input)
	}
}

// https://stackoverflow.com/a/55947800
func TestDeferStatementAndFunction(t *testing.T) {
	start := time.Now()
	time.Sleep(1 * time.Second)
	defer fmt.Println("Time elapsed - defer statement", time.Since(start)) // call to time.Since() is not deferred
	defer func() { fmt.Println("Time elapsed - defer function", time.Since(start)) }()
	time.Sleep(1 * time.Second)
}
