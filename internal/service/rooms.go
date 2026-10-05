package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/myau69/bookem/internal/models"
	"github.com/myau69/bookem/internal/repository"
)

type RoomService struct {
	repo repository.RoomRepository
	now  func() time.Time
}

func NewRoomService(repo repository.RoomRepository, now func() time.Time) *RoomService {
	return &RoomService{repo: repo, now: now}
}

func (s *RoomService) Create(
	ctx context.Context,
	name string,
	description *string,
	capacity *int,
) (models.Room, error) {
	room, err := models.NewRoom(uuid.New(), name, description, capacity, s.now())
	if err != nil {
		return models.Room{}, err
	}
	saved, err := s.repo.Create(ctx, room)
	if err != nil {
		return models.Room{}, fmt.Errorf("save room: %w", err)
	}
	return saved, nil
}

func (s *RoomService) List(ctx context.Context) ([]models.Room, error) {
	return s.repo.List(ctx)
}
