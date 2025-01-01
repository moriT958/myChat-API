package repository

import (
	"context"
	"myChat-API2/internal/domain"
)

// Implemention of domain user repository interface.
var _ domain.IUserRepository = (*UserRepository)(nil)

type UserRepository struct {
	db DBTX
}

func NewUserRepository(db DBTX) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(ctx context.Context, user domain.User) error {
	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (domain.User, error) {
	return domain.User{}, nil
}

func (r *UserRepository) GetByName(ctx context.Context, name string) (domain.User, error) {
	return domain.User{}, nil
}
