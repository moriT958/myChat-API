package server

import (
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

//type BaseChatResponse struct {
//	ID        string `json:"id"`
//	Body      string `json:"body"`
//	CreatedAt string `json:"createdAt"`
//	RoomID    string `json:"roomId"`
//	UserID    string `json:"userId"`
//}
//
//type CreateChatRequest struct {
//	Body   string `json:"body"`
//	RoomID string `json:"roomId"`
//}

func TestChatHandler(t *testing.T) {

	ws := NewServerWS(&MockChatService{})
	srv := httptest.NewServer(ws)
	defer srv.Close()

	wsURL := "ws" + strings.TrimPrefix(srv.URL, "http") + "/ws"
	wsConn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("could not open a ws connection on %s %v", wsURL, err)
	}
	defer wsConn.Close()

	reqBody := []byte(`
	{
	"id":"test-id-1",
	"body": "test-body-1",
	"createdAt": "test-2001-11-04",
	"roomId": "test-roomID-1",
	"userID": "test-userID-1"
	}
	`)
	if err := wsConn.WriteJSON(reqBody); err != nil {
		t.Fatal(err)
	}

	t.Run("crrectly response chat", func(t *testing.T) {
		want := map[string]string{"body": "test-body-1", "roomId": "test-roomID-1"}
		var got map[string]string
		if err := wsConn.ReadJSON(&got); err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %v, got %v", want, got)
		}
	})
}
