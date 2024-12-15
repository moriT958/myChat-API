package repository

import (
	"context"
	"errors"
	"fmt"
	"myChat-API/internal/domain"
	"myChat-API/internal/query"

	"github.com/google/uuid"
)

type PostRepository struct {
	*query.Queries
}

func NewPostRepository(q *query.Queries) *PostRepository {
	return &PostRepository{Queries: q}
}

func (r *PostRepository) GetThreadByUuid(ctx context.Context, postUuid string) (domain.Thread, error) {

	params, err := uuid.Parse(postUuid)
	if err != nil {
		return domain.Thread{}, errors.New("failed to parse post uuid")
	}
	p, err := r.Queries.GetPostByUuid(ctx, params)
	if err != nil {
		return domain.Thread{}, err
	}

	t, err := r.Queries.GetThreadById(ctx, p.ThreadID)
	if err != nil {
		return domain.Thread{}, err
	}

	threadEntity := domain.Thread{
		Uuid:   t.Uuid.String(),
		Topic:  t.Topic,
		UserId: int(t.UserID),
	}
	return threadEntity, nil
}

func (r *PostRepository) SavePost(ctx context.Context, p domain.Post) error {
	params := query.CreatePostParams{
		Uuid:     uuid.MustParse(p.Uuid),
		Body:     p.Body,
		ThreadID: int32(p.ThreadId),
		UserID:   int32(p.UserId),
	}
	if err := r.Queries.CreatePost(ctx, params); err != nil {
		return fmt.Errorf("failed to save post data: %w\n", err)
	}

	return nil
}

func (r *PostRepository) GetPostsByThreadUuid(ctx context.Context, threadUuid string) ([]domain.Post, error) {
	t, err := r.Queries.GetThreadByUuid(ctx, uuid.MustParse(threadUuid))
	if err != nil {
		return nil, err
	}

	pl, err := r.Queries.GetPostByThreadId(ctx, int32(t.ID))
	if err != nil {
		return nil, err
	}

	postEntities := make([]domain.Post, len(pl))
	for i, p := range pl {
		postEntities[i] = domain.Post{
			Uuid:     p.Uuid.String(),
			Body:     p.Body,
			ThreadId: int(p.ThreadID),
			UserId:   int(p.UserID),
		}
	}

	return postEntities, nil
}

func (r *PostRepository) GetCreatedAt(ctx context.Context, postUuid string) (string, error) {
	p, err := r.Queries.GetPostByUuid(ctx, uuid.MustParse(postUuid))
	if err != nil {
		return "", err
	}
	return p.CreatedAt.Format("2006-01-02 15:04:05"), nil

}
