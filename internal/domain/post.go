package domain

import "context"

type IPostRepository interface {
	Save(context.Context, Post) error
	GetByID(context.Context, string) (Post, error)
	GetByThreadID(context.Context, string) ([]Post, error)
	GetByUserID(context.Context, string) ([]Post, error)
}

type Post struct {
	ID        string
	Body      string
	ThreadID  string
	UserID    string
	CreatedAt string
}
