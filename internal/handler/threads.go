package handler

import (
	"encoding/json"
	"log"
	"myChat-API/internal/domain"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func (h *Handler) CreateThreadHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// リクエストの処理
	var reqSchema CreateThreadRequest
	err := json.NewDecoder(r.Body).Decode(&reqSchema)
	if err != nil || reqSchema.Topic == "" {
		log.Println(err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Create Thread.
	username := GetUsername(r.Context())
	threadUuid, err := h.Ts.CreateThread(ctx, reqSchema.Topic, username)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to your user info.", http.StatusInternalServerError)
		return
	}

	// レスポンスの作成処理
	res := CreateThreadResponses{
		Uuid: threadUuid,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(&res); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func getQueryParams(r *http.Request) (map[string]int, error) {
	res := map[string]int{
		"offset": 0,
		"limit":  3,
	}

	qParams := r.URL.Query()

	if val, exists := qParams["offset"]; exists && len(val) > 0 {
		offset, err := strconv.Atoi(val[0])
		if err != nil {
			return res, err
		}
		res["offset"] = offset
	}

	if val, exists := qParams["limit"]; exists && len(val) > 0 {
		limit, err := strconv.Atoi(val[0])
		if err != nil {
			return res, err
		}
		res["limit"] = limit
	}

	return res, nil
}

func (h *Handler) ReadThreadListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q, err := getQueryParams(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Parameters", http.StatusBadRequest)
		return
	}

	threads, err := h.Ts.ReadAllThreads(ctx, q["limit"], q["offset"])
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to read threads.", http.StatusInternalServerError)
		return
	}

	// レスポンスの作成処理
	var res GetThreadListResponse
	for _, t := range threads {
		createdAt, err := h.Ts.Tr.GetCreatedAt(ctx, t.Uuid)
		if err != nil {
			log.Println(err)
			http.Error(w, "Failed to get createdAt.", http.StatusInternalServerError)
			return
		}
		res.Threads = append(res.Threads, BaseThreadResponse{
			Uuid:      t.Uuid,
			Topic:     t.Topic,
			CreatedAt: createdAt,
			UserId:    t.UserId,
		})
	}

	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *Handler) ReadThreadDetailHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// リクエストの処理
	var threadUuid string = r.PathValue("uuid")

	thread, posts, err := h.Ts.ReadThreadDetail(ctx, threadUuid)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to read thread.", http.StatusInternalServerError)
		return
	}

	// レスポンスの処理
	createdAt, err := h.Ts.Tr.GetCreatedAt(ctx, thread.Uuid)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get createdAt.", http.StatusInternalServerError)
		return
	}
	res := GetThreadDetailResponse{
		Uuid:      thread.Uuid,
		Topic:     thread.Topic,
		CreatedAt: createdAt,
		Posts:     make([]PostOnThread, len(posts)),
	}
	for _, p := range posts {
		createdAt, err := h.Ps.GetCreatedAt(ctx, thread.Uuid)
		if err != nil {
			log.Println(err)
			http.Error(w, "Failed to get createdAt.", http.StatusInternalServerError)
			return
		}
		res.Posts = append(res.Posts, PostOnThread{
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
