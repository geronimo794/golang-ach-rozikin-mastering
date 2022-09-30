package controller

import (
	"net/http"

	"github.com/geronimo794/golang-ach-rozikin-mastering/project1/helper"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project1/model"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project1/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type TodoControllerImpl struct {
	TodoService service.TodoService
	Validate    *validator.Validate
}

func NewTodoController(todoService service.TodoService, validate *validator.Validate) TodoController {
	return &TodoControllerImpl{
		TodoService: todoService,
		Validate:    validate,
	}
}
func (controller *TodoControllerImpl) Create(e echo.Context) error {
	// Gather the form data
	request_data := model.RequestTodo{}
	request_data.Name = e.FormValue("name")
	request_data.Priority = e.FormValue("priority")

	// Validate the form data
	err := controller.Validate.Struct(request_data)

	// If form data failed to validate
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.BuildResponse(nil, []any{err.Error()}))
	}

	// Create repsponse
	response_data := model.Todo{}
	response_data, err = controller.TodoService.Create(e.Request().Context(), request_data)

	// If failed to save the data
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.BuildResponse(nil, []any{err.Error()}))
	}
	return e.JSON(http.StatusOK, helper.BuildResponse(response_data, nil))
}
