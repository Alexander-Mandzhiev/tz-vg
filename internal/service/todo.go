package service

import (
	"context"
	"tz-vg/internal/entity"
	"tz-vg/internal/repository"
)

type TodoService struct {
	repo repository.Todo
}

func NewTodoService(repo repository.Todo) *TodoService {
	return &TodoService{repo: repo}
}

func (a *TodoService) Create(ctx context.Context, note *entity.Todo) error {
	const op = "service.Create"
	return a.repo.Create(ctx, note)
}
func (a *TodoService) GetAll(ctx context.Context) ([]entity.Todo, error) {
	const op = "service.GetAll"
	return a.repo.GetAll(ctx)
}

func (a *TodoService) GetOne(ctx context.Context, id int) (entity.Todo, error) {
	const op = "service.GetOne"
	return a.repo.GetOne(ctx, id)
}
func (a *TodoService) Update(ctx context.Context, id int, todo *entity.Todo) error {
	const op = "service.Update"
	return a.repo.Update(ctx, id, todo)
}
func (a *TodoService) Delete(ctx context.Context, id int) error {
	const op = "service.Delete"
	return a.repo.Delete(ctx, id)
}
