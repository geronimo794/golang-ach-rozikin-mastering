package app

import (
	"fmt"

	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/model"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	DB_USER     string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

func NewDatabase() *gorm.DB {
	dbConfig := getDBConfigFromEnv()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		dbConfig.DB_USER,
		dbConfig.DB_PASSWORD,
		dbConfig.DB_HOST,
		dbConfig.DB_PORT,
		dbConfig.DB_NAME,
	)

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
func getDBConfigFromEnv() DBConfig {
	return DBConfig{
		DB_USER:     utils.GetValue("DB_USER"),
		DB_PASSWORD: utils.GetValue("DB_PASSWORD"),
		DB_HOST:     utils.GetValue("DB_HOST"),
		DB_PORT:     utils.GetValue("DB_PORT"),
		DB_NAME:     utils.GetValue("DB_NAME"),
	}

}
