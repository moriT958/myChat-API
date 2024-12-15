package handler

import (
	"log"
	"myChat-API/internal/domain"
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
		var inMsg InMessage
		err := conn.ReadJSON(&inMsg)
		if err != nil {
			log.Println(err, "Marker")
			break
		}
		log.Printf("Recieved: %+v\n", inMsg)

		// Save to DB
		threadId, err := h.DAO.GetThreadIdByUuid(inMsg.ThreadUuid)
		if err != nil {
			log.Println(err)
			break
		}

		post := domain.Post{
			Uuid:      uuid.New(),
			Body:      inMsg.Body,
			ThreadId:  threadId,
			CreatedAt: time.Now(),
		}
		if err := h.DAO.SavePost(post); err != nil {
			log.Println(err)
			break
		}

		outMsg := OutMessage{
			Uuid:       post.Uuid.String(),
			Body:       post.Body,
			ThreadUuid: inMsg.ThreadUuid,
			CreatedAt:  post.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		h.Hub.broadcast <- outMsg
	}
}
