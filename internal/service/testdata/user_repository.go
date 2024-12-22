package testdata

import (
	"context"
	"myChat-API2/internal/domain"
	"time"
)

var users = []domain.User{
	domain.User{
		ID:       "uuid1",
		Name:     "user1",
		Password: "pass1",
	},
	domain.User{
		ID:       "uuid2",
		Name:     "user2",
		Password: "pass2",
	},
	domain.User{
		ID:       "uuid3",
		Name:     "user3",
		Password: "pass3",
	},
}

type MockUserRepository struct{}

func (m *MockUserRepository) Save(ctx context.Context, user domain.User) error {
	return nil
}

func (m *MockUserRepository) GetByID(ctx context.Context, userID string) (domain.User, error) {
	return users[0], nil
}

func (m *MockUserRepository) GetByName(ctx context.Context, username string) (domain.User, error) {
	return users[0], nil
}

func (m *MockUserRepository) GetCreatedAtByID(ctx context.Context, userID string) (string, error) {
	return time.Now().Format("2006-01-02 15:04:05"), nil
}
