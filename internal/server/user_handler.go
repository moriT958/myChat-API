package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *TodoServer) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var userID string = r.PathValue("userID")

	user, createdAt, err := h.AuthService.SeeUserDetail(ctx, userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to read user data.", http.StatusInternalServerError)
		return
	}

	// Generate response.
	res := GetUserResponse{
		Uuid:      user.ID,
		Username:  user.Name,
		CreatedAt: createdAt,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
