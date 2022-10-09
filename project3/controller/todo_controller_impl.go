package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/helper"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/model"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/model/web"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/service"
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
	request_data := web.RequestTodo{}
	request_data.Name = e.FormValue("name")
	request_data.Priority = e.FormValue("priority")

	// Validate the form data
	err := controller.Validate.Struct(request_data)

	// // If form data failed to validate
	if err != nil {
		errs := helper.CreateValidationErrorResponse(err)
		return helper.BuildJsonResponse(e, http.StatusBadRequest, nil, errs)
	}

	// Create repsponse
	response_data := controller.TodoService.Create(e.Request().Context(), request_data)

	return helper.BuildJsonResponse(e, http.StatusCreated, response_data, nil)
}
func (controller *TodoControllerImpl) FindAll(e echo.Context) error {
	// Gather the form data
	request_data := web.RequestParameterTodo{
		Status: -1,
	}
	echo.QueryParamsBinder(e).
		String("keyword", &request_data.Keyword).
		Int8("status", &request_data.Status)

	response_data := controller.TodoService.FindAll(e.Request().Context(), request_data)
	if len(response_data) == 0 {
		return helper.BuildJsonResponse(e, http.StatusNotFound, nil, nil)
	}

	return helper.BuildJsonResponse(e, http.StatusOK, response_data, nil)
}
func (controller *TodoControllerImpl) FindById(e echo.Context) error {
	// Gather request
	param_id := e.Param("id")
	id, _ := strconv.Atoi(param_id)

	response_data := controller.TodoService.FindById(e.Request().Context(), id)
	if (model.Todo{}) == response_data {
		return helper.BuildJsonResponse(e, http.StatusNotFound, nil, nil)
	}
	return helper.BuildJsonResponse(e, http.StatusOK, response_data, nil)
}

func (controller *TodoControllerImpl) Update(e echo.Context) error {
	// Get header parameter data
	param_id := e.Param("id")
	id, _ := strconv.Atoi(param_id)

	// Get form data request body
	request_data := model.Todo{}
	request_data.Name = e.FormValue("name")
	request_data.Priority = e.FormValue("priority")

	// Validate the form data
	err := controller.Validate.Struct(request_data)

	// // If form data failed to validate
	if err != nil {
		errs := helper.CreateValidationErrorResponse(err)
		return helper.BuildJsonResponse(e, http.StatusBadRequest, nil, errs)
	}
	request_data.Id = id
	fmt.Println(request_data)
	response_data := controller.TodoService.Update(e.Request().Context(), request_data)

	// If empty data then not found
	if (model.Todo{}) == response_data {
		return helper.BuildJsonResponse(e, http.StatusNotFound, nil, nil)
	}

	return helper.BuildJsonResponse(e, http.StatusOK, response_data, nil)
}
func (controller *TodoControllerImpl) Delete(e echo.Context) error {
	// Gather request
	param_id := e.Param("id")
	id, _ := strconv.Atoi(param_id)

	response_data := controller.TodoService.Delete(e.Request().Context(), id)
	if (model.Todo{}) == response_data {
		return helper.BuildJsonResponse(e, http.StatusNotFound, nil, nil)
	}
	return helper.BuildJsonResponse(e, http.StatusOK, response_data, nil)
}
func (controller *TodoControllerImpl) ReverseIsDone(e echo.Context) error {
	// Gather request
	param_id := e.Param("id")
	id, _ := strconv.Atoi(param_id)

	response_data := controller.TodoService.ReverseIsDone(e.Request().Context(), id)
	if (model.Todo{}) == response_data {
		return helper.BuildJsonResponse(e, http.StatusNotFound, nil, nil)
	}
	return helper.BuildJsonResponse(e, http.StatusOK, response_data, nil)
}
