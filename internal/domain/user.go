package domain

import "context"

type IUserRepository interface {
	Save(context.Context, User) error
	GetByID(context.Context, string) (User, error)
	GetByName(context.Context, string) (User, error)
	GetCreatedAtByID(context.Context, string) (string, error)
}

type User struct {
	ID       string
	Name     string
	Password string
}

// TODO:
// Write Password safety check validation logic here.
// Write as User model's method.
