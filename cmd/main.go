package main

import (
	"encoding/json"
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
	mux.HandleFunc("GET /hello-world", helloWorldHandler)

	mux.HandleFunc("POST /threads", h.CreateThreadHandler)
	mux.HandleFunc("GET /threads", h.ReadThreadListHandler)
	mux.HandleFunc("GET /threads/{uuid}", h.ReadThreadDetailHandler)

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

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(http.StatusOK)

	res := map[string]string{
		"msg": "Hello, World!",
	}

	json.NewEncoder(w).Encode(res)
}
