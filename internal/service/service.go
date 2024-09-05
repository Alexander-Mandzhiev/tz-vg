package service

import (
	"context"
	"tz-vg/internal/entity"
	"tz-vg/internal/repository"
)

type Todo interface {
	Create(ctx context.Context, note *entity.Todo) error
	GetAll(ctx context.Context) ([]entity.Todo, error)
	GetOne(ctx context.Context, id int) (entity.Todo, error)
	Update(ctx context.Context, id int, todo *entity.Todo) error
	Delete(ctx context.Context, id int) error
}
type Service struct {
	Todo
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Todo: NewTodoService(repository.Todo),
	}
}
