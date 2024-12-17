package repository

import (
	"context"
	"myChat-API2/internal/domain"
	"myChat-API2/internal/query"
	"time"

	"github.com/google/uuid"
)

type ThreadRepository struct {
	*query.Queries
}

func NewThreadRepository(q *query.Queries) *ThreadRepository {
	return &ThreadRepository{Queries: q}
}

func (r *ThreadRepository) Save(ctx context.Context, thread domain.Thread) error {
	createdAt, err := time.Parse("2006-01-02 15:04:05", thread.CreatedAt)
	if err != nil {
		return err
	}

	params := query.CreateThreadParams{
		Uuid:      uuid.MustParse(thread.ID),
		Topic:     thread.Topic,
		CreatedAt: createdAt,
		// Uuid_2 is UserID(type uuid).
		Uuid_2: uuid.MustParse(thread.UserID),
	}

	if err := r.Queries.CreateThread(ctx, params); err != nil {
		return err
	}
	return nil
}

func (r *ThreadRepository) GetAll(ctx context.Context, limit int, offset int) ([]domain.Thread, error) {
	params := query.GetAllThreadsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	tl, err := r.Queries.GetAllThreads(ctx, params)
	if err != nil {
		return nil, err
	}

	threads := make([]domain.Thread, len(tl))
	for i, t := range tl {
		threads[i] = domain.Thread{
			ID:        t.ThreadUuid.String(),
			Topic:     t.Topic,
			CreatedAt: t.ThreadCreatedAt.Format("2006-01-02 15:04:05"),
			UserID:    t.UserUuid.String(),
		}
	}
	return threads, nil
}

func (r *ThreadRepository) GetByID(ctx context.Context, id string) (domain.Thread, error) {
	t, err := r.Queries.GetThreadByUuid(ctx, uuid.MustParse(id))
	if err != nil {
		return domain.Thread{}, err
	}

	thread := domain.Thread{
		ID:        t.ThreadUuid.String(),
		Topic:     t.Topic,
		CreatedAt: t.ThreadCreatedAt.Format("2006-01-02 15:04:05"),
		UserID:    t.UserUuid.String(),
	}
	return thread, nil
}

func (r *ThreadRepository) GetByUserID(ctx context.Context, userId string) ([]domain.Thread, error) {
	tl, err := r.Queries.GetThreadByUserUuid(ctx, uuid.MustParse(userId))
	if err != nil {
		return nil, err
	}

	threads := make([]domain.Thread, len(tl))
	for i, t := range tl {
		threads[i] = domain.Thread{
			ID:        t.ThreadUuid.String(),
			Topic:     t.Topic,
			CreatedAt: t.ThreadCreatedAt.Format("2006-01-02 15:04:05"),
			UserID:    t.UserUuid.String(),
		}
	}
	return threads, nil
}
