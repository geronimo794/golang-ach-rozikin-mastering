package helper

import (
	"github.com/geronimo794/golang-ach-rozikin-mastering/project1/model"
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
func BuildResponse(data any, err []any) model.ResponseStandard {
	response := model.ResponseStandard{
		Data:  data,
		Error: err,
	}
	return response
}
