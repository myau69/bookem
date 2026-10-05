package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/myau69/bookem/internal/bootstrap"
	"github.com/myau69/bookem/internal/models"
	"github.com/myau69/bookem/internal/repository/postgres"
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
	capacity := 8
	room, err := models.NewRoom(uuid.New(), "Alpha", nil, &capacity, time.Now())
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	saved, err := repo.Create(ctx, room)
	if err != nil {
		return err
	}
	fmt.Println("Сохранена комната:", saved.ID, saved.Name)
	rooms, err := repo.List(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Всего комнат:", len(rooms))
	return nil
}
