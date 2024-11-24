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

	s := http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      mux,
		ReadTimeout:  time.Duration(10 * int64(time.Second)),
		WriteTimeout: time.Duration(600 * int64(time.Second)),
	}
	log.Fatal(s.ListenAndServe())
}
