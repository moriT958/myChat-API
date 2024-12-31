package main

import (
	"database/sql"
	"log"
	"log/slog"
	"myChat-API2/internal/config"
	"myChat-API2/internal/query"
	"myChat-API2/internal/repository"
	"myChat-API2/internal/server"
	"myChat-API2/internal/service"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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

	dsn := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Init Dependency
	queries := query.New(db)
	userRepo := repository.NewUserRepository(queries)
	threadRepo := repository.NewThreadRepository(queries)
	postRepo := repository.NewPostRepository(queries)
	authService := service.NewAuthService(userRepo)
	chatService := service.NewChatService(threadRepo, postRepo, userRepo)
	srv := server.NewTodoServer(authService, chatService)

	srv.Run()
}
