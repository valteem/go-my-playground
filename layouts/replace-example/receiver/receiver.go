package receiver

import (
	"net/http"
)

type Receiver struct {
	port    string
	path    string
	handler http.Handler
}

func NewReceiver(port, path string, h http.Handler) *Receiver {
	return &Receiver{port: port, path: path, handler: h}
}

func (r *Receiver) Receive() {

	mux := http.NewServeMux()

	mux.Handle(r.path, r.handler)

	http.ListenAndServe(r.port, mux)

}
