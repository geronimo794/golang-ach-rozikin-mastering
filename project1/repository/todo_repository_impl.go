package repository

import (
	"context"
	"strings"

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

func (repository TodoRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB, param model.RequestParameterTodo) (todos []model.Todo) {

	if len(strings.Trim(param.Keyword, " ")) > 0 {
		tx = tx.Where("name LIKE ?", "%"+param.Keyword+"%")
	}

	if param.Status != -1 {
		tx = tx.Where("status = ?", param.Status)
	}
	err := tx.Find(&todos).Error
	helper.PanicIfError(err)
	return todos
}
