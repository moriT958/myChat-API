package repository

import (
	"context"
	"myChat-API2/internal/domain"
	"myChat-API2/internal/query"
)

var _ domain.IChatRepository = (*ChatRepository)(nil)

type ChatRepository struct {
	*query.Queries
}

func NewChatRepository(q *query.Queries) *ChatRepository {
	return &ChatRepository{Queries: q}
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

//func (r *ChatRepository) Save(ctx context.Context, chat domain.Chat) error {
//	createdAt, err := time.Parse("2006-01-02 15:04:05", chat.CreatedAt)
//	if err != nil {
//		return err
//	}
//
//	params := query.CreatePostParams{
//		Uuid:      uuid.MustParse(chat.ID),
//		Body:      chat.Body,
//		CreatedAt: createdAt,
//		Uuid_2:    uuid.MustParse(chat.RoomID),
//		Uuid_3:    uuid.MustParse(chat.UserID),
//	}
//	if err := r.Queries.CreatePost(ctx, params); err != nil {
//		return err
//	}
//	return nil
//}
//
//func (r *ChatRepository) GetByID(ctx context.Context, id string) (domain.Chat, error) {
//	p, err := r.Queries.GetPostByUuid(ctx, uuid.MustParse(id))
//	if err != nil {
//		return domain.Chat{}, err
//	}
//
//	post := domain.Chat{
//		ID:        p.PostUuid.String(),
//		Body:      p.Body,
//		RoomID:    p.ThreadUuid.String(),
//		UserID:    p.UserUuid.String(),
//		CreatedAt: p.PostCreatedAt.Format("2006-01-02 15:04:05"),
//	}
//	return post, nil
//}
//
//func (r *ChatRepository) GetByRoomID(ctx context.Context, threadId string) ([]domain.Post, error) {
//	pl, err := r.Queries.GetPostByThreadUuid(ctx, uuid.MustParse(threadId))
//	if err != nil {
//		return nil, err
//	}
//
//	posts := make([]domain.Post, len(pl))
//	for i, p := range pl {
//		posts[i] = domain.Post{
//			ID:        p.PostUuid.String(),
//			Body:      p.Body,
//			ThreadID:  p.ThreadUuid.String(),
//			UserID:    p.UserUuid.String(),
//			CreatedAt: p.PostCreatedAt.Format("2006-01-02 15:04:05"),
//		}
//	}
//	return posts, nil
//}
