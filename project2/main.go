package main

import (
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/app"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/controller"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/repository"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := app.NewDatabase()
	e := echo.New()
	validate := validator.New()

	// Auth API
	authService := service.NewAuthService()
	authController := controller.NewAuthController(authService, validate)
	app.SetRouterAuth(e, authController)

	// Todo API
	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository, db, validate)
	todoController := controller.NewTodoController(todoService, validate)
	app.SetRouterTodo(e, todoController)

	e.Use(middleware.Recover())
	e.Start(":3000")
}
