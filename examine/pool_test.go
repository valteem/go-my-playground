package examine

import (
	"net/http"
	"net/url"
	"sync"
	"testing"

	"github.com/colega/zeropool"
)

func initHttpHeader() *http.Header {

	h := &http.Header{}

	h.Set("Host", "developer.mozilla.org")
	h.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:50.0) Gecko/20100101 Firefox/50.0")
	h.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,q=0.8")
	h.Set("Accept-Language", "en-US,en;q=0.5")
	h.Set("Accept-Encoding", "gzip, deflate, br")
	h.Set("Referer", "https://developer.mozilla.org/testpage.html")
	h.Set("Connection", "keep-alive")
	h.Set("Upgrade-Insecure-Requests", "1")
	h.Set("If-Modified-Since", "Tue, 02 Jul 2024 07:02:00 GMT")
	h.Set("If-None-Match", "c561c68d0ba92bbeb8b0fff2a9199f722e3a621a")
	h.Set("Cache-Control", "max-age=0")

	return h

}

func initSimpleHttpHeader() *http.Header {

	h := &http.Header{}

	h.Set("Host", "developer.mozilla.org")

	return h

}

func initHttpRequest() *http.Request {

	r := &http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme: "http",
			Host:   "localhost:33445",
			Path:   "/segment1/segment2/segment3",
		},
		Proto:  "HTTP/2",
		Header: *initHttpHeader(),
	}

	return r

}

func initSimpleHttpRequest() *http.Request {

	r := &http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme: "http",
			Host:   "localhost:33445",
			Path:   "/segment1/segment2/segment3",
		},
		Proto:  "HTTP/2",
		Header: *initSimpleHttpHeader(),
	}

	return r

}

func initHttpRequestNoHeader() *http.Request {

	r := &http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme: "http",
			Host:   "localhost:33445",
			Path:   "/segment1/segment2/segment3",
		},
		Proto: "HTTP/2",
	}

	return r

}

func initHttpRequestEmpty() *http.Request {

	return &http.Request{}

}
func BenchmarkPool(b *testing.B) {

	tests := []struct {
		name     string
		initFunc func() *http.Request
	}{
		{
			name:     "LargeHeader",
			initFunc: initHttpRequest,
		},
		{
			name:     "SmallHeader",
			initFunc: initSimpleHttpRequest,
		},
		{
			name:     "NoHeader",
			initFunc: initHttpRequestNoHeader,
		},
		{
			name:     "EmptyRequest",
			initFunc: initHttpRequestEmpty,
		},
	}

	for _, tc := range tests {
		b.Run(tc.name, func(b *testing.B) {
			p := &sync.Pool{New: func() any { return tc.initFunc() }}
			ch := make(chan *http.Request, b.N)
			wg := sync.WaitGroup{}
			wg.Add(b.N)

			for i := 0; i < b.N; i++ {
				// getting objects from pool
				go func() {
					r, _ := p.Get().(*http.Request)
					ch <- r
				}()
			}

			// putting objects back to the pool
			go func() {
				for r := range ch {
					p.Put(r)
					wg.Done()
				}
			}()
			wg.Wait()
		})
	}
}

func BenchmarkSyncPoolVsZeropool(b *testing.B) {

	b.Run("zeropool", func(b *testing.B) {
		p := zeropool.New(func() []int { return make([]int, 128) })

		s := p.Get()
		p.Put(s)

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			s := p.Get()
			p.Put(s)
		}
	})

	b.Run("sync.pool", func(b *testing.B) {
		p := &sync.Pool{New: func() any { return make([]int, 128) }}

		s := p.Get().([]int)
		p.Put(s)

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			s := p.Get().([]int)
			// When passing a value that is not a pointer to a function that accepts an interface,
			// the value needs to be placed on the heap, which means an additional allocation
			p.Put(s) // SA6002
		}
	})

}
