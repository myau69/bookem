package repository

import (
	"context"

	"github.com/myau69/bookem/internal/models"
)

type RoomRepository interface {
	Create(ctx context.Context, room models.Room) (models.Room, error)
	List(ctx context.Context) ([]models.Room, error)
}
