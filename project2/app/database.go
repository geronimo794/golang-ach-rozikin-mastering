package app

import (
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newDatabaseProcess(env string) *gorm.DB {
	var dsn string
	switch env {
	case "production":
		dsn = "root:menjadilebihbaik@tcp(127.0.0.1:3306)/todo_project_golang?charset=utf8mb4&parseTime=True&loc=Local"
	case "test":
		dsn = "root:menjadilebihbaik@tcp(127.0.0.1:3306)/todo_project_golang_test?charset=utf8mb4&parseTime=True&loc=Local"
	default:
		dsn = "root:menjadilebihbaik@tcp(127.0.0.1:3306)/todo_project_golang_test?charset=utf8mb4&parseTime=True&loc=Local"
	}

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
func NewDatabaseProduction() *gorm.DB {
	return newDatabaseProcess("production")
}
func NewDatabaseTest() *gorm.DB {
	return newDatabaseProcess("test")
}
func migrateTable(db *gorm.DB) {
	db.AutoMigrate(&model.Todo{})
}
