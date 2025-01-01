package repository

import (
	"context"
	"myChat-API2/internal/domain"
)

var _ domain.IRoomRepository = (*RoomRepository)(nil)

type RoomRepository struct {
	db DBTX
}

func NewRoomRepository(db DBTX) *RoomRepository {
	return &RoomRepository{db: db}
}

func (r *RoomRepository) Save(ctx context.Context, room domain.Room) error {
	return nil
}

func (r *RoomRepository) GetAll(ctx context.Context, page int) ([]domain.Room, error) {
	return nil, nil
}

func (r *RoomRepository) GetByID(ctx context.Context, id string) (domain.Room, error) {
	return domain.Room{}, nil
}

func (r *RoomRepository) GetByUserID(ctx context.Context, userId string) ([]domain.Room, error) {
	return nil, nil
}
