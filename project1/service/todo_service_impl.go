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
	// Create transaction in this service
	tx := helper.StartTransaction(service.DB)
	defer helper.CommitOrRollback(tx)

	todo := model.Todo{
		Name:     request.Name,
		Priority: request.Priority,
	}

	todo = service.TodoRepository.Create(ctx, service.DB, todo)

	return todo
}
func (service *TodoServiceImpl) FindAll(ctx context.Context, request model.RequestParameterTodo) []model.Todo {
	return service.TodoRepository.FindAll(ctx, service.DB, request)
}
func (service *TodoServiceImpl) FindById(ctx context.Context, id int) model.Todo {
	return service.TodoRepository.FindById(ctx, service.DB, id)
}

func (service *TodoServiceImpl) Update(ctx context.Context, todo model.Todo) model.Todo {
	// Create transaction in this service
	tx := helper.StartTransaction(service.DB)
	defer helper.CommitOrRollback(tx)

	// Check if db todo exist
	db_todo := service.TodoRepository.FindById(ctx, tx, todo.Id)
	if db_todo == (model.Todo{}) {
		return db_todo
	}

	todo = service.TodoRepository.Update(ctx, service.DB, todo)

	return todo
}
