package server

import (
	"context"
	"myChat-API2/internal/domain"
	"myChat-API2/internal/service"
)

var _ service.IAuthService = (*MockAuthService)(nil)
var _ service.IChatService = (*MockChatService)(nil)

type MockAuthService struct{}

func (m *MockAuthService) Signup(ctx context.Context, username string, password string) (string, error) {
	return "test-userID-1", nil
}

func (m *MockAuthService) Login(ctx context.Context, username string, password string) (string, error) {
	return "test-token-1", nil
}

type MockChatService struct{}

func (m *MockChatService) CreateRoom(ctx context.Context, name string, userID string) (string, error) {
	return "test-roomID-1", nil
}

func (m *MockChatService) ShowAllRooms(ctx context.Context, page int) ([]domain.Room, error) {
	rooms := []domain.Room{
		{
			ID:        "test-roomID-1",
			Name:      "test-room-name-1",
			CreatedAt: "test-room-createdAt-1",
			UserID:    "test-userID-1",
		}, {
			ID:        "test-roomID-2",
			Name:      "test-room-name-2",
			CreatedAt: "test-room-createdAt-2",
			UserID:    "test-userID-2",
		}, {
			ID:        "test-roomID-3",
			Name:      "test-room-name-3",
			CreatedAt: "test-room-createdAt-3",
			UserID:    "test-userID-3",
		},
	}
	return rooms, nil
}

func (m *MockChatService) SeeRoomDetail(ctx context.Context, roomID string) (domain.Room, []domain.Chat, error) {
	room := domain.Room{
		ID:        "test-roomID-1",
		Name:      "test-room-name-1",
		CreatedAt: "test-room-createdAt-1",
		UserID:    "test-userID-1",
	}
	chats := []domain.Chat{
		{
			ID:        "test-chatID-1",
			Body:      "test-body-1",
			RoomID:    "test-roomID-1",
			UserID:    "test-userID-1",
			CreatedAt: "test-createdAt-1",
		}, {
			ID:        "test-chatID-2",
			Body:      "test-body-2",
			RoomID:    "test-roomID-2",
			UserID:    "test-userID-2",
			CreatedAt: "test-createdAt-2",
		}, {
			ID:        "test-chatID-3",
			Body:      "test-body-3",
			RoomID:    "test-roomID-3",
			UserID:    "test-userID-3",
			CreatedAt: "test-createdAt-3",
		},
	}
	return room, chats, nil
}

func (m *MockChatService) CreateChat(ctx context.Context, body string, roomID string, userID string) (domain.Chat, error) {
	chat := domain.Chat{
		ID:        "test-chatID-1",
		Body:      "test-body-1",
		RoomID:    "test-roomID-1",
		UserID:    "test-userID-1",
		CreatedAt: "test-createdAt-1",
	}
	return chat, nil
}
