// https://sysdig.com/blog/golang-expvar-custom-metrics/

package main

import (
	"expvar"
	"net/http"
	"time"
)

type Vars struct {
	Counter int     `json:"app.counter"`
	Load    float64 `json:"app.load"`
	Status  string  `json:"app.status"`
}

func init() {
	go http.ListenAndServe(":3001", nil)
}

func RunServer() {

	var (
		counter = expvar.NewInt("app.counter")
		load    = expvar.NewFloat("app.load")
		status  = expvar.NewString("app.status")
	)

	for {
		counter.Set(42)
		load.Set(42.0)
		status.Set("running")
		time.Sleep(1 * time.Second)
	}

}
