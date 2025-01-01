package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"log/slog"
	"myChat-API2/internal/config"
	"myChat-API2/internal/dependency"
	"os"
)

func init() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := config.Load("config.json"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	s := dependency.InitServer()
	s.Run()
}
