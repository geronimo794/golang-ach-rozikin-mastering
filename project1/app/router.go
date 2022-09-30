package app

import (
	"github.com/geronimo794/golang-ach-rozikin-mastering/project1/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetRouter(e *echo.Echo, todoController controller.TodoController) {
	e.POST("/todo", todoController.Create)

	e.Use(middleware.Recover())

}
