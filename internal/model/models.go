package model

import (
	"time"

	"github.com/google/uuid"
)

type Thread struct {
	Id        int
	Uuid      uuid.UUID
	Topic     string
	CreatedAt time.Time
}

type Post struct {
	Id        int
	Uuid      uuid.UUID
	Body      string
	ThreadId  int
	CreatedAt time.Time
}
