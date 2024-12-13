package main

import (
	"log"
	"myChat-API/internal/dao"
	"myChat-API/internal/handler"
	"net/http"
	"os"
	"time"
)

func main() {

	dsn := os.Getenv("DATABASE_URL")
	d, err := dao.New(dsn)
	if err != nil {
		log.Fatal(err)
	}

	h := handler.New(d)

	mux := http.NewServeMux()

	// thread functions
	mux.Handle("POST /threads", h.AuthMiddleware(h.CreateThreadHandler))
	mux.HandleFunc("GET /threads", h.ReadThreadListHandler)
	mux.HandleFunc("GET /threads/{uuid}", h.ReadThreadDetailHandler)

	// get user info function
	mux.HandleFunc("GET /users/{username}", h.GetUserHandler)

	// auth functions
	mux.HandleFunc("POST /signup", h.SignupHandler)
	mux.HandleFunc("POST /login", h.LoginHandler)

	// posting functions
	go h.Hub.Start()
	mux.HandleFunc("/ws", h.PostHandler)

	s := http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      mux,
		ReadTimeout:  time.Duration(10 * int64(time.Second)),
		WriteTimeout: time.Duration(600 * int64(time.Second)),
	}
	log.Printf("Server started at port %s!\n", s.Addr)
	log.Fatal(s.ListenAndServe())
}
