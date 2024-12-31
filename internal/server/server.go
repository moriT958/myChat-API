package server

import (
	"fmt"
	"log"
	"myChat-API2/internal/config"
	"myChat-API2/internal/service"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type TodoServer struct {
	http.Server
	AuthService service.IAuthService
	ChatService service.IChatService
}

func NewTodoServer(
	as service.IAuthService,
	cs service.IChatService,
) *TodoServer {
	s := new(TodoServer)
	mux := http.NewServeMux()
	mux.Handle("POST /threads", s.AuthMiddleware(s.CreateThreadHandler))
	mux.Handle("GET /threads", http.HandlerFunc(s.GetThreadListHandler))
	mux.Handle("GET /threads/{threadID}", http.HandlerFunc(s.ReadThreadDetailHandler))

	// get user info function
	mux.HandleFunc("GET /users/{userID}", http.HandlerFunc(s.GetUserHandler))

	// auth functions
	mux.HandleFunc("POST /signup", http.HandlerFunc(s.SignupHandler))
	mux.HandleFunc("POST /login", http.HandlerFunc(s.LoginHandler))

	hub := NewHub()
	ws := NewServerWS(hub, cs)

	// posting functions
	go ws.Hub.Start()
	mux.Handle("/ws", s.AuthMiddleware(ws.PostHandler))

	s.Addr = config.Address()
	s.Handler = mux

	s.ReadTimeout = time.Duration(config.ReadTimeout() * int64(time.Second))
	s.WriteTimeout = time.Duration(config.WriteTimeout() * int64(time.Second))

	s.AuthService = as
	s.ChatService = cs

	return s
}

func (s *TodoServer) Run() {
	fmt.Printf("JustDoIt Version %s:\nServer starting at %s...\n", config.Version(), s.Addr)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	<-shutdown
	if err := s.Close(); err != nil {
		log.Fatalf("Server failed to shutdown gracefully: %v", err)
	}
}
