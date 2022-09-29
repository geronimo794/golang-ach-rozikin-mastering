package helper

import (
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
