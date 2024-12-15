package service

import (
	"context"
	"myChat-API/internal/domain"

	"github.com/google/uuid"
)

type ThreadService struct {
	Tr domain.ThreadRepositorier
	Pr domain.PostRepositorier
	Ur domain.UserRepositorier
}

func NewThreadService(
	tr domain.ThreadRepositorier,
	pr domain.PostRepositorier,
	ur domain.UserRepositorier,
) *ThreadService {
	return &ThreadService{
		Tr: tr,
		Pr: pr,
		Ur: ur,
	}
}

func (s *ThreadService) CreateThread(ctx context.Context, topic string, username string) (string, error) {
	u, err := s.Ur.GetUserByUsername(ctx, username)
	if err != nil {
		return "", err
	}
	userId, err := s.Ur.GetUserId(ctx, u.Uuid)
	if err != nil {
		return "", err
	}

	t := domain.Thread{
		Uuid:   uuid.NewString(),
		Topic:  topic,
		UserId: userId,
	}
	if err := s.Tr.SaveThread(ctx, t); err != nil {
		return "", err
	}

	return t.Uuid, nil
}

func (s *ThreadService) ReadAllThreads(ctx context.Context, limit int, offset int) ([]domain.Thread, error) {
	t, err := s.Tr.GetThreads(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (s *ThreadService) ReadThreadDetail(ctx context.Context, threadUuid string) (domain.Thread, []domain.Post, error) {
	t, err := s.Tr.GetThreadByUuid(ctx, threadUuid)
	if err != nil {
		return domain.Thread{}, nil, err
	}

	p, err := s.Pr.GetPostsByThreadUuid(ctx, threadUuid)
	if err != nil {
		return domain.Thread{}, nil, err
	}

	return t, p, nil
}
