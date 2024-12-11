package schema

//
// Thread schemas (rest)
//

type BaseThreadResponse struct {
	Uuid      string `json:"uuid"`
	Topic     string `json:"topic"`
	CreatedAt string `json:"createdAt"`
	Username  string `json:"username"`
}

type CreateThreadRequest struct {
	Topic    string `json:"topic"`
	Username string `json:"username"`
}

type GetThreadListResponse struct {
	Threads []BaseThreadResponse `json:"threads"`
}

type GetThreadDetailResponse struct {
	Uuid      string         `json:"uuid"`
	Topic     string         `json:"topic"`
	CreatedAt string         `json:"createdAt"`
	Posts     []PostOnThread `json:"posts"`
	Username  string         `json:"username"`
}

type PostOnThread struct {
	Uuid      string `json:"uuid"`
	Body      string `json:"body"`
	CreatedAt string `json:"createdAt"`
	Username  string `json:"username"`
}

// User schemas (rest)
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserResponce struct {
	Uuid string `json:"uuid"`
}

type GetUserResponse struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	CreateAt string `json:"create_at"`
}

//
// Post schemas (websocket)
//

type InMessage struct {
	Body       string `json:"body"`
	ThreadUuid string `json:"threadUuid"`
	Username   string `json:"username"`
}

type OutMessage struct {
	Uuid       string `json:"uuid"`
	Body       string `json:"body"`
	ThreadUuid string `json:"threadUuid"`
	CreatedAt  string `json:"createdAt"`
	Username   string `json:"username"`
}
