package app

import (
	"github.com/geronimo794/golang-ach-rozikin-mastering/project1/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	dsn := "root:menjadilebihbaik@tcp(127.0.0.1:3306)/todo_project_golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	// Migrate database model to database
	migrateTable(db)


	tx := db.Session(&gorm.Session{SkipDefaultTransaction: true})
	return tx
}
func migrateTable(db *gorm.DB) {
	db.AutoMigrate(&model.Todo{})
}