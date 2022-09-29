package controller

import (
	"net/http"

	"github.com/geronimo794/golang-ach-rozikin-mastering/project1/model"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project1/service"
	"github.com/labstack/echo/v4"
)

type TodoControllerImpl struct {
	TodoService service.TodoService
}

func NewTodoController(todoService service.TodoService) TodoController {
	return &TodoControllerImpl{
		TodoService: todoService,
	}
}
func (controller *TodoControllerImpl) Create(e echo.Context) error {
	request_data := model.RequestTodo{}
	err := e.Bind(&request_data)
	if err != nil {
		return e.String(http.StatusBadRequest, err.Error())
	}
	controller.TodoService.Create(e.Request().Context(), request_data)

	return e.String(http.StatusOK, request_data.Name)
}
