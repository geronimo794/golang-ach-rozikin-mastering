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

func (service *TodoServiceImpl) Create(ctx context.Context, request model.RequestTodo) (model.Todo, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return model.Todo{}, err
	}

	tx := service.DB.Begin()
	if tx.Error != nil {
		return model.Todo{}, err
	}

	defer helper.CommitOrRollback(tx)

	todo := model.Todo{
		Name:     request.Name,
		Priority: request.Priority,
	}

	todo = service.TodoRepository.Create(ctx, service.DB, todo)

	return todo, nil

}
