package main

import (
	"github.com/geronimo794/golang-ach-rozikin-mastering/project1/app"
	"github.com/labstack/echo/v4"
)

func main() {
	app.NewDatabase()

	e := echo.New()
	app.SetRouter(e)
	e.Logger.Fatal(e.Start(":3000"))
}
