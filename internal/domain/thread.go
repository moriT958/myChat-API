package domain

import "context"

type IThreadRepository interface {
	Save(context.Context, Thread) error
	GetAll(context.Context, int, int) ([]Thread, error)
	GetByID(context.Context, string) (Thread, error)
	GetByUserID(context.Context, string) ([]Thread, error)
}

type Thread struct {
	ID        string
	Topic     string
	CreatedAt string
	UserID    string
}
