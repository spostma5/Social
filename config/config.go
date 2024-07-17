package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

var Envs Config

type Config struct {
	Address string
	Port    string
}

func InitEnvs() {
	err := godotenv.Load()
	if err != nil {
		slog.Debug("Error loading .env file", "err", err)
	}

	Envs.Address = getEnv("HOST_ADDRESS", "0.0.0.0")
	Envs.Port = getEnv("HOST_PORT", "8080")
}

func getEnv(key string, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}

	return fallback
}
