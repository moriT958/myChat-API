package domain

import "context"

type PostRepositorier interface {
	GetThreadByUuid(context.Context, string) (Thread, error)
	SavePost(context.Context, Post) error
	GetPostsByThreadUuid(context.Context, string) ([]Post, error)
	GetCreatedAt(context.Context, string) (string, error)
}

type ThreadRepositorier interface {
	SaveThread(context.Context, Thread) error
	GetThreads(context.Context, int, int) ([]Thread, error)
	GetThreadByUuid(context.Context, string) (Thread, error)
	GetThreadId(context.Context, string) (int, error)
	GetCreatedAt(context.Context, string) (string, error)
}

type UserRepositorier interface {
	SaveUser(context.Context, User) error
	GetUserByUsername(context.Context, string) (User, error)
	GetUserId(context.Context, string) (int, error)
	GetCreatedAt(context.Context, string) (string, error)
}
