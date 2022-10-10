package controller

import (
	"github.com/labstack/echo/v4"
)

type GraphController interface {
	Query(echo.Context) error
	PlayGround(echo.Context) error
}
