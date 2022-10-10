package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/app"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/controller"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/graph"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/graph/generated"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/repository"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := app.NewDatabase()
	e := echo.New()
	validate := validator.New()

	// Auth API
	authService := service.NewAuthService()
	authController := controller.NewAuthController(authService, validate)
	app.SetRouterAuth(e, authController)

	// Set autentification group
	eGroup := app.SetAuthJWTGroup(e)

	// Todo API
	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository, db, validate)
	todoController := controller.NewTodoController(todoService, validate)
	app.SetRouterTodo(eGroup, todoController)

	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{TodoService: todoService}}))

	e.Use(middleware.Recover())
	e.Start(":3000")
}

// const defaultPort = "8080"

// func main() {
// 	database.Connect()

// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = defaultPort
// 	}

// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

// 	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
// 	http.Handle("/query", srv)

// 	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
// 	log.Fatal(http.ListenAndServe(":"+port, nil))
// }
