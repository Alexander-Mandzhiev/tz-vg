package repository

import (
	"context"
	"errors"
	"tz-vg/internal/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrInvalidCredentials = errors.New("неправильный формат данных")
	ErrInternalServerErr  = errors.New("проблема на сервере")
	ErrTaskNotFound       = errors.New("задача не найдена")
)

type Todo interface {
	Create(ctx context.Context, note *entity.Todo) error
	GetAll(ctx context.Context) ([]entity.Todo, error)
	GetOne(ctx context.Context, id int) (entity.Todo, error)
	Update(ctx context.Context, id int, todo *entity.Todo) error
	Delete(ctx context.Context, id int) error
}

type Repository struct {
	pool *pgxpool.Pool
	Todo
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
		Todo: NewTodo(pool),
	}
}

func (p *Repository) Close() {
	p.pool.Close()
}
