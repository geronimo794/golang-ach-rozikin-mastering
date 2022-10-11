package test

import (
	"net/http"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/graph"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/graph/generated"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/repository"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/service"
	"github.com/go-playground/validator/v10"
	"github.com/steinfletcher/apitest"
)

// Setup graphQL Handler
func setUpGraphHandler() *http.ServeMux {
	// Create depedency
	validate := validator.New()
	db := NewDatabaseTest()
	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository, db, validate)

	// Data sample
	ClearTodosData(db)
	CreateSampleTodoData(db, todoRepository)

	// Create handler for graphQL
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{TodoService: todoService}}))

	mux := http.NewServeMux()
	mux.Handle("/query", srv)

	return mux
}

// Struct test case
type TestCaseGraph struct {
	ReqQuery string
	Exp      ExpectationResultTest
}

/**
* Test Query
**/
// todos : Failed
func TestTodos_Failed(t *testing.T) {
	graphQLHander := setUpGraphHandler()

	var testCase = []TestCaseGraph{
		{
			ReqQuery: "",
		},
	}

	// Doing test with table test
	for _, v := range testCase {
		test := apitest.New().
			Handler(graphQLHander).
			Post("/query").
			GraphQLQuery(v.ReqQuery).
			Expect(t)

		test.Status(http.StatusOK).
			Body(v.Exp.ExpectedData)

		test.End()
	}
}

// todos : Success
// todo : Failed
// todo : Success

/**
* Test Mutation
**/
// createTodo : Failed
// createTodo : Success
// updateTodo : Failed
// updateTodo : Success
// deleteTodo : Failed
// deleteTodo : Success
// reverseStatusTodo : Failed
// reverseStatusTodo : Success
