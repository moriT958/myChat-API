package model

import "time"

type Thread struct {
	Uuid      string
	Topic     string
	CreatedAt time.Time
}

type Post struct {
	Uuid      string
	Body      string
	ThreadId  int
	CreatedAt time.Time
}
