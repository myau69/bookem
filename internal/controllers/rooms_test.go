package controllers

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/myau69/bookem/internal/models"
	"github.com/myau69/bookem/internal/service"
)

type roomHTTPStore struct {
	items []models.Room
}

func (s *roomHTTPStore) Create(ctx context.Context, room models.Room) (models.Room, error) {
	s.items = append(s.items, room)
	return room, nil
}

func (s *roomHTTPStore) List(ctx context.Context) ([]models.Room, error) {
	return s.items, nil
}

func TestRoomHTTP(t *testing.T) {
	store := &roomHTTPStore{}
	handler := NewRoomsHandler(service.NewRoomService(store, time.Now))
	list := httptest.NewRecorder()
	handler.List(list, httptest.NewRequest("GET", "/rooms/list", nil))
	if list.Code != 200 || strings.TrimSpace(list.Body.String()) != `{"rooms":[]}` {
		t.Fatalf("empty list: %d %s", list.Code, list.Body.String())
	}
	for _, body := range []string{`"name":" "`, `"name":"A", "capacity":"0"`, `"name":"A"{}`} {
		rec := httptest.NewRecorder()
		handler.Create(rec, httptest.NewRequest("POST", "/rooms/create", strings.NewReader(body)))
		if rec.Code != 400 {
			t.Fatalf("invalid body %s: status %d", body, rec.Code)
		}
	}
	if len(store.items) != 0 {
		t.Fatalf("invalid request changed store")
	}
	rec := httptest.NewRecorder()
	handler.Create(rec, httptest.NewRequest("POST", "/rooms/create", strings.NewReader(`{"name":" Alpha ","capacity":8}`)))
	if rec.Code != 201 {
		t.Fatalf("create: %d %s", rec.Code, rec.Body.String())
	}
	var response CreateRoomResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}
	if response.Room.Name != "Alpha" ||
		response.Room.Capacity == nil ||
		*response.Room.Capacity != 8 ||
		len(store.items) != 1 {
		t.Fatalf("unexpected error: %v", response)
	}
}
