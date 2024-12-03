package handler

import (
	"log"
	"myChat-API/internal/model"
	"myChat-API/internal/schema"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (h *Handler) PostHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := h.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{conn: conn}
	h.Hub.register <- client
	defer func() {
		h.Hub.unregister <- client
	}()

	for {
		var inMsg schema.InMessage
		err := conn.ReadJSON(&inMsg)
		if err != nil {
			log.Println(err, "Marker")
			break
		}
		log.Printf("Recieved: %+v\n", inMsg)

		// Save to DB
		var thread_id int
		q := `SELECT id FROM threads WHERE uuid = $1;`
		row := h.Data.DB.QueryRow(q, inMsg.ThreadUuid)
		if err := row.Scan(&thread_id); err != nil {
			log.Println(err)
			break
		}

		post := model.Post{
			Uuid:      uuid.NewString(),
			Body:      inMsg.Body,
			ThreadId:  thread_id,
			CreatedAt: time.Now(),
		}

		q = `INSERT INTO posts (uuid, body, thread_id, created_at) VALUES ($1, $2, $3, $4);`
		if _, err := h.Data.DB.Exec(q, post.Uuid, post.Body, post.ThreadId, post.CreatedAt); err != nil {
			log.Println(err)
			http.Error(w, "Failed to scan db data", http.StatusInternalServerError)
			return
		}

		outMsg := schema.OutMessage{
			Uuid:       post.Uuid,
			Body:       post.Body,
			ThreadUuid: inMsg.ThreadUuid,
			CreatedAt:  post.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		h.Hub.broadcast <- outMsg
	}
}
