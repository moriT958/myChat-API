package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func CreateThreadHandler(w http.ResponseWriter, r *http.Request) {
	var reqSchema CreateThreadRequest
	if err := json.NewDecoder(r.Body).Decode(&reqSchema); err != nil {
		log.Println(err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	res := BaseThreadResponse{
		Uuid:      uuid.New().String(),
		Topic:     reqSchema.Topic,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Println("encoding error: ", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func ReadThreadList(w http.ResponseWriter, r *http.Request) {
	var err error

	// Query parameters
	var offset int
	var limit int

	query := r.URL.Query()

	if val := query.Get("offset"); val != "" {
		offset, err = strconv.Atoi(val)
		if err != nil {
			log.Println(err)
			http.Error(w, "Invalid offset value", http.StatusBadRequest)
			return
		}
		if offset < 0 {
			offset = 0
		}
	}

	if val := query.Get("limit"); val != "" {
		limit, err = strconv.Atoi(val)
		if err != nil {
			log.Println(err)
			http.Error(w, "Invalid limit value", http.StatusBadRequest)
			return
		}
		if limit < 1 {
			limit = len(Data.Threads)
		}
	}

	// Calculate the number of data.
	start := offset
	end := offset + limit
	if start > len(Data.Threads) {
		start = len(Data.Threads)
	}
	if end > len(Data.Threads) {
		end = len(Data.Threads)
	}

	res := GetThreadListResponse{
		Threads: Data.Threads[start:end],
	}

	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
