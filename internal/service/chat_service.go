package service

import (
	"context"
	"fmt"
	"myChat-API2/internal/domain"
	"time"

	"github.com/google/uuid"
)

type IChatService interface {
	CreateThread(context.Context, string, string) (string, error)
	ShowAllThreads(context.Context, int, int) ([]domain.Thread, error)
	SeeThreadDetail(context.Context, string) (domain.Thread, []domain.Post, error)
	CreatePost(context.Context, string, string, string) (domain.Post, error)
}

type ChatService struct {
	ThreadRepo domain.IThreadRepository
	PostRepo   domain.IPostRepository
	UserRepo   domain.IUserRepository
}

func NewChatService(
	tr domain.IThreadRepository,
	pr domain.IPostRepository,
	ur domain.IUserRepository,
) *ChatService {
	return &ChatService{
		ThreadRepo: tr,
		PostRepo:   pr,
		UserRepo:   ur,
	}
}

func (s *ChatService) CreateThread(ctx context.Context, topic string, userID string) (string, error) {
	u, err := s.UserRepo.GetByID(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("User doesn't exist: %w", err)
	}

	thread := domain.Thread{
		ID:        uuid.NewString(),
		Topic:     topic,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UserID:    u.ID,
	}

	if err := s.ThreadRepo.Save(ctx, thread); err != nil {
		return "", err
	}

	return thread.ID, nil
}

func (s *ChatService) ShowAllThreads(ctx context.Context, limit int, offset int) ([]domain.Thread, error) {
	threads, err := s.ThreadRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("Failed to load threads data: %w", err)
	}
	return threads, nil
}

func (s *ChatService) SeeThreadDetail(ctx context.Context, threadID string) (domain.Thread, []domain.Post, error) {
	thread, err := s.ThreadRepo.GetByID(ctx, threadID)
	if err != nil {
		return domain.Thread{}, nil, err
	}

	posts, err := s.PostRepo.GetByThreadID(ctx, threadID)
	if err != nil {
		return thread, nil, fmt.Errorf("Failed to load posts data: %w", err)
	}

	return thread, posts, nil
}

func (s *ChatService) CreatePost(ctx context.Context, body string, threadID string, userID string) (domain.Post, error) {
	if _, err := s.UserRepo.GetByID(ctx, userID); err != nil {
		return domain.Post{}, fmt.Errorf("User doesn't exist: %w", err)
	}

	if _, err := s.ThreadRepo.GetByID(ctx, threadID); err != nil {
		return domain.Post{}, fmt.Errorf("Thread doesn't exist: %w", err)
	}

	post := domain.Post{
		ID:        uuid.NewString(),
		Body:      body,
		ThreadID:  threadID,
		UserID:    userID,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	if err := s.PostRepo.Save(ctx, post); err != nil {
		return domain.Post{}, fmt.Errorf("Failed to post: %w", err)
	}

	return post, nil
}
