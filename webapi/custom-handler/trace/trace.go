package trace

import (
	"fmt"
	"net/http"
	"runtime"
)

func WithStackTrace(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				const size = 64 << 10
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)] // for calling goroutine only
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "error:\n%v\n", err)
				w.Write(buf)
			}
		}()
		h.ServeHTTP(w, r)
	})
}
