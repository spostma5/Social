package api

import (
	"log/slog"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func HandleMiddleware(h http.Handler, m ...Middleware) http.Handler {
	if len(m) < 1 {
		return h
	}

	wrapped := h

	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}

	return wrapped
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Handling request", "Method", r.Method, "Path", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
