package server

//
// Thread schemas (rest)
//

type BaseRoomResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UserID    string `json:"userId"`
}

type CreateRoomRequest struct {
	Name string `json:"name"`
}

type CreateRoomResponses struct {
	ID string `json:"id"`
}

type GetRoomListResponse struct {
	Rooms []BaseRoomResponse `json:"rooms"`
}

type GetRoomDetailResponse struct {
	ID        string             `json:"id"`
	Name      string             `json:"name"`
	CreatedAt string             `json:"createdAt"`
	Chats     []BaseChatResponse `json:"chats"`
	UserID    string             `json:"userId"`
}

// User schemas (rest)
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	ID string `json:"id"`
}

type GetTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type GetUserResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"create_at"`
}

//
// Post schemas (websocket)
//

type BaseChatResponse struct {
	ID        string `json:"id"`
	Body      string `json:"body"`
	CreatedAt string `json:"createdAt"`
	RoomID    string `json:"roomId"`
	UserID    string `json:"userId"`
}

type CreateChatRequest struct {
	Body   string `json:"body"`
	RoomID string `json:"roomId"`
}
