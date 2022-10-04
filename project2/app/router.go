package app

import (
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/controller"
	"github.com/labstack/echo/v4"
)

func SetRouterAuth(e *echo.Echo, authController controller.AuthController) {
	e.POST("/login", authController.Login)

}

func SetRouterTodo(e *echo.Echo, todoController controller.TodoController) {
	e.POST("/todo", todoController.Create)
	e.GET("/todo", todoController.FindAll)
	e.GET("/todo/:id", todoController.FindById)
	e.PUT("/todo/:id", todoController.Update)
	e.DELETE("/todo/:id", todoController.Delete)
	e.PUT("/todo/:id/reverse-status", todoController.ReverseStatus)
}
