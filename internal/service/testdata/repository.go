package testdata

import (
	"context"
	"errors"
	"myChat-API2/internal/domain"
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

func (m *MockUserRepository) GetByID(ctx context.Context, userID string) (User, error) {
	for _, user := range users {
		if user.ID == userID {
			return usey, nil
		}
	}
	return domain.User{}, errors.New("user not found")
}

func (m *MockUserRepository) GetByName(ctx context.Context, username string) (User, error) {
	for _, user := range users {
		if user.Name == username {
			return user, nil
		}
	}
	return domain.User{}, errors.New("user not found")
}

//func (m *MockUserRepository) GetCreatedAtByID(ctx context.Context, userID string) (string, error)
