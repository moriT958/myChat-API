package server

import (
	"log"
	"myChat-API2/internal/service"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var _ http.Handler = (*ServerWS)(nil)

type ServerWS struct {
	Hub         *Hub
	ChatService service.IChatService
}

func NewServerWS(cs service.IChatService) *ServerWS {
	hub := NewHub()
	go hub.Start()

	return &ServerWS{
		Hub:         hub,
		ChatService: cs,
	}
}

func (ws *ServerWS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ws.chatHandler(w, r)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Client struct {
	conn *websocket.Conn
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan BaseChatResponse
	register   chan *Client
	unregister chan *Client
	mu         sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan BaseChatResponse),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (hub *Hub) Start() {
	for {
		select {
		case client := <-hub.register:
			hub.mu.Lock()
			hub.clients[client] = true
			hub.mu.Unlock()
		case client := <-hub.unregister:
			hub.mu.Lock()
			if _, ok := hub.clients[client]; ok {
				delete(hub.clients, client)
				client.conn.Close()
			}
			hub.mu.Unlock()
		case msg := <-hub.broadcast:
			hub.mu.RLock()
			for client := range hub.clients {
				err := client.conn.WriteJSON(msg)
				if err != nil {
					log.Println(err)
					client.conn.Close()
					hub.mu.RUnlock()
					hub.mu.Lock()
					delete(hub.clients, client)
					hub.mu.Unlock()
					hub.mu.RLock()
				}
			}
			hub.mu.RUnlock()
		}
	}
}
