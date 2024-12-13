package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        int
	Uuid      uuid.UUID
	Username  string
	Password  string
	CreatedAt time.Time
}

type Thread struct {
	Id        int
	Uuid      uuid.UUID
	Topic     string
	CreatedAt time.Time
	UserId    int
}

type Post struct {
	Id        int
	Uuid      uuid.UUID
	Body      string
	ThreadId  int
	CreatedAt time.Time
}
