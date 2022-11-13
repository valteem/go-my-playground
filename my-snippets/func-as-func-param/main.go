package main

import (
	"fmt"
)

type Headers map[string]string

type Request struct {
	Body string
	Headers *Headers
}

type Response struct{
	Body string
	Headers *Headers
}

type Handler func(Response, *Request)

type muxEntry struct {
	h Handler
	pattern string
}

type mux struct {
	es []muxEntry
}

func (mx *mux) HandleFunc(pattern string, handler Handler) {
	if handler == nil {
		panic(("http: nil handler"))
	}
	e := muxEntry{h: handler, pattern: pattern}
	mx.es = append(mx.es, e)
}

func hello(rw Response, rq *Request) {

	fmt.Printf("Hello: request header are ")
	i := 0
	for k, v := range *rq.Headers {
		fmt.Printf("%s:%s", k, v)
		if i < (len(*rq.Headers) - 1) {
			fmt.Print(", ")
		}
		i++
	}

	fmt.Printf(" and response headers are ")
	i = 0
	for k, v := range *rw.Headers {
		fmt.Printf("%s:%s", k, v)
		if i < (len(*rw.Headers) - 1) {
			fmt.Print(", ")		
		}
		i++

	}

}

func main() {

	hReq := make(Headers)
	hReq["hrq1"] = "vrq1"
	hReq["hrq2"] = "vrq2"

	hResp := make(Headers)
	hResp["hrsp1"] = "vrsp1"
	hResp["hrsp2"] = "vrsp2"

	req := Request{Body: "rqbody", Headers: &hReq}
	resp := Response{Body: "rspbody", Headers: &hResp}
	fmt.Println(req, resp)

	mx := mux{}
	mx.HandleFunc("hello", hello)
	fmt.Println(mx)
}