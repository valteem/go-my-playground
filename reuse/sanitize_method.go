package reuse

import (
	"net/http"
)

var methodMap = map[string]string {
	"GET":http.MethodGet, "get":http.MethodGet,
	"POST":http.MethodPost, "post":http.MethodPost,
}

func SanitizeMethod(m string) string {
	if a, ok := methodMap[m]; ok {
		return a
	}
	return "OTHER"
}