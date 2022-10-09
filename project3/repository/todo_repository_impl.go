package repository

import (
	"context"
	"strconv"
	"strings"

	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/helper"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/model"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/model/web"
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
func (repository TodoRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB, param web.RequestParameterTodo) (todos []model.Todo) {

	if len(strings.Trim(param.Keyword, " ")) > 0 {
		tx = tx.Where("name LIKE ?", "%"+param.Keyword+"%")
	}

	findIsDone, err := strconv.ParseBool(param.IsDone)
	if err == nil {
		strFindIsDone := "0"
		if findIsDone {
			strFindIsDone = "1"
		}
		tx = tx.Where("is_done = ?", strFindIsDone)
	}
	err = tx.Find(&todos).Error
	helper.PanicIfError(err)
	return todos
}
func (repository TodoRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, id int) model.Todo {
	todo := model.Todo{}
	tx = tx.First(&todo, id)
	helper.PanicIfError(tx.Error)
	return todo
}
func (repository TodoRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, todo model.Todo) model.Todo {
	err := tx.Save(&todo).Error
	helper.PanicIfError(err)
	return todo
}
func (repository TodoRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, id int) {
	err := tx.Delete(&model.Todo{}, id).Error
	helper.PanicIfError(err)
}
