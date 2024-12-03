package config

import (
	"log"
	"log/slog"
	"os"
)

func NewLogger() *slog.Logger {
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	logWriter := slog.NewTextHandler(f, &slog.HandlerOptions{
		AddSource: true,
	})

	logger := *slog.New(logWriter)

	return &logger
}
