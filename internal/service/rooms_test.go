package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/myau69/bookem/internal/models"
)

type roomStoreFake struct {
	saved models.Room
	calls int
	err   error
}

func (f *roomStoreFake) Create(
	ctx context.Context,
	room models.Room,
) (models.Room, error) {
	f.calls++
	f.saved = room
	return room, f.err
}

func (f *roomStoreFake) List(
	ctx context.Context,
) ([]models.Room, error) {
	return []models.Room{}, nil
}

func TestRoomServiceCreate(t *testing.T) {
	now := time.Date(2006, 07, 21, 14, 55, 00, 00, time.UTC)
	store := &roomStoreFake{}
	svc := NewRoomService(store, func() time.Time { return now })
	room, err := svc.Create(context.Background(), "Alpha", nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if store.calls != 1 || room.ID == uuid.Nil || store.saved.ID != room.ID {
		t.Fatal("expected one saved room with assigned ID")
	}
	if store.saved.Name != "Alpha" || !store.saved.CreatedAt.Equal(now) {
		t.Fatalf("unexpected value: %+v", store.saved)
	}
	_, err = svc.Create(context.Background(), " ", nil, nil)
	if !errors.Is(err, models.ErrRoomRequiredName) || store.calls != 1 {
		t.Fatal("invalid room must not reach the repository")
	}
	storageErr := errors.New("storage unavaliable")
	store.err = storageErr
	_, err = svc.Create(context.Background(), "Beta", nil, nil)
	if !errors.Is(err, storageErr) {
		t.Fatalf("storage lost: %v", err)
	}
}
