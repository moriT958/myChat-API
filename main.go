package main

import (
	"database/sql"
	"log"
	"myChat-API2/internal/handler"
	"myChat-API2/internal/query"
	"myChat-API2/internal/repository"
	"myChat-API2/internal/service"
	"net/http"
	"os"
	"time"

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
	handlers := handler.NewHandlers(authService, chatService)
	hub := handler.NewHub()
	wsHandler := handler.NewWSHandler(hub, chatService)

	mux := http.NewServeMux()

	// thread functions
	mux.Handle("POST /threads", handlers.AuthMiddleware(handlers.CreateThreadHandler))
	mux.HandleFunc("GET /threads", handlers.GetThreadListHandler)
	mux.HandleFunc("GET /threads/{threadID}", handlers.ReadThreadDetailHandler)

	// get user info function
	mux.HandleFunc("GET /users/{userID}", handlers.GetUserHandler)

	// auth functions
	mux.HandleFunc("POST /signup", handlers.SignupHandler)
	mux.HandleFunc("POST /login", handlers.LoginHandler)

	// posting functions
	go wsHandler.Hub.Start()
	mux.Handle("/ws", handlers.AuthMiddleware(wsHandler.PostHandler))

	s := http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      mux,
		ReadTimeout:  time.Duration(10 * int64(time.Second)),
		WriteTimeout: time.Duration(600 * int64(time.Second)),
	}
	log.Printf("Server started at port %s!\n", s.Addr)
	log.Fatal(s.ListenAndServe())
}
