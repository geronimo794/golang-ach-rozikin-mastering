package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type GraphControllerImpl struct {
	PlaygroundHandler http.Handler
	QueryHandler      http.Handler
}

func NewGraphController(p http.Handler, q http.Handler) GraphController {
	return &GraphControllerImpl{
		PlaygroundHandler: p,
		QueryHandler:      q,
	}
}
func (controller *GraphControllerImpl) Query(e echo.Context) error {
	controller.QueryHandler.ServeHTTP(e.Response(), e.Request())

	return nil
}

func (controller *GraphControllerImpl) PlayGround(e echo.Context) error {

	controller.PlaygroundHandler.ServeHTTP(e.Response(), e.Request())
	return nil
}
