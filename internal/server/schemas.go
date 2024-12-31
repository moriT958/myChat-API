package server

//
// Thread schemas (rest)
//

type BaseThreadResponse struct {
	Uuid      string `json:"id"`
	Topic     string `json:"topic"`
	CreatedAt string `json:"createdAt"`
	UserUuid  string `json:"userId"`
}

type CreateThreadRequest struct {
	Topic string `json:"topic"`
}

type CreateThreadResponses struct {
	Uuid string `json:"id"`
}

type GetThreadListResponse struct {
	Threads []BaseThreadResponse `json:"threads"`
}

type GetThreadDetailResponse struct {
	Uuid      string             `json:"id"`
	Topic     string             `json:"topic"`
	CreatedAt string             `json:"createdAt"`
	Posts     []BasePostResponse `json:"posts"`
	UserUuid  string             `json:"userId"`
}

// User schemas (rest)
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Uuid string `json:"id"`
}

type GetTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type GetUserResponse struct {
	Uuid      string `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"create_at"`
}

//
// Post schemas (websocket)
//

type BasePostResponse struct {
	Uuid       string `json:"id"`
	Body       string `json:"body"`
	CreatedAt  string `json:"createdAt"`
	ThreadUuid string `json:"threadId"`
	UserUuid   string `json:"userId"`
}

type CreatePostRequest struct {
	Body       string `json:"body"`
	ThreadUuid string `json:"threadId"`
}
