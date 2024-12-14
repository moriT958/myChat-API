package handler

import (
	"encoding/json"
	"log"
	"myChat-API/internal/model"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func (h *Handler) CreateThreadHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストの処理
	var reqSchema CreateThreadRequest
	err := json.NewDecoder(r.Body).Decode(&reqSchema)
	if err != nil || reqSchema.Topic == "" {
		log.Println(err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	username := GetUsername(r.Context())
	u, err := h.DAO.GetUserByUsername(username)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get User data.", http.StatusInternalServerError)
		return
	}

	t := model.Thread{
		Uuid:      uuid.New(),
		Topic:     reqSchema.Topic,
		CreatedAt: time.Now(),
		UserId:    u.Id,
	}

	if err := h.DAO.SaveThread(t); err != nil {
		log.Println(err)
		http.Error(w, "Failed to Save Thread", http.StatusInternalServerError)
		return
	}

	// レスポンスの作成処理
	res := BaseThreadResponse{
		Uuid:      t.Uuid.String(),
		Topic:     t.Topic,
		CreatedAt: t.CreatedAt.Format("2006-01-02 15:04:05"),
		UserId:    t.UserId,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf8")
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

	q, err := getQueryParams(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Parameters", http.StatusBadRequest)
		return
	}

	threads, err := h.DAO.GetThreads(q["limit"], q["offset"])
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed Get threads", http.StatusInternalServerError)
		return
	}

	// レスポンスの作成処理
	var res GetThreadListResponse
	for _, t := range threads {
		res.Threads = append(res.Threads, BaseThreadResponse{
			Uuid:      t.Uuid.String(),
			Topic:     t.Topic,
			CreatedAt: t.CreatedAt.Format("2006-01-02 15:04:05"),
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
	// リクエストの処理
	var uuid string = r.PathValue("uuid")

	thread, err := h.DAO.GetThreadByUuid(uuid)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get thread", http.StatusInternalServerError)
		return
	}

	posts, err := h.DAO.GetPostSByThreadId(thread.Id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get posts", http.StatusInternalServerError)
		return
	}

	// レスポンスの処理
	res := GetThreadDetailResponse{
		Uuid:      thread.Uuid.String(),
		Topic:     thread.Topic,
		CreatedAt: thread.CreatedAt.Format("2006-01-02 15:04:05"),
		Posts:     make([]PostOnThread, 0),
	}
	for _, p := range posts {
		res.Posts = append(res.Posts, PostOnThread{
			Uuid:      p.Uuid.String(),
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
