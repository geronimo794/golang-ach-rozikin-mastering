package test

import (
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/helper"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/model"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseTest() *gorm.DB {
	dsn := "root:menjadilebihbaik@tcp(127.0.0.1:3306)/todo_project_golang_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	// Migrate database model to database
	migrateTable(db)

	return db
}

func migrateTable(db *gorm.DB) {
	db.AutoMigrate(&model.Todo{})
}

/**
* To do data set
**/
func CreateSampleTodoData(db *gorm.DB, todoRepository repository.TodoRepository) {
	todoData := model.Todo{
		Name:     "Sample data low",
		Priority: "low",
		IsDone:   false,
	}
	todoRepository.Create(db.Statement.Context, db, todoData)

	todoData = model.Todo{
		Name:     "Sample data medium",
		Priority: "medium",
		IsDone:   false,
	}
	todoRepository.Create(db.Statement.Context, db, todoData)

	todoData = model.Todo{
		Name:     "Sample data high",
		Priority: "high",
		IsDone:   true,
	}
	todoRepository.Create(db.Statement.Context, db, todoData)

}
func ClearTodosData(db *gorm.DB) {
	// Truncate table
	err := db.Exec("TRUNCATE TABLE todos").Error
	helper.PanicIfError(err)
}

type ExpectationResultTest struct {
	ExpectedCode           int
	ExpectedContainData    string
	NotExpectedContainData string
}
