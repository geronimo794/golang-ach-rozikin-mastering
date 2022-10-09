package controller

import (
	"github.com/labstack/echo/v4"
)

type AuthController interface {
	Authenticate(echo.Context) error
}
