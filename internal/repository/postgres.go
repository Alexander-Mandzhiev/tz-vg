package repository

import (
	"context"
	"errors"
	"fmt"
	"time"
	"tz-vg/internal/entity"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryTodo struct {
	pool *pgxpool.Pool
}

func NewTodo(pool *pgxpool.Pool) *RepositoryTodo {
	return &RepositoryTodo{pool: pool}
}

func (r *RepositoryTodo) Create(ctx context.Context, todo *entity.Todo) error {
	const op = "repository.Create"
	var pgErr *pgconn.PgError
	query := "INSERT INTO tasks ( title, description, due_date, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING *"

	if err := r.pool.QueryRow(ctx, query, todo.Title, todo.Description, todo.DueDate, time.Now(), time.Now()).
		Scan(&todo.ID, &todo.Title, &todo.Description, &todo.DueDate, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.InvalidParameterValue {
			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		return fmt.Errorf("%s: %w", op, ErrInternalServerErr)
	}
	return nil
}

func (r *RepositoryTodo) GetAll(ctx context.Context) ([]entity.Todo, error) {
	const op = "repository.GetAll"
	query := "SELECT * FROM tasks"
	var todos []entity.Todo
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return todos, fmt.Errorf("%s: %w", op, ErrInternalServerErr)
	}
	defer rows.Close()

	for rows.Next() {
		var todo entity.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.DueDate, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return todos, fmt.Errorf("%s: %w", op, ErrInternalServerErr)
		}
		todos = append(todos, todo)
	}
	if err := rows.Err(); err != nil {
		return todos, fmt.Errorf("%s: %w", op, ErrInternalServerErr)
	}
	return todos, nil
}

func (r *RepositoryTodo) GetOne(ctx context.Context, id int) (entity.Todo, error) {
	const op = "repository.GetOne"
	query := "SELECT * FROM tasks WHERE id = $1"
	var todo entity.Todo

	if err := r.pool.QueryRow(ctx, query, id).
		Scan(&todo.ID, &todo.Title, &todo.Description, &todo.DueDate, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
		if err == pgx.ErrNoRows {
			return todo, fmt.Errorf("%s: %w", op, ErrTaskNotFound)
		}
		return todo, fmt.Errorf("%s: %w", op, ErrInternalServerErr)
	}
	return todo, nil
}

func (r *RepositoryTodo) Update(ctx context.Context, id int, todo *entity.Todo) error {
	const op = "repository.Update"
	query := "UPDATE tasks SET title = $2, description = $3, due_date = $4, updated_at = $5 WHERE id = $1 RETURNING *"

	if err := r.pool.QueryRow(ctx, query, id, todo.Title, todo.Description, todo.DueDate, time.Now()).
		Scan(&todo.ID, &todo.Title, &todo.Description, &todo.DueDate, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("%s: %w", op, ErrTaskNotFound)
		}
		return fmt.Errorf("%s: %w", op, ErrInternalServerErr)
	}
	return nil
}
func (r *RepositoryTodo) Delete(ctx context.Context, id int) error {
	const op = "repository.Delete"
	query := "DELETE FROM tasks WHERE id = $1 RETURNING id"
	if err := r.pool.QueryRow(ctx, query, id).Scan(&id); err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("%s: %w", op, ErrTaskNotFound)
		}
		return fmt.Errorf("%s: %w", op, ErrInternalServerErr)
	}
	return nil
}
