package repository

import (
	"context"

	"github.com/geronimo794/golang-ach-rozikin-mastering/project1/model"
	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(ctx context.Context, tx *gorm.DB, todo model.Todo) model.Todo
	FindAll(ctx context.Context, tx *gorm.DB, param model.RequestParameterTodo) []model.Todo
}
