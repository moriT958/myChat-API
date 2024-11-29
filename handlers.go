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
	err := json.NewDecoder(r.Body).Decode(&reqSchema)
	if err != nil || reqSchema.Topic == "" {
		logger.Error("Invalid JSON format")
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	var res BaseThreadResponse
	var createdAtTime time.Time

	q := `INSERT INTO threads (uuid, topic, created_at) VALUES ($1, $2, now()) RETURNING uuid, topic, created_at;`
	row := db.QueryRow(q, uuid.New().String(), reqSchema.Topic)
	if err := row.Scan(&res.Uuid, &res.Topic, &createdAtTime); err != nil {
		log.Println(err)
		http.Error(w, "Failed to scan db data", http.StatusInternalServerError)
		return
	}
	res.CreatedAt = createdAtTime.Format("2006-01-02 15:04:05")

	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(&res); err != nil {
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
		limit = 3
	}

	var res GetThreadListResponse
	sql := `SELECT uuid, topic, created_at FROM threads LIMIT $1 OFFSET $2;`
	rows, err := db.Query(sql, limit, offset)
	if err != nil {
		logger.Error("Failed to get threads")
		http.Error(w, "Failed to get threads", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	res.Threads = make([]BaseThreadResponse, 0)
	var tmpTime time.Time
	for rows.Next() {
		var t BaseThreadResponse
		if err := rows.Scan(&t.Uuid, &t.Topic, &tmpTime); err != nil {
			log.Println(err)
			http.Error(w, "Failed to create response", http.StatusInternalServerError)
			return
		}
		t.CreatedAt = tmpTime.Format("2006-01-02 15:04:05")
		res.Threads = append(res.Threads, t)
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

	var res GetThreadDetailResponse
	var tmpTime time.Time
	var thread_id int
	q := `SELECT id, uuid, topic, created_at FROM threads WHERE uuid = $1;`
	row := db.QueryRow(q, uuid)
	if err := row.Scan(&thread_id, &res.Uuid, &res.Topic, &tmpTime); err != nil {
		log.Println(err)
		http.Error(w, "Failed to scan db data", http.StatusInternalServerError)
		return
	}
	res.CreatedAt = tmpTime.Format("2006-01-02 15:04:05")

	q = `SELECT uuid, body, created_at FROM posts WHERE thread_id = $1;`
	rows, err := db.Query(q, thread_id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get db data", http.StatusInternalServerError)
		return
	}
	res.Posts = make([]postsOnThread, 0)
	for rows.Next() {
		var p postsOnThread
		if err := rows.Scan(&p.Uuid, &p.Body, &tmpTime); err != nil {
			log.Println("failed to get posts: ", err)
			http.Error(w, "Failed to scan db data", http.StatusInternalServerError)
			return
		}
		p.CreatedAt = tmpTime.Format("2006-01-02 15:04:05")
		res.Posts = append(res.Posts, p)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// TODO: Post機能のDB接続
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
