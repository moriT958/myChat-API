package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

var logger slog.Logger

func main() {
	var err error

	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	logWriter := slog.NewTextHandler(f, &slog.HandlerOptions{
		AddSource: true,
	})
	logger = *slog.New(logWriter)

	dsn := os.Getenv("DATABASE_URL")
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello-world", helloWorldHandler)

	mux.HandleFunc("POST /threads", CreateThreadHandler)
	mux.HandleFunc("GET /threads", ReadThreadListHandler)
	mux.HandleFunc("GET /threads/{uuid}", ReadThreadDetailHandler)

	mux.HandleFunc("POST /posts", CreatePostHandler)
	mux.HandleFunc("GET /posts/{threadUuid}", GetPostListHandler)

	s := http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      mux,
		ReadTimeout:  time.Duration(10 * int64(time.Second)),
		WriteTimeout: time.Duration(600 * int64(time.Second)),
	}
	log.Printf("Server started at port %s!\n", s.Addr)
	log.Fatal(s.ListenAndServe())
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(http.StatusOK)

	res := map[string]string{
		"msg": "Hello, World!",
	}

	json.NewEncoder(w).Encode(res)
}
