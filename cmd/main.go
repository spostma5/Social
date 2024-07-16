package main

import (
	"context"
	"log/slog"

	"github.com/spostma5/social/cmd/api"
)

func main() {
	server := api.NewAPIServer("172.105.24.66:8080")
	ctx := context.Background()
	if err := server.Run(&ctx); err != nil {
		slog.Error("Closing server with error", "err", err)
	}
}
