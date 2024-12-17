package repository

import (
	"context"
	"myChat-API2/internal/domain"
	"myChat-API2/internal/query"
	"time"

	"github.com/google/uuid"
)

type PostRepository struct {
	*query.Queries
}

func NewPostRepository(q *query.Queries) *PostRepository {
	return &PostRepository{Queries: q}
}

func (r *PostRepository) Save(ctx context.Context, post domain.Post) error {
	createdAt, err := time.Parse("2006-01-02 15:04:05", post.CreatedAt)
	if err != nil {
		return err
	}

	params := query.CreatePostParams{
		Uuid:      uuid.MustParse(post.ID),
		Body:      post.Body,
		CreatedAt: createdAt,
		Uuid_2:    uuid.MustParse(post.ThreadID),
		Uuid_3:    uuid.MustParse(post.UserID),
	}
	if err := r.Queries.CreatePost(ctx, params); err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) GetByID(ctx context.Context, id string) (domain.Post, error) {
	p, err := r.Queries.GetPostByUuid(ctx, uuid.MustParse(id))
	if err != nil {
		return domain.Post{}, err
	}

	post := domain.Post{
		ID:        p.PostUuid.String(),
		Body:      p.Body,
		ThreadID:  p.ThreadUuid.String(),
		UserID:    p.UserUuid.String(),
		CreatedAt: p.PostCreatedAt.Format("2006-01-02 15:04:05"),
	}
	return post, nil
}

func (r *PostRepository) GetByThreadID(ctx context.Context, threadId string) ([]domain.Post, error) {
	pl, err := r.Queries.GetPostByThreadUuid(ctx, uuid.MustParse(threadId))
	if err != nil {
		return nil, err
	}

	posts := make([]domain.Post, len(pl))
	for i, p := range pl {
		posts[i] = domain.Post{
			ID:        p.PostUuid.String(),
			Body:      p.Body,
			ThreadID:  p.ThreadUuid.String(),
			UserID:    p.UserUuid.String(),
			CreatedAt: p.PostCreatedAt.Format("2006-01-02 15:04:05"),
		}
	}
	return posts, nil
}

func (r *PostRepository) GetByUserID(ctx context.Context, userId string) ([]domain.Post, error) {
	pl, err := r.Queries.GetPostByThreadUuid(ctx, uuid.MustParse(userId))
	if err != nil {
		return nil, err
	}

	posts := make([]domain.Post, len(pl))
	for i, p := range pl {
		posts[i] = domain.Post{
			ID:        p.PostUuid.String(),
			Body:      p.Body,
			ThreadID:  p.ThreadUuid.String(),
			UserID:    p.UserUuid.String(),
			CreatedAt: p.PostCreatedAt.Format("2006-01-02 15:04:05"),
		}
	}
	return posts, nil
}
