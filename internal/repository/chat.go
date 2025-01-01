package repository

import (
	"context"
	"myChat-API2/internal/domain"
)

var _ domain.IChatRepository = (*ChatRepository)(nil)

type ChatRepository struct {
	store map[int]domain.Chat
}

func NewChatRepository() *ChatRepository {
	cr := new(ChatRepository)
	cr.store = make(map[int]domain.Chat)
	return cr
}
func (r *ChatRepository) Save(ctx context.Context, chat domain.Chat) error {
	return nil
}

func (r *ChatRepository) GetByID(ctx context.Context, id string) (domain.Chat, error) {
	return domain.Chat{}, nil
}

func (r *ChatRepository) GetByRoomID(ctx context.Context, roomID string) ([]domain.Chat, error) {
	return nil, nil
}
