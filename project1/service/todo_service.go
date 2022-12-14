package service

import (
	"context"

	"github.com/geronimo794/golang-ach-rozikin-mastering/project1/model"
)

type TodoService interface {
	Create(ctx context.Context, request model.RequestTodo) model.Todo
	FindAll(ctx context.Context, request model.RequestParameterTodo) []model.Todo
	FindById(ctx context.Context, id int) model.Todo
	Update(ctx context.Context, todo model.Todo) model.Todo
	Delete(ctx context.Context, id int) model.Todo
	ReverseStatus(ctx context.Context, id int) model.Todo
}
