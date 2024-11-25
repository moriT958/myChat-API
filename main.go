package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf8")
		w.WriteHeader(http.StatusOK)

		res := map[string]string{
			"msg": "Hello, World!",
		}

		json.NewEncoder(w).Encode(res)
	})

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
	log.Fatal(s.ListenAndServe())
}
