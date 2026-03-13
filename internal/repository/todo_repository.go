package repository

import (
	"context"
	"time"
	"todo/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTodo(pool *pgxpool.Pool, title string, complted bool) (*models.Todo, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	var query string = `
		INSERT INTO todos_user (title , completed)
		VALUES ($1,$2)
		RETURNING id , title , completed, created_at , updated_at

	`

	var todo model.Todo

	var err error = pool.QueryRow(ctx, query, title, complted).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &todo, nil
}
