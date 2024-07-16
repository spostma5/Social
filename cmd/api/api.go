package api

import (
	"context"
	"log/slog"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	slog.Info("New API Server created", "Address", addr)

	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run(ctx *context.Context) error {
	slog.Info("Running server...")

	mux := http.NewServeMux()
	v1 := http.StripPrefix("/v1", mux)

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi hun\n"))
	})

	middleware := []Middleware{
		LogMiddleware,
	}

	mh := HandleMiddleware(v1, middleware...)

	server := http.Server{
		Addr:    s.addr,
		Handler: mh,
	}

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
