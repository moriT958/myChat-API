package domain

import "context"

type IRoomRepository interface {
	Save(context.Context, Room) error
	GetAll(context.Context, int, int) ([]Room, error)
	GetByID(context.Context, string) (Room, error)
	GetByUserID(context.Context, string) ([]Room, error)
}

type Room struct {
	ID        string
	Topic     string
	CreatedAt string
	UserID    string
}
