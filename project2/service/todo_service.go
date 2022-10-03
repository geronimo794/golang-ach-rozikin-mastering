package service

import (
	"context"

	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/model"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/model/web"
)

type TodoService interface {
	Create(ctx context.Context, request web.RequestTodo) model.Todo
	FindAll(ctx context.Context, request web.RequestParameterTodo) []model.Todo
	FindById(ctx context.Context, id int) model.Todo
	Update(ctx context.Context, todo model.Todo) model.Todo
	Delete(ctx context.Context, id int) model.Todo
	ReverseStatus(ctx context.Context, id int) model.Todo
}
