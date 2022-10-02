package repository

import (
	"context"

	"github.com/geronimo794/golang-ach-rozikin-mastering/project1/model"
	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(ctx context.Context, tx *gorm.DB, todo model.Todo) model.Todo
	FindAll(ctx context.Context, tx *gorm.DB, param model.RequestParameterTodo) []model.Todo
	FindById(ctx context.Context, tx *gorm.DB, id int) model.Todo
	Update(ctx context.Context, tx *gorm.DB, todo model.Todo) model.Todo
	Delete(ctx context.Context, tx *gorm.DB, id int)
}
