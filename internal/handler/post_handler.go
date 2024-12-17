package handler

import (
	"log"
	"net/http"
)

func (ws *WSHandler) PostHandler(w http.ResponseWriter, r *http.Request) {
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
		var newPost CreatePostRequest
		err := conn.ReadJSON(&newPost)
		if err != nil {
			log.Println(err, "Marker")
			break
		}
		log.Printf("Recieved: %+v\n", newPost)

		body := newPost.Body
		threadID := newPost.ThreadUuid
		userID := getUserID(ctx)
		post, err := ws.ChatService.CreatePost(ctx, body, threadID, userID)
		if err != nil {
			log.Println("Failed to post: ", err)
			break
		}

		res := BasePostResponse{
			Uuid:       post.ID,
			Body:       post.Body,
			CreatedAt:  post.CreatedAt,
			ThreadUuid: post.ThreadID,
			UserUuid:   post.UserID,
		}

		ws.Hub.broadcast <- res
	}
}
