package app

import (

	"github.com/geronimo794/golang-ach-rozikin-mastering/project1/controller"
	"github.com/labstack/echo/v4"
)

func SetRouter(e *echo.Echo, todoController controller.TodoController) {
	e.POST("/todo", todoController.Create)

}