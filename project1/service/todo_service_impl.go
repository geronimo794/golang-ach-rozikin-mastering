package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	"github.com/geronimo794/golang-ach-rozikin-mastering/project1/helper"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project1/model"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project1/repository"
)

type TodoServiceImpl struct {
	TodoRepository repository.TodoRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewTodoService(todoRepository repository.TodoRepository, db *gorm.DB, validate *validator.Validate) TodoService {
	return &TodoServiceImpl{
		TodoRepository: todoRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *TodoServiceImpl) Create(ctx context.Context, request model.RequestTodo) model.Todo {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	helper.PanicIfError(tx.Error)
	defer helper.CommitOrRollback(tx)

	todo := model.Todo{
		Name: request.Name,
	}

	todo = service.TodoRepository.Create(ctx, service.DB, todo)

	return todo
}
