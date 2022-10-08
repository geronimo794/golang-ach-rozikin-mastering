package app

import (
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/config"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/controller"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/model/claim"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetAuthJWTGroup(e *echo.Echo) *echo.Group {
	jwtConfig := middleware.JWTConfig{
		Claims:     &claim.JwtCustomClaims{},
		SigningKey: []byte(config.JWTKEY),
	}
	eJWT := e.Group("")
	eJWT.Use(middleware.JWTWithConfig(jwtConfig))
	return eJWT
}

func SetRouterAuth(e *echo.Echo, authController controller.AuthController) {
	e.POST("/login", authController.Authenticate)
}

func SetRouterTodo(e *echo.Group, todoController controller.TodoController) {
	e.POST("/todo", todoController.Create)
	e.GET("/todo", todoController.FindAll)
	e.GET("/todo/:id", todoController.FindById)
	e.PUT("/todo/:id", todoController.Update)
	e.DELETE("/todo/:id", todoController.Delete)
	e.PUT("/todo/:id/reverse-status", todoController.ReverseStatus)
}
