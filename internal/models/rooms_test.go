package models

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestNewRoom(t *testing.T) {
	positive, zero, negative := 8, 0, -2
	longName := strings.Repeat("a", 101)
	now := time.Date(2026, 9, 28, 16, 45, 40, 342, time.UTC)
	tests := []struct {
		testName string
		id       uuid.UUID //room id, etc.
		name     string
		capacity *int
		wantErr  error
	}{
		{"valid", uuid.New(), "Alpha", &positive, nil},
		{"no id", uuid.Nil, "Alpha", &positive, ErrRoomRequiredID},
		{"empty name", uuid.New(), "", &positive, ErrRoomRequiredName},
		{"long name", uuid.New(), longName, &positive, ErrRoomNameTooLong},
		{"negative capacity", uuid.New(), "Alpha", &negative, ErrRoomCapacityInvalid},
		{"zero capacity", uuid.New(), "Альфа", &zero, ErrRoomCapacityInvalid},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			room, err := NewRoom(tt.id, tt.name, nil, tt.capacity, now)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("error = %v, want = %v", err, tt.wantErr)
			}
			if tt.wantErr != nil {
				return
			}
			if room.Name != "Alpha" || room.ID != tt.id || !room.CreatedAt.Equal(now) {
				t.Fatalf("unexpected room: %+v", room)
			}
		})
	}
}
