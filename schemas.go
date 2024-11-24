package main

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
	Uuid      string             `json:"uuid"`
	Topic     string             `json:"topic"`
	CreatedAt string             `json:"createdAt"`
	Posts     []BasePostResponse `json:"posts"`
}

//
// Post schemas
//

type BasePostResponse struct {
	Uuid       string `json:"uuid"`
	Body       string `json:"body"`
	ThreadUuid string `json:"threadId"`
	CreatedAt  string `json:"createdAt"`
}

type CreatePostRequest struct {
	Body       string `json:"body"`
	ThreadUuid string `json:"threadUuid"`
}

type GetPostListResponse struct {
	Posts []BasePostResponse `json:"posts"`
}
