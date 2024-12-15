package handler

import (
	"myChat-API/internal/service"
	"net/http"

	"github.com/gorilla/websocket"
)

type Handler struct {
	As  *service.AuthService
	Ts  *service.ThreadService
	Hub *Hub
	websocket.Upgrader
}

func New(as *service.AuthService, ts *service.ThreadService) *Handler {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	hub := NewHub()

	return &Handler{
		As:       as,
		Ts:       ts,
		Hub:      hub,
		Upgrader: upgrader,
	}
}
