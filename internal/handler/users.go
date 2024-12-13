package handler

import (
	"encoding/json"
	"log"
	"myChat-API/internal/schema"
	"net/http"
)

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
