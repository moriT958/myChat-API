package handler

import (
	"myChat-API/internal/dao"
	"net/http"

	"github.com/gorilla/websocket"
)

type Handler struct {
	DAO *dao.DAO
	Hub *Hub
	websocket.Upgrader
}

func New(d *dao.DAO) *Handler {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	hub := NewHub()

	return &Handler{
		DAO:      d,
		Hub:      hub,
		Upgrader: upgrader,
	}
}
