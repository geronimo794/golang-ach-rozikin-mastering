package service

import (
	"context"

	"github.com/geronimo794/golang-ach-rozikin-mastering/project1/model"
)

type TodoService interface {
	Create(ctx context.Context, request model.RequestTodo) (model.Todo, error)
	// FindAll(ctx context.Context, request model.RequestParameterTodo) model.Todo
	// Update(ctx context.Context, id int, request model.RequestTodo) model.Todo
	// Delete(ctx context.Context, id int, request model.RequestTodo)
	// ReverseStatus(ctx context.Context, id int) model.Todo
}
