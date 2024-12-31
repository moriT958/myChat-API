package domain

import "context"

type IUserRepository interface {
	Save(context.Context, User) error
	GetByID(context.Context, string) (User, error)
	GetByName(context.Context, string) (User, error)
}

type User struct {
	ID       string
	Name     string
	Password string
}
