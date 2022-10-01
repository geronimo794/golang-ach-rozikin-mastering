package repository

import (
	"context"

	"github.com/geronimo794/golang-ach-rozikin-mastering/project1/helper"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project1/model"
	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
}

func NewTodoRepository() TodoRepository {
	return &TodoRepositoryImpl{}
}
func (repository TodoRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, todo model.Todo) model.Todo {
	err := tx.Create(&todo).Error
	helper.PanicIfError(err)
	return todo
}
