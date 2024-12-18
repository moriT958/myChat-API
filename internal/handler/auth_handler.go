package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handlers) SignupHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Username == "" || req.Password == "" {
		log.Println(err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	userID, err := h.AuthService.Signup(ctx, req.Username, req.Password)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to signup.", http.StatusInternalServerError)
		return
	}

	res := CreateUserResponse{
		Uuid: userID,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handlers) LoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := r.FormValue("username")
	password := r.FormValue("password")

	token, err := h.AuthService.Login(ctx, username, password)
	if err != nil {
		log.Println(err)
		http.Error(w, "username or password is incorrect.", http.StatusUnauthorized)
		return
	}

	res := GetTokenResponse{
		AccessToken: token,
		TokenType:   "Bearer",
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
