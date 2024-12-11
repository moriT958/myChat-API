package handler

import (
	"encoding/json"
	"log"
	"myChat-API/internal/model"
	"myChat-API/internal/schema"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストの処理
	var req schema.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Username == "" || req.Password == "" {
		log.Println(err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	u := model.User{
		Uuid:      uuid.New(),
		Username:  req.Username,
		Password:  req.Password,
		CreatedAt: time.Now(),
	}

	if err := h.DAO.SaveUser(u); err != nil {
		log.Println(err)
		http.Error(w, "Failed to Save User", http.StatusInternalServerError)
		return
	}

	// レスポンスの作成処理
	res := schema.CreateUserResponse{
		Uuid:     u.Uuid.String(),
		CreateAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(&res); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストの処理
	var username string = r.PathValue("username")

	user, err := h.DAO.GetUserByUsername(username)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	// レスポンスの処理
	res := schema.GetUserResponse{
		Uuid:      user.Uuid.String(),
		Username:  user.Username,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
