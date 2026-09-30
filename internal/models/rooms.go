package models

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
)

var (
	ErrRoomRequiredID      = errors.New("room: id is required")
	ErrRoomRequiredName    = errors.New("room: name is required")
	ErrRoomNameTooLong     = errors.New("room: name must be less than 100 characters")
	ErrRoomCapacityInvalid = errors.New("room: capacity must be better than 0")
)

type Room struct {
	ID          uuid.UUID
	Name        string
	Description *string
	Capacity    *int
	CreatedAt   time.Time
}

func NewRoom(
	id uuid.UUID,
	name string,
	description *string,
	capacity *int,
	createdAt time.Time) (Room, error) {
	room := Room{
		ID:          id,
		Name:        strings.TrimSpace(name),
		Description: description,
		Capacity:    capacity,
		CreatedAt:   createdAt.UTC(),
	}
	if err := room.Validate(); err != nil {
		return Room{}, err
	}
	return room, nil
}

func (r Room) Validate() error {
	if r.ID == uuid.Nil {
		return ErrRoomRequiredID
	}
	name := strings.TrimSpace(r.Name)
	if name == "" {
		return ErrRoomRequiredName
	}
	if utf8.RuneCountInString(name) > 100 {
		return ErrRoomNameTooLong
	}
	if r.Capacity != nil && *r.Capacity <= 0 {
		return ErrRoomCapacityInvalid
	}
	return nil
}
