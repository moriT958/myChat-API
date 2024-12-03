package model

import "time"

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	CreatedAt time.Time
}

type Post struct {
	Id        int
	Uuid      string
	Body      string
	ThreadId  int
	CreatedAt time.Time
}
