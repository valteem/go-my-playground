package reuse_test

import (
	"fmt"
	"testing"
)

func TestSelect(t *testing.T) {
	var msg string
	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)
	chTerm := make(chan struct{}, 1)
	chRec := make(chan struct{}, 1)
	var s struct{}

	go func(msg *string) {
		for {
			select {
			case <-ch1:
				*msg = "1"
			case <-ch2:
				*msg = "2"
			case <-chTerm:
				chRec <- s
				return
			default:
				*msg = "0"
			}
		}
	}(&msg)

	ch1 <- s
	ch2 <- s // blocks if channel capacity is not set
	chTerm <- s

	for range chRec {
		fmt.Println(msg)
		break
	}

}
