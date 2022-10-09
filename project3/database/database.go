package database

import (
	"fmt"
	"log"

	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/graph/model"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	var (
		DB_USER     string = utils.GetValue("DB_USER")
		DB_PASSWORD string = utils.GetValue("DB_PASSWORD")
		DB_HOST     string = utils.GetValue("DB_HOST")
		DB_PORT     string = utils.GetValue("DB_PORT")
		DB_NAME     string = utils.GetValue("DB_NAME")
	)
	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		DB_USER,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("Success connected to the database")

	DB.AutoMigrate(&model.Todo{})
}
