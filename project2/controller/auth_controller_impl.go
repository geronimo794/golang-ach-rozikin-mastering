package controller

import (
	"net/http"

	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/helper"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/model/web"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
	Validate    *validator.Validate
}

func NewAuthController(authService service.AuthService, validate *validator.Validate) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
		Validate:    validate,
	}
}

func (controller *AuthControllerImpl) Authenticate(e echo.Context) error {
	// Gather the form data
	request_data := web.RequestAuth{}
	e.Bind(&request_data)

	// Validation form
	err := controller.Validate.Struct(request_data)

	// If form data failed to validate
	if err != nil {
		errs := helper.CreateValidationErrorResponse(err)
		return helper.BuildJsonResponse(e, http.StatusBadRequest, nil, errs)
	}

	// Check the credential user with login service. If failed response false
	if !controller.AuthService.Login(e.Request().Context(), request_data) {
		return helper.BuildJsonResponse(e, http.StatusUnauthorized, nil, nil)
	}

	// Create repsponse
	response_data := controller.AuthService.GenerateToken(e.Request().Context(), request_data)

	return helper.BuildJsonResponse(e, http.StatusCreated, response_data, nil)

}
