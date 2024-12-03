package schema

//
// Thread schemas
//

type BaseThreadResponse struct {
	Uuid      string `json:"uuid"`
	Topic     string `json:"topic"`
	CreatedAt string `json:"createdAt"`
}

type CreateThreadRequest struct {
	Topic string `json:"topic"`
}

type GetThreadListResponse struct {
	Threads []BaseThreadResponse `json:"threads"`
}

type GetThreadDetailResponse struct {
	Uuid      string         `json:"uuid"`
	Topic     string         `json:"topic"`
	CreatedAt string         `json:"createdAt"`
	Posts     []PostOnThread `json:"posts"`
}

type PostOnThread struct {
	Uuid      string `json:"uuid"`
	Body      string `json:"body"`
	CreatedAt string `json:"createdAt"`
}

//
// Post schemas (websocket)
//

type InMessage struct {
	Body       string `json:"body"`
	ThreadUuid string `json:"threadUuid"`
}

type OutMessage struct {
	Uuid       string `json:"uuid"`
	Body       string `json:"body"`
	ThreadUuid string `json:"threadUuid"`
	CreatedAt  string `json:"createdAt"`
}
