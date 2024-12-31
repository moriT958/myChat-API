package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (h *TodoServer) CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parse Request
	var reqNewRoom CreateRoomRequest
	err := json.NewDecoder(r.Body).Decode(&reqNewRoom)
	if err != nil || reqNewRoom.Name == "" {
		log.Println(err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Create Thread.
	userID := getUserID(ctx)
	roomID, err := h.ChatService.CreateRoom(ctx, reqNewRoom.Name, userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to post thread.", http.StatusInternalServerError)
		return
	}

	// Generate response.
	res := CreateRoomResponses{
		ID: roomID,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(&res); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *TodoServer) GetRoomListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q, err := getQueryParams(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Parameters", http.StatusBadRequest)
		return
	}

	rooms, err := h.ChatService.ShowAllRooms(ctx, q["page"])
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to read rooms.", http.StatusInternalServerError)
		return
	}

	// Generate response.
	var res GetRoomListResponse
	res.Rooms = make([]BaseRoomResponse, len(rooms))
	for i, r := range rooms {
		res.Rooms[i] = BaseRoomResponse{
			ID:        r.ID,
			Name:      r.Name,
			CreatedAt: r.CreatedAt,
			UserID:    r.UserID,
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

// TODO:
// need to fix
func getQueryParams(r *http.Request) (map[string]int, error) {
	res := map[string]int{
		"offset": 0,
		"limit":  3,
	}

	qParams := r.URL.Query()

	if val, exists := qParams["offset"]; exists && len(val) > 0 {
		offset, err := strconv.Atoi(val[0])
		if err != nil {
			return res, err
		}
		res["offset"] = offset
	}

	if val, exists := qParams["limit"]; exists && len(val) > 0 {
		limit, err := strconv.Atoi(val[0])
		if err != nil {
			return res, err
		}
		res["limit"] = limit
	}

	return res, nil
}

func (h *TodoServer) ReadRoomDetailHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var roomID string = r.PathValue("roomID")

	room, chats, err := h.ChatService.SeeRoomDetail(ctx, roomID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to read room.", http.StatusInternalServerError)
		return
	}

	// Generate response.
	res := GetRoomDetailResponse{
		ID:        room.ID,
		Name:      room.Name,
		CreatedAt: room.CreatedAt,
		Chats:     make([]BaseChatResponse, len(chats)),
		UserID:    room.UserID,
	}

	for i, c := range chats {
		res.Chats[i] = BaseChatResponse{
			ID:        c.ID,
			Body:      c.Body,
			CreatedAt: c.CreatedAt,
			RoomID:    c.RoomID,
			UserID:    c.UserID,
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
