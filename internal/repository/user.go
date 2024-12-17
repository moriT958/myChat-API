package repository

import (
	"context"
	"myChat-API2/internal/domain"
	"myChat-API2/internal/query"
	"time"

	"github.com/google/uuid"
)

// Implemention of domain user repository interface.
type UserRepository struct {
	*query.Queries
}

func NewUserRepository(q *query.Queries) *UserRepository {
	return &UserRepository{Queries: q}
}

func (r *UserRepository) Save(ctx context.Context, user domain.User) error {
	params := query.CreateUserParams{
		Uuid:      uuid.MustParse(user.ID),
		Username:  user.Name,
		Password:  user.Password,
		CreatedAt: time.Now(),
	}
	if err := r.Queries.CreateUser(ctx, params); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (domain.User, error) {
	u, err := r.Queries.GetUserByUuid(ctx, uuid.MustParse(id))
	if err != nil {
		return domain.User{}, err
	}

	user := domain.User{
		ID:       u.Uuid.String(),
		Name:     u.Username,
		Password: u.Password,
	}
	return user, nil
}

func (r *UserRepository) GetByName(ctx context.Context, name string) (domain.User, error) {
	u, err := r.Queries.GetUserByName(ctx, name)
	if err != nil {
		return domain.User{}, err
	}

	user := domain.User{
		ID:       u.Uuid.String(),
		Name:     u.Username,
		Password: u.Password,
	}
	return user, nil
}

func (r *UserRepository) GetCreatedAtByID(ctx context.Context, userID string) (string, error) {
	u, err := r.GetUserByUuid(ctx, uuid.MustParse(userID))
	if err != nil {
		return "", err
	}

	return u.CreatedAt.Format("2006-01-02 15:04:05"), nil
}
