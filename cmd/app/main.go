package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/myau69/bookem/internal/bootstrap"
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	capacity := 8
	room, err := svc.Create(ctx, "Alpha", nil, &capacity)
	if err != nil {
		return err
	}
	fmt.Println("Сохранена комната:", room.ID, room.Name)
	rooms, err := svc.List(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Всего комнат:", len(rooms))
	return nil
}
