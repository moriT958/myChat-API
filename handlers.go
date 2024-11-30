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
	// リクエストの処理
	var reqSchema CreateThreadRequest
	err := json.NewDecoder(r.Body).Decode(&reqSchema)
	if err != nil || reqSchema.Topic == "" {
		logger.Error("Invalid JSON format")
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// データベースへの保存処理
	var t Thread
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
	var t Thread
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
		t Thread
		p Post
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

// TODO: modelsを使用して、DB処理とリクエスト・レスポンスの処理を分ける
func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var reqSchema CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&reqSchema); err != nil {
		log.Println(err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	var thread_id int
	q := "SELECT id FROM threads WHERE uuid = $1;"
	row := db.QueryRow(q, reqSchema.ThreadUuid)
	if err := row.Scan(&thread_id); err != nil {
		log.Println(err)
		http.Error(w, "Failed scan db data", http.StatusInternalServerError)
		return
	}

	var res BasePostResponse
	var tmpTime time.Time
	q = `INSERT INTO posts (uuid, body, thread_id, created_at) VALUES ($1, $2, $3, now()) RETURNING uuid, body, created_at;`
	row = db.QueryRow(q, uuid.NewString(), reqSchema.Body, thread_id)
	if err := row.Scan(&res.Uuid, &res.Body, &tmpTime); err != nil {
		log.Println(err)
		http.Error(w, "Failed to scan db data", http.StatusInternalServerError)
		return
	}
	res.ThreadUuid = reqSchema.ThreadUuid
	res.CreatedAt = tmpTime.Format("2006-01-02 15:04:05")

	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func GetPostListHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	threadUuid := r.PathValue("threadUuid")

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

	var res GetPostListResponse
	res.Posts = make([]postOnThread, 0)

	sql := `SELECT posts.uuid, posts.body, posts.created_at 
			FROM posts JOIN threads ON posts.thread_id = threads.id 
			WHERE threads.uuid = $1
			OFFSET $2
			LIMIT $3;`
	rows, err := db.Query(sql, threadUuid, offset, limit)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get post data", http.StatusInternalServerError)
		return
	}
	var p postOnThread
	var tmpTime time.Time
	for rows.Next() {
		if err := rows.Scan(&p.Uuid, &p.Body, &tmpTime); err != nil {
			log.Println(err)
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
