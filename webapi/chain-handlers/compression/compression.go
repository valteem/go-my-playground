package compression

import (
	"compress/gzip"
	"net/http"
)

type gzipResponseWriter struct {
	http.ResponseWriter
	cw *gzip.Writer
}

func (g *gzipResponseWriter) Write(p []byte) (int, error) {
	return g.cw.Write(p)
}

func (g *gzipResponseWriter) Close() {
	g.cw.Close()
}

func newGzipResponseWriter(w http.ResponseWriter, r *http.Request) *gzipResponseWriter {
	return &gzipResponseWriter{
		ResponseWriter: w,
		cw:             gzip.NewWriter(w),
	}
}

func WithCompressionGzip(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gw := newGzipResponseWriter(w, r)
		defer gw.cw.Close() // https://stackoverflow.com/a/60923716
		h.ServeHTTP(gw, r)

	})
}
