package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflect"
	"sync"
	"testing"

	"github.com/SomeFancyAccount/receiver"
	"github.com/SomeFancyAccount/sender"
)

type Order struct {
	Article string `json:"article"`
	Qty     int    `json:"qty"`
}

func TestApp(t *testing.T) {

	ch := make(chan []byte)

	h := func(ch chan []byte) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			b, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			ch <- b
			w.WriteHeader(http.StatusOK)
		})
	}

	sender := sender.NewSender("http://localhost:3001/add")
	receiver := receiver.NewReceiver(":3001", "/add", h(ch))

	orderSend := Order{Article: "some article", Qty: 42}
	var orderReceive Order

	go receiver.Receive()

	var wg sync.WaitGroup
	wg.Add(1)
	/*
		go func(ch chan []byte, wg *sync.WaitGroup) {
			select {
			case b := <-ch:
				err := json.Unmarshal(b, &orderReceive)
				wg.Done()
				if err != nil {
					log.Fatalf("failed to unmarshal JSON %v: %v", b, err)
				}
				return
			}
		}(ch, &wg)
	*/

	go func(ch chan []byte, wg *sync.WaitGroup) {
		for b := range ch {
			err := json.Unmarshal(b, &orderReceive)
			wg.Done()
			if err != nil {
				log.Fatalf("failed to unmarshal JSON %v: %v", b, err)
			}
		}
	}(ch, &wg)

	sender.Send(orderSend)

	wg.Wait()

	if !reflect.DeepEqual(orderReceive, orderSend) {
		t.Errorf("receive\n%v\nexpect\n%v\n", orderReceive, orderSend)
	}

}
