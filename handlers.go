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

func ReadThreadListHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var offset, limit int

	q := r.URL.Query()

	if val, exists := q["offset"]; exists && len(val) > 0 {
		offset, err = strconv.Atoi(val[0])
		if err != nil {
			log.Println(err)
			offset = 0
			http.Error(w, "Invalid offset", http.StatusBadRequest)
		}
		if offset < 0 {
			offset *= -1
		}
	} else {
		offset = 0 // set default val.
	}

	if val, exists := q["limit"]; exists && len(val) > 0 {
		limit, err = strconv.Atoi(val[0])
		if err != nil {
			log.Println(err)
			limit = 3
			http.Error(w, "Invalid offset:", http.StatusBadRequest)
		}
		if limit < 0 {
			limit *= -1
		}
	} else {
		log.Println("check")
		limit = 3
	}

	// Calculate the number of data.
	start := offset
	end := offset + limit
	if start > len(ThreadData) {
		start = len(ThreadData)
	}
	if end > len(ThreadData) {
		end = len(ThreadData)
	}

	res := GetThreadListResponse{
		Threads: ThreadData[start:end],
	}

	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func ReadThreadDetailHandler(w http.ResponseWriter, r *http.Request) {
	var uuid string = r.PathValue("uuid")

	res := GetThreadDetailResponse{
		Uuid:      uuid,
		Topic:     "Test Thread Detail",
		CreatedAt: "2006-01-02 15:04:05",
		Posts: []BasePostResponse{
			{
				Uuid:       "1234",
				Body:       "test post 1",
				ThreadUuid: uuid,
				CreatedAt:  "2006-01-02 15:04:05",
			},
			{
				Uuid:       "abcd",
				Body:       "test post 2",
				ThreadUuid: uuid,
				CreatedAt:  "2006-01-02 15:04:05",
			},
		},
	}

	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var reqSchema CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&reqSchema); err != nil {
		log.Println(err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
	}

	res := BasePostResponse{
		Uuid:       uuid.NewString(),
		Body:       reqSchema.Body,
		ThreadUuid: reqSchema.ThreadUuid,
		CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
	}

	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func GetPostListHandler(w http.ResponseWriter, r *http.Request) {
	threadUuid := r.PathValue("threadUuid")
	var res GetPostListResponse
	res.Posts = []BasePostResponse{}
	for _, p := range PostData {
		if p.ThreadUuid == threadUuid {
			res.Posts = append(res.Posts, p)
		}
	}

	q := r.URL.Query()

	var err error
	var offset, limit int

	if val, exits := q["offset"]; exits && len(val) > 0 {
		offset, err = strconv.Atoi(val[0])
		if err != nil || offset < 0 {
			log.Println("Invalid offset param", err)
			http.Error(w, "Invalid offset params", http.StatusBadRequest)
			return
		}
	} else {
		offset = 0
	}

	if val, exists := q["limit"]; exists && len(val) > 0 {
		limit, err = strconv.Atoi(val[0])
		if err != nil || limit < 0 {
			log.Println("Invalid limit param", err)
			http.Error(w, "Invalid limit params", http.StatusBadRequest)
			return
		}
	} else {
		limit = 3
	}

	start := offset
	end := offset + limit
	if start > len(res.Posts) {
		start = len(res.Posts)
	}
	if end > len(res.Posts) {
		end = len(res.Posts)
	}
	res.Posts = res.Posts[start:end]

	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
