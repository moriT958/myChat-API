package main

import (
	"encoding/json"
	"log"
	"myChat-API/model"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

func CreateThreadHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストの処理
	var reqSchema CreateThreadRequest
	err := json.NewDecoder(r.Body).Decode(&reqSchema)
	if err != nil || reqSchema.Topic == "" {
		logger.Error("Invalid JSON format")
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// データベースへの保存処理
	var t model.Thread
	q := `INSERT INTO threads (uuid, topic, created_at) VALUES ($1, $2, now()) RETURNING uuid, topic, created_at;`
	row := db.QueryRow(q, uuid.NewString(), reqSchema.Topic)
	if err := row.Scan(&t.Uuid, &t.Topic, &t.CreatedAt); err != nil {
		log.Println(err)
		http.Error(w, "Failed to scan db data", http.StatusInternalServerError)
		return
	}

	// レスポンスの作成処理
	res := BaseThreadResponse{
		Uuid:      t.Uuid,
		Topic:     t.Topic,
		CreatedAt: t.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(&res); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func ReadThreadListHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	// クエリパラメータの処理
	var (
		offset = 0
		limit  = 3
	)
	qParams := r.URL.Query()
	if val, exists := qParams["offset"]; exists && len(val) > 0 {
		offset, err = strconv.Atoi(val[0])
		if err != nil {
			log.Println(err)
			http.Error(w, "Invalid offset", http.StatusBadRequest)
		}
	}
	if val, exists := qParams["limit"]; exists && len(val) > 0 {
		limit, err = strconv.Atoi(val[0])
		if err != nil {
			log.Println(err)
			http.Error(w, "Invalid offset:", http.StatusBadRequest)
		}
	}

	// データベースの格納処理
	var t model.Thread
	sql := `SELECT uuid, topic, created_at FROM threads LIMIT $1 OFFSET $2;`
	rows, err := db.Query(sql, limit, offset)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get threads", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// レスポンスの作成処理
	res := GetThreadListResponse{
		Threads: make([]BaseThreadResponse, 0),
	}
	for rows.Next() {
		if err := rows.Scan(&t.Uuid, &t.Topic, &t.CreatedAt); err != nil {
			log.Println(err)
			http.Error(w, "Failed to create response", http.StatusInternalServerError)
			return
		}
		res.Threads = append(res.Threads, BaseThreadResponse{
			t.Uuid,
			t.Topic,
			t.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func ReadThreadDetailHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストの処理
	var uuid string = r.PathValue("uuid")

	// データベースの処理
	var (
		t model.Thread
		p model.Post
	)
	var thread_id int

	q := `SELECT id, uuid, topic, created_at FROM threads WHERE uuid = $1;`
	row := db.QueryRow(q, uuid)
	if err := row.Scan(&thread_id, &t.Uuid, &t.Topic, &t.CreatedAt); err != nil {
		log.Println(err)
		http.Error(w, "Failed to scan db data", http.StatusInternalServerError)
		return
	}

	q = `SELECT uuid, body, created_at FROM posts WHERE thread_id = $1;`
	rows, err := db.Query(q, thread_id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get db data", http.StatusInternalServerError)
		return
	}

	// レスポンスの処理
	res := GetThreadDetailResponse{
		Uuid:      t.Uuid,
		Topic:     t.Topic,
		CreatedAt: t.CreatedAt.Format("2006-01-02 15:04:05"),
		Posts:     make([]postOnThread, 0),
	}
	for rows.Next() {
		if err := rows.Scan(&p.Uuid, &p.Body, &p.CreatedAt); err != nil {
			log.Println("failed to get posts: ", err)
			http.Error(w, "Failed to scan db data", http.StatusInternalServerError)
			return
		}
		res.Posts = append(res.Posts, postOnThread{
			Uuid:      p.Uuid,
			Body:      p.Body,
			CreatedAt: p.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
