package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handlers struct {
	Rooms *RoomsHandler
}

func NewRouter(h Handlers) http.Handler {
	r := chi.NewRouter()
	r.HandleFunc("/_info", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, 200, map[string]string{"status": "ok"})
	})
	r.Post("/rooms/create", h.Rooms.Create)
	r.Get("/rooms/list", h.Rooms.List)
	return r
}
