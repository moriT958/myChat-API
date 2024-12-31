package server

import (
	"log"
	"net/http"
)

func (ws *ServerWS) ChatHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{conn: conn}
	ws.Hub.register <- client
	defer func() {
		ws.Hub.unregister <- client
	}()

	for {
		var newChat CreateChatRequest
		err := conn.ReadJSON(&newChat)
		if err != nil {
			log.Println(err, "Marker")
			break
		}
		log.Printf("Recieved: %+v\n", newChat)

		body := newChat.Body
		roomID := newChat.RoomID
		userID := getUserID(ctx)
		chat, err := ws.ChatService.CreateChat(ctx, body, roomID, userID)
		if err != nil {
			log.Println("Failed to post: ", err)
			break
		}

		res := BaseChatResponse{
			ID:        chat.ID,
			Body:      chat.Body,
			CreatedAt: chat.CreatedAt,
			RoomID:    chat.RoomID,
			UserID:    chat.UserID,
		}

		ws.Hub.broadcast <- res
	}
}
