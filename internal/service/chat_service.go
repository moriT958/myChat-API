package service

import (
	"context"
	"fmt"
	"myChat-API2/internal/domain"
	"time"

	"github.com/google/uuid"
)

type IChatService interface {
	CreateRoom(context.Context, string, string) (string, error)
	ShowAllRooms(context.Context, int) ([]domain.Room, error)
	SeeRoomDetail(context.Context, string) (domain.Room, []domain.Chat, error)
	CreateChat(context.Context, string, string, string) (domain.Chat, error)
}

var _ IChatService = (*ChatService)(nil)

type ChatService struct {
	RoomRepo domain.IRoomRepository
	ChatRepo domain.IChatRepository
	UserRepo domain.IUserRepository
}

func NewChatService(
	rr domain.IRoomRepository,
	cr domain.IChatRepository,
	ur domain.IUserRepository,
) *ChatService {
	return &ChatService{
		RoomRepo: rr,
		ChatRepo: cr,
		UserRepo: ur,
	}
}

func (s *ChatService) CreateRoom(ctx context.Context, topic string, userID string) (string, error) {
	u, err := s.UserRepo.GetByID(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("User doesn't exist: %w", err)
	}

	room := domain.Room{
		ID:        uuid.NewString(),
		Name:      topic,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UserID:    u.ID,
	}

	if err := s.RoomRepo.Save(ctx, room); err != nil {
		return "", err
	}

	return room.ID, nil
}

func (s *ChatService) ShowAllRooms(ctx context.Context, page int) ([]domain.Room, error) {
	rooms, err := s.RoomRepo.GetAll(ctx, page)
	if err != nil {
		return nil, fmt.Errorf("Failed to load threads data: %w", err)
	}
	return rooms, nil
}

func (s *ChatService) SeeRoomDetail(ctx context.Context, roomID string) (domain.Room, []domain.Chat, error) {
	room, err := s.RoomRepo.GetByID(ctx, roomID)
	if err != nil {
		return domain.Room{}, nil, err
	}

	chats, err := s.ChatRepo.GetByRoomID(ctx, roomID)
	if err != nil {
		return room, nil, fmt.Errorf("Failed to load posts data: %w", err)
	}

	return room, chats, nil
}

func (s *ChatService) CreateChat(ctx context.Context, body string, roomID string, userID string) (domain.Chat, error) {
	if _, err := s.UserRepo.GetByID(ctx, userID); err != nil {
		return domain.Chat{}, fmt.Errorf("User doesn't exist: %w", err)
	}

	if _, err := s.RoomRepo.GetByID(ctx, roomID); err != nil {
		return domain.Chat{}, fmt.Errorf("Thread doesn't exist: %w", err)
	}

	post := domain.Chat{
		ID:        uuid.NewString(),
		Body:      body,
		RoomID:    roomID,
		UserID:    userID,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	if err := s.ChatRepo.Save(ctx, post); err != nil {
		return domain.Chat{}, fmt.Errorf("Failed to post: %w", err)
	}

	return post, nil
}
