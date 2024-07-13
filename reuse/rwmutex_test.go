// https://stackoverflow.com/a/19168242

package reuse_test

import (
	"fmt"
	"sync"
	"testing"
)

const (
	numRead  = 50
	numWrite = 5
)

type messageRW struct {
	message string
	lock    sync.RWMutex
}

func readMessage(c *messageRW, readerId int, chReaderId chan int, chMessage chan string, chDone chan bool) {
	c.lock.RLock()
	chReaderId <- readerId
	chMessage <- c.message
	c.lock.RUnlock()
	chDone <- true
}

func writeMessage(c *messageRW, writerId int, chWriterId chan int, m string, chDone chan bool) {
	c.lock.Lock()
	c.message = m
	chWriterId <- writerId
	c.lock.Unlock()
	chDone <- true
}

func TestRLock(t *testing.T) {

	c := messageRW{message: "initial", lock: sync.RWMutex{}}
	chReader := make(chan int)
	chWriter := make(chan int)
	chMessage := make(chan string)
	chDone := make(chan bool)
	var readsId, writesId []int
	var msgRead []string

	go func() {
		for i := 0; i < numRead; i++ {
			readMessage(&c, i, chReader, chMessage, chDone)
		}
	}()

	go func() {
		for i := 0; i < numWrite; i++ {
			writeMessage(&c, i, chWriter, fmt.Sprintf("message%02d", i), chDone)
		}
	}()

	go func() {
		for i := 0; i < numRead; i++ {
			readsId = append(readsId, <-chReader)
		}
		chDone <- true
	}()

	go func() {
		for i := 0; i < numWrite; i++ {
			writesId = append(writesId, <-chWriter)
		}
		chDone <- true
	}()

	go func() {
		for i := 0; i < numRead; i++ {
			msgRead = append(msgRead, <-chMessage)
		}
		chDone <- true
	}()

	for i := 0; i < numRead+numWrite+3; i++ {
		<-chDone
	}

	if len(readsId) != numRead {
		t.Errorf("number of reads: get %d, expect %d", len(readsId), numRead)
	}

}
