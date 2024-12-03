package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// TODO: modelsを使用して、DB処理とリクエスト・レスポンスの処理を分ける
func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var reqSchema CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&reqSchema); err != nil {
		log.Println(err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	var thread_id int
	q := "SELECT id FROM threads WHERE uuid = $1;"
	row := db.QueryRow(q, reqSchema.ThreadUuid)
	if err := row.Scan(&thread_id); err != nil {
		log.Println(err)
		http.Error(w, "Failed scan db data", http.StatusInternalServerError)
		return
	}

	var res BasePostResponse
	var tmpTime time.Time
	q = `INSERT INTO posts (uuid, body, thread_id, created_at) VALUES ($1, $2, $3, now()) RETURNING uuid, body, created_at;`
	row = db.QueryRow(q, uuid.NewString(), reqSchema.Body, thread_id)
	if err := row.Scan(&res.Uuid, &res.Body, &tmpTime); err != nil {
		log.Println(err)
		http.Error(w, "Failed to scan db data", http.StatusInternalServerError)
		return
	}
	res.ThreadUuid = reqSchema.ThreadUuid
	res.CreatedAt = tmpTime.Format("2006-01-02 15:04:05")

	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func GetPostListHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	threadUuid := r.PathValue("threadUuid")

	// クエリパラメータの処理
	var (
		offset = 0
		limit  = 3
	)
	qParams := r.URL.Query()
	if val, exists := qParams["offset"]; exists && len(val) > 0 {
		offset, err = strconv.Atoi(val[0])
		if err != nil {
			log.Println(err)
			http.Error(w, "Invalid offset", http.StatusBadRequest)
		}
	}
	if val, exists := qParams["limit"]; exists && len(val) > 0 {
		limit, err = strconv.Atoi(val[0])
		if err != nil {
			log.Println(err)
			http.Error(w, "Invalid offset:", http.StatusBadRequest)
		}
	}

	var res GetPostListResponse
	res.Posts = make([]postOnThread, 0)

	sql := `SELECT posts.uuid, posts.body, posts.created_at 
			FROM posts JOIN threads ON posts.thread_id = threads.id 
			WHERE threads.uuid = $1
			OFFSET $2
			LIMIT $3;`
	rows, err := db.Query(sql, threadUuid, offset, limit)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get post data", http.StatusInternalServerError)
		return
	}
	var p postOnThread
	var tmpTime time.Time
	for rows.Next() {
		if err := rows.Scan(&p.Uuid, &p.Body, &tmpTime); err != nil {
			log.Println(err)
			http.Error(w, "Failed to scan db data", http.StatusInternalServerError)
			return
		}
		p.CreatedAt = tmpTime.Format("2006-01-02 15:04:05")
		res.Posts = append(res.Posts, p)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

type InMessage struct {
	Body       string `json:"body"`
	ThreadUuid string `json:"threadUuid"`
}

type OutMessage struct {
	Uuid       string `json:"uuid"`
	Body       string `json:"body"`
	ThreadUuid string `json:"threadUuid"`
	CreatedAt  string `json:"createdAt"`
}

type Client struct {
	conn *websocket.Conn
}

type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan OutMessage
	register   chan *Client
	unregister chan *Client
	mu         sync.Mutex
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan OutMessage),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (m *ClientManager) Start() {
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
			m.mu.Lock()
			for client := range m.clients {
				err := client.conn.WriteJSON(msg)
				if err != nil {
					log.Println(err)
					client.conn.Close()
					delete(m.clients, client)
				}
			}
			m.mu.Unlock()
		}
	}
}

var manager = NewClientManager()

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func wsPostHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{conn: conn}
	manager.register <- client
	defer func() {
		manager.unregister <- client
	}()

	for {
		var inMsg InMessage
		err := conn.ReadJSON(&inMsg)
		if err != nil {
			log.Println(err, "Marker")
			break
		}

		log.Printf("Recieved: %+v\n", inMsg)

		outMsg := OutMessage{
			Uuid:       uuid.NewString(),
			Body:       inMsg.Body,
			ThreadUuid: inMsg.ThreadUuid,
			CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
		}
		manager.broadcast <- outMsg
	}
}
