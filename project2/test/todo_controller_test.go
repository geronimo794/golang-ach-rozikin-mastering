package test

import (
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/controller"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/model"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/repository"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/service"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func setUpTodoTestRouterController() controller.TodoController {
	validate := validator.New()
	db := NewDatabaseTest()

	// Auth API
	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository, db, validate)
	todoController := controller.NewTodoController(todoService, validate)
	return todoController
}
func setUpTodoDatabaseTest(db *gorm.DB, todoRepository repository.TodoRepository) {
	// Truncate table
	db.Exec("TRUNCATE TABLE todos")

	// Create single todo data
	todoData := model.Todo{
		Name:     "Example content todo data",
		Priority: "high",
	}
	todoRepository.Create(db.Statement.Context, db, todoData)
}
