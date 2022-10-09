package controller

import (
	"github.com/labstack/echo/v4"
)

type TodoController interface {
	Create(echo.Context) error
	FindAll(echo.Context) error
	FindById(echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
	ReverseIsDone(echo.Context) error
}
