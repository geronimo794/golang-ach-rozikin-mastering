package helper

import (
	"errors"
	"net/http"
	"strings"

	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/model/web"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func PanicIfError(err error) {
	if err != nil {
		// If the error is error not found, no need to be panic
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			panic(err)
		}
	}
}
func StartTransaction(inputtx *gorm.DB) (tx *gorm.DB) {
	tx = inputtx.Begin()
	if tx.Error != nil {
		PanicIfError(tx.Error)
	}
	return tx
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
func BuildJsonResponse(e echo.Context, http_status int, data any, err []web.ErrorResponse) error {
	response := web.ResponseStandard{
		Code:    http_status,
		Message: http.StatusText(http_status),
		Data:    data,
		Errors:  err,
	}
	return e.JSON(http_status, response)
}
func CreateErrorResponse(title string, message ...string) web.ErrorResponse {
	errResp := web.ErrorResponse{
		Title:    title,
		Messages: message,
	}
	return errResp
}
func CreateValidationErrorResponse(validatorError error) (sliceErrorResponse []web.ErrorResponse) {
	var message string
	for _, err := range validatorError.(validator.ValidationErrors) {
		message = strings.Trim("is "+err.ActualTag()+" "+err.Param(), " ")
		sliceErrorResponse = append(sliceErrorResponse, CreateErrorResponse(strings.ToLower(err.Field()), message))
	}
	return sliceErrorResponse
}
