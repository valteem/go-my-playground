// https://fale.io/blog/2018/08/31/a-small-http-debug-server-in-go

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"
)

var realUrl = "https://rbc.ru"

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8888", nil); err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	filePrefix := fmt.Sprintf("%s_%s", time.Now().Format("20221125-101830"), r.Method)

	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	if err := ioutil.WriteFile(filePrefix + "_request.txt", dump, 0644); err != nil {
		fmt.Println(err)
	}

	nr, err := http.NewRequest(r.Method, realUrl + r.URL.String(), r.Body)
	if err != nil {
		fmt.Println(err)
	}
	nr.Header = r.Header // for some unknown reason http.NewRequest sets method, URL and body, but skips headers
	response, err := http.DefaultClient.Do(nr)
	if err != nil {
		fmt.Println(err)
	}

	responseDump, err := httputil.DumpResponse(response, true)
	if err != nil {
		fmt.Println(responseDump)
	}
	if err := ioutil.WriteFile(filePrefix+"_response.txt", responseDump, 0644); err != nil {
		fmt.Println(err)
	}

	io.Copy(w, response.Body)
}