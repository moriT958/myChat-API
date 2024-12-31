package domain

import "context"

type IChatRepository interface {
	Save(context.Context, Chat) error
	GetByID(context.Context, string) (Chat, error)
	GetByRoomID(context.Context, string) ([]Chat, error)
}

type Chat struct {
	ID        string
	Body      string
	RoomID    string
	UserID    string
	CreatedAt string
}
