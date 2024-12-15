package handler

import (
	"log"
	"myChat-API/internal/domain"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (h *Handler) PostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
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
		threadUuid, err := uuid.Parse(inMsg.ThreadUuid)
		if err != nil {
			log.Println(err)
			http.Error(w, "Failed to parse uuid", http.StatusBadRequest)
			return
		}
		threadId, err := h.Queries.GetThreadIdByUuid(ctx, threadUuid)
		if err != nil {
			log.Println(err)
			break
		}
		t, err := h.

		post := domain.Post{
			Uuid:      uuid.New(),
			Body:      inMsg.Body,
			ThreadId:  int(threadId),
			CreatedAt: time.Now(),
		}
		if err := h.Queries.CreatePost(ctx, post); err != nil {
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
