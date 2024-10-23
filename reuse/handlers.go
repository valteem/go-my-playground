// https://github.com/prometheus/prometheus/blob/7c7116fea8343795cae6da42960cacd0207a2af8/web/web.go#L117

package reuse

import (
	"log/slog"
	"net/http"
	"runtime"
)

func LogStackTraceOnHandlerPanic(h http.Handler, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				const size = 2 << 18
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)] // false - only for current goroutine
				logger.Error("panic serving request", "client", r.RemoteAddr, "url", r.URL, "error", err, "stack", buf)
				panic(err)
			}
		}()
		h.ServeHTTP(w, r)
	})
}
