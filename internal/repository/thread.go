package repository

import (
	"context"
	"myChat-API/internal/domain"
	"myChat-API/internal/query"
	"time"

	"github.com/google/uuid"
)

type ThreadRepositoy struct {
	*query.Queries
}

func NewThreadRepository(q *query.Queries) *ThreadRepositoy {
	return &ThreadRepositoy{Queries: q}
}

func (r *ThreadRepositoy) SaveThread(ctx context.Context, t domain.Thread) error {
	params := query.CreateThreadParams{
		Uuid:      uuid.New(),
		Topic:     t.Topic,
		CreatedAt: time.Now(),
		UserID:    int32(t.UserId),
	}
	if err := r.Queries.CreateThread(ctx, params); err != nil {
		return err
	}

	return nil
}

func (r *ThreadRepositoy) GetThreads(ctx context.Context, limit int, offset int) ([]domain.Thread, error) {
	params := query.GetAllThreadsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	tl, err := r.Queries.GetAllThreads(ctx, params)
	if err != nil {
		return []domain.Thread{}, err
	}

	threadEntityList := make([]domain.Thread, limit)
	for i, t := range tl {
		threadEntityList[i] = domain.Thread{
			Uuid:   t.Uuid.String(),
			Topic:  t.Topic,
			UserId: int(t.UserID),
		}
	}

	return threadEntityList, nil
}

func (r *ThreadRepositoy) GetThreadByUuid(ctx context.Context, threadUuid string) (domain.Thread, error) {
	param, err := uuid.Parse(threadUuid)
	if err != nil {
		return domain.Thread{}, err
	}
	t, err := r.Queries.GetThreadByUuid(ctx, param)
	if err != nil {
		return domain.Thread{}, err
	}

	threadEntity := domain.Thread{
		Uuid:   threadUuid,
		Topic:  t.Topic,
		UserId: int(t.UserID),
	}
	return threadEntity, nil
}

func (r *ThreadRepositoy) GetThreadId(ctx context.Context, threadUuid string) (int, error) {
	t, err := r.Queries.GetThreadByUuid(ctx, uuid.MustParse(threadUuid))
	if err != nil {
		return 0, err
	}
	return int(t.ID), nil
}

func (r *ThreadRepository) GetCreatedAt(ctx context.Context, threadUuid string) (string, error) {
	p, err := r.Queries.GetPostByUuid(ctx, uuid.MustParse(threadUuid))
	if err != nil {
		return "", err
	}
	return p.CreatedAt.Format("2006-01-02 15:04:05"), nil

}
