package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (h *TodoServer) CreateThreadHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parse Request
	var reqNewThread CreateThreadRequest
	err := json.NewDecoder(r.Body).Decode(&reqNewThread)
	if err != nil || reqNewThread.Topic == "" {
		log.Println(err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Create Thread.
	userID := getUserID(ctx)
	threadID, err := h.ChatService.CreateThread(ctx, reqNewThread.Topic, userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to post thread.", http.StatusInternalServerError)
		return
	}

	// Generate response.
	res := CreateThreadResponses{
		Uuid: threadID,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(&res); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *TodoServer) GetThreadListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q, err := getQueryParams(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Parameters", http.StatusBadRequest)
		return
	}

	threads, err := h.ChatService.ShowAllThreads(ctx, q["limit"], q["offset"])
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to read threads.", http.StatusInternalServerError)
		return
	}

	// Generate response.
	var res GetThreadListResponse
	res.Threads = make([]BaseThreadResponse, len(threads))
	for i, t := range threads {
		res.Threads[i] = BaseThreadResponse{
			Uuid:      t.ID,
			Topic:     t.Topic,
			CreatedAt: t.CreatedAt,
			UserUuid:  t.UserID,
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
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

func (h *TodoServer) ReadThreadDetailHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var threadID string = r.PathValue("threadID")

	thread, posts, err := h.ChatService.SeeThreadDetail(ctx, threadID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to read thread.", http.StatusInternalServerError)
		return
	}

	// Generate response.
	res := GetThreadDetailResponse{
		Uuid:      thread.ID,
		Topic:     thread.Topic,
		CreatedAt: thread.CreatedAt,
		Posts:     make([]BasePostResponse, len(posts)),
		UserUuid:  thread.UserID,
	}

	for i, p := range posts {
		res.Posts[i] = BasePostResponse{
			Uuid:       p.ID,
			Body:       p.Body,
			CreatedAt:  p.CreatedAt,
			ThreadUuid: p.ThreadID,
			UserUuid:   p.UserID,
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
