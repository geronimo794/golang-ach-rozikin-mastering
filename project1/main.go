package main

import (
	"github.com/geronimo794/golang-ach-rozikin-mastering/project1/app"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project1/controller"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project1/repository"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project1/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	db := app.NewDatabase()
	e := echo.New()
	validate := validator.New()
	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository, db, validate)
	todoController := controller.NewTodoController(todoService, validate)
	app.SetRouter(e, todoController)
	e.Start(":3000")
}
