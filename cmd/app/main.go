package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/myau69/bookem/internal/bootstrap"
	"github.com/myau69/bookem/internal/controllers"
	"github.com/myau69/bookem/internal/repository/postgres"
	"github.com/myau69/bookem/internal/service"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	databaseURL := os.Getenv("DATABASE_URL")
	if strings.TrimSpace(databaseURL) == "" {
		return errors.New("DATABASE_URL is required; load .env before starting")
	}
	if err := bootstrap.RunMigrations(databaseURL, "migrations"); err != nil {
		return err
	}
	db, err := bootstrap.OpenDatabase(databaseURL)
	if err != nil {
		return err
	}
	defer func() {
		_ = db.Close()
	}()
	repo := postgres.NewRoomsRepository(db)
	svc := service.NewRoomService(repo, time.Now)
	handler := controllers.NewRoomsHandler(svc)
	router := controllers.NewRouter(controllers.Handlers{Rooms: handler})
	server := &http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Println("HTTP server listening on :8080")
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
