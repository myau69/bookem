package postgres

import (
	"context"
	"database/sql"

	"github.com/myau69/bookem/internal/models"
)

type RoomsRepository struct {
	db *sql.DB
}

func NewRoomsRepository(db *sql.DB) *RoomsRepository {
	return &RoomsRepository{db: db}
}

type rowScanner interface {
	Scan(dest ...any) error
}

func scanRoom(row rowScanner) (models.Room, error) {
	var room models.Room
	var descriprion sql.NullString
	var capacity sql.NullInt64
	if err := row.Scan(&room.ID, &room.Name, &descriprion, &capacity, &room.CreatedAt); err != nil {
		return models.Room{}, err
	}
	if descriprion.Valid {
		value := descriprion.String
		room.Description = &value
	}
	if capacity.Valid {
		value := int(capacity.Int64)
		room.Capacity = &value
	}
	room.CreatedAt = room.CreatedAt.UTC()
	return room, nil
}

func (r *RoomsRepository) Create(ctx context.Context, room models.Room) (models.Room, error) {
	const query = `
	INSERT INTO rooms (id, name, description, capacity, created_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, name, description, capacity, created_at
	`
	return scanRoom(r.db.QueryRowContext(ctx, query, room.ID, room.Name, room.Description, room.Capacity, room.CreatedAt))
}

func (r *RoomsRepository) List(ctx context.Context) ([]models.Room, error) {
	const query = `
	SELECT id, name, description, capacity, created_at
	FROM rooms ORDER BY created_at DESC, id DESC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	rooms := make([]models.Room, 0)
	for rows.Next() {
		room, err := scanRoom(rows)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return rooms, nil
}
