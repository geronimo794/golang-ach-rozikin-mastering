package helper

import (
	"net/http"
	"strings"

	"github.com/geronimo794/golang-ach-rozikin-mastering/project1/model"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
func CommitOrRollback(tx *gorm.DB) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback().Error
		PanicIfError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit().Error
		PanicIfError(errorCommit)
	}
}
func BuildJsonResponse(e echo.Context, http_status int, data any, err []model.ErrorResponse) error {
	response := model.ResponseStandard{
		Code:    http_status,
		Message: http.StatusText(http_status),
		Data:    data,
		Errors:  err,
	}
	return e.JSON(http_status, response)
}
func CreateErrorResponse(title string, message ...string) model.ErrorResponse {
	errResp := model.ErrorResponse{
		Title:    title,
		Messages: message,
	}
	return errResp
}
func CreateValidationErrorResponse(validatorError error) (sliceErrorResponse []model.ErrorResponse) {
	var message string
	for _, err := range validatorError.(validator.ValidationErrors) {
		message = strings.Trim("is "+err.ActualTag()+" "+err.Param(), " ")
		sliceErrorResponse = append(sliceErrorResponse, CreateErrorResponse(err.Field(), message))
	}
	return sliceErrorResponse
}
func CreateEchoBindErrorResponse(echoBindError error) (sliceErrorResponse []model.ErrorResponse) {
	var message string
	for _, err := range echoBindError.(echo.Err) {
		message = strings.Trim("is "+err.ActualTag()+" "+err.Param(), " ")
		sliceErrorResponse = append(sliceErrorResponse, CreateErrorResponse(err.Field(), message))
	}
	return sliceErrorResponse
}
