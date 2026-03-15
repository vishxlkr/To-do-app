package repository

import (
	"context"
	"time"
	"todo/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTodo(pool *pgxpool.Pool, title string, completed bool) (*models.Todo, error) {
	// this creates a context with time 5 seconds , if the db takes more than 5 sec then it stop automatically.
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	var query string = `
		INSERT INTO todos_user (title , completed)
		VALUES ($1,$2)
		RETURNING id , title , completed, created_at , updated_at

	`

	var todo models.Todo

	// this query will only exist for 5 sec , given by ctx variable
	var err error = pool.QueryRow(ctx, query, title, completed).Scan(
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

func GetAllTodos(pool *pgxpool.Pool) ([]models.Todo, error) {
	var ctx context.Context
	var cancel context.CancelFunc

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	var query string = `
		SELECT id, title, completed, created_at, updated_at, user_id
		FROM todos
		
		ORDER BY created_at DESC
	`

	var rows , err =  pool.Query(ctx, query)

	if err!= nil {
		return nil, err
	}

	var todos []models.Todo  = []models.Todo{}

	for rows.Next(){
		var todo models.Todo

		// its takes the row values and put it into todo variable 
		var err = rows.scan(
			&todo.ID,
			&todo.Title,&todo.Completed,&todo.CreatedAt, &todo.UpdatedAt
		)

		if err!= nil {
			return nil , err
		}

		todos = append(todos, todo)
	}

	var err
}
