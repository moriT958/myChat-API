package repository

import (
	"context"
	"myChat-API/internal/domain"
	"myChat-API/internal/query"
	"time"

	"github.com/google/uuid"
)

type UserRepository struct {
	*query.Queries
}

func NewUserRepository(q *query.Queries) *UserRepository {
	return &UserRepository{Queries: q}
}

func (r *UserRepository) SaveUser(ctx context.Context, u domain.User) error {
	params := query.CreateUserParams{
		Uuid:      uuid.MustParse(u.Uuid),
		Username:  u.Username,
		Password:  u.Password,
		CreatedAt: time.Now(),
	}

	if err := r.Queries.CreateUser(ctx, params); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {
	u, err := r.Queries.GetUserByName(ctx, username)
	if err != nil {
		return domain.User{}, err
	}

	userEntity := domain.User{
		Username: u.Username,
		Password: u.Password,
	}
	return userEntity, nil
}

func (r *UserRepository) GetUserId(ctx context.Context, userUuid string) (int, error) {
	u, err := r.Queries.GetUserByUuid(ctx, uuid.MustParse(userUuid))
	if err != nil {
		return 0, err
	}

	return int(u.ID), nil
}

func (r *UserRepository) GetCreatedAt(ctx context.Context, userUuid string) (string, error) {
	p, err := r.Queries.GetPostByUuid(ctx, uuid.MustParse(userUuid))
	if err != nil {
		return "", err
	}
	return p.CreatedAt.Format("2006-01-02 15:04:05"), nil

}
