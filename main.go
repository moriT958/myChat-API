package main

import (
	"database/sql"
	"log"
	"myChat-API2/internal/handler"
	"myChat-API2/internal/query"
	"myChat-API2/internal/repository"
	"myChat-API2/internal/service"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file.")
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
	srv := handler.NewTodoServer(authService, chatService)

	srv.Run()
}
