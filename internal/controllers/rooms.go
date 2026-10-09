package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/myau69/bookem/internal/models"
)

type RoomsUseCase interface {
	Create(context.Context, string, *string, *int) (models.Room, error)
	List(context.Context) ([]models.Room, error)
}

type RoomsHandler struct {
	svc RoomsUseCase
}

func NewRoomsHandler(svc RoomsUseCase) *RoomsHandler {
	return &RoomsHandler{svc: svc}
}

type createRoomRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Capacity    *int    `json:"capacity"`
}

type RoomResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Capacity    *int      `json:"capacity"`
	CreatedAt   time.Time `json:"createdAt"`
}

type CreateRoomResponse struct {
	Room RoomResponse `json:"room"`
}

type RoomsListResponse struct {
	Rooms []RoomResponse `json:"rooms"`
}

func mapRoom(room models.Room) RoomResponse {
	return RoomResponse{
		ID:          room.ID,
		Name:        room.Name,
		Description: room.Description,
		Capacity:    room.Capacity,
		CreatedAt:   room.CreatedAt.UTC(),
	}
}

func (h *RoomsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createRoomRequest
	if err := decodeJson(w, r, &req); err != nil {
		writeError(w, 400, "INVALID_REQUEST", "invalid request body")
		return
	}
	room, err := h.svc.Create(r.Context(), req.Name, req.Description, req.Capacity)
	if err != nil {
		if errors.Is(err, models.ErrRoomRequiredName) ||
			errors.Is(err, models.ErrRoomNameTooLong) ||
			errors.Is(err, models.ErrRoomCapacityInvalid) {
			writeError(w, 400, "INVALID_REQUEST", "invalid request body")
			return
		}
		log.Printf("create room: %v", err)
		writeError(w, 500, "INTERNAL_ERROR", "internal server error")
		return
	}
	writeJSON(w, 201, CreateRoomResponse{Room: mapRoom(room)})
}

func (h *RoomsHandler) List(w http.ResponseWriter, r *http.Request) {
	rooms, err := h.svc.List(r.Context())
	if err != nil {
		log.Printf("list rooms: %v", err)
		writeError(w, 500, "INTERNAL_ERROR", "internal server error")
		return
	}
	out := make([]RoomResponse, 0, len(rooms))
	for _, room := range rooms {
		out = append(out, mapRoom(room))
	}
	writeJSON(w, 200, RoomsListResponse{Rooms: out})
}
