package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/helper"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/model"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/model/web"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/repository"
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

func (service *TodoServiceImpl) Create(ctx context.Context, request web.RequestTodo) model.Todo {
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
func (service *TodoServiceImpl) FindAll(ctx context.Context, request web.RequestParameterTodo) []model.Todo {
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
func (service *TodoServiceImpl) Delete(ctx context.Context, id int) model.Todo {
	// Create transaction in this service
	tx := helper.StartTransaction(service.DB)
	defer helper.CommitOrRollback(tx)

	// Check if db todo exist
	db_todo := service.TodoRepository.FindById(ctx, tx, id)
	if db_todo == (model.Todo{}) {
		return db_todo
	}

	service.TodoRepository.Delete(ctx, service.DB, id)

	return db_todo
}
func (service *TodoServiceImpl) ReverseStatus(ctx context.Context, id int) model.Todo {
	// Create transaction in this service
	tx := helper.StartTransaction(service.DB)
	defer helper.CommitOrRollback(tx)

	// Check if db todo exist
	db_todo := service.TodoRepository.FindById(ctx, tx, id)
	if db_todo == (model.Todo{}) {
		return db_todo
	}

	// Reverse the status
	if db_todo.Status == 0 {
		db_todo.Status = 1
	} else {
		db_todo.Status = 0
	}
	db_todo = service.TodoRepository.Update(ctx, service.DB, db_todo)

	return db_todo
}
