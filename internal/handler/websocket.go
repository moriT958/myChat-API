package handler

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan OutMessage
	register   chan *Client
	unregister chan *Client
	mu         sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan OutMessage),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (m *Hub) Start() {
	for {
		select {
		case client := <-m.register:
			m.mu.Lock()
			m.clients[client] = true
			m.mu.Unlock()
		case client := <-m.unregister:
			m.mu.Lock()
			if _, ok := m.clients[client]; ok {
				delete(m.clients, client)
				client.conn.Close()
			}
			m.mu.Unlock()
		case msg := <-m.broadcast:
			m.mu.RLock()
			for client := range m.clients {
				err := client.conn.WriteJSON(msg)
				if err != nil {
					log.Println(err)
					client.conn.Close()
					m.mu.RUnlock()
					m.mu.Lock()
					delete(m.clients, client)
					m.mu.Unlock()
					m.mu.RLock()
				}
			}
			m.mu.RUnlock()
		}
	}
}
