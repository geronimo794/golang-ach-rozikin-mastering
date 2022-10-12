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
		// Query no selected field
		{
			ReqQuery: `query{
				todos{
				}
			}`,
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusUnprocessableEntity,
			},
		},
		// Query wrong selected
		{
			ReqQuery: `query{
						todos{
							nama
						}
					}`,
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusUnprocessableEntity,
			},
		},
		// Query wrong input parameter selected
		{
			ReqQuery: `query{
							todos(input: {}){
								name
							}
						}`,
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusUnprocessableEntity,
			},
		},
	}

	// Doing test with table test
	for _, v := range testCase {
		apitest.New().
			Handler(graphQLHander).
			Post("/query").
			GraphQLQuery(v.ReqQuery).
			Expect(t).
			Status(v.Exp.ExpectedCode).
			End()
	}
}

// todos : Success
func TestTodos_Success(t *testing.T) {
	graphQLHander := setUpGraphHandler()

	var testCase = []TestCaseGraph{
		// Query select name
		{
			ReqQuery: `query{
				todos{
					name
				}
			}`,
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusOK,
			},
		},
		// Query select id, name
		{
			ReqQuery: `query{
				todos{
					id
					name
				}
			}`,
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusOK,
			},
		},
		// Query select id, name, priority
		{
			ReqQuery: `query{
				todos{
					id
					name
					priority
				}
			}`,
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusOK,
			},
		},
	}

	// Doing test with table test
	for _, v := range testCase {
		test := apitest.New().
			Handler(graphQLHander).
			Post("/query").
			GraphQLQuery(v.ReqQuery).
			Expect(t).
			Status(v.Exp.ExpectedCode)

		test.End()
	}
}

// todo : Failed
func TestTodo_Failed(t *testing.T) {
	graphQLHander := setUpGraphHandler()

	var testCase = []TestCaseGraph{
		// Query no selected field
		{
			ReqQuery: `query{
				todo{
				}
			}`,
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusUnprocessableEntity,
			},
		},
		// Query wrong selected
		{
			ReqQuery: `query{
						todo{
							namanya
						}
					}`,
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusUnprocessableEntity,
			},
		},
		// Query wrong input parameter selected
		{
			ReqQuery: `query{
							todo(id: xxx){
								name
							}
						}`,
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusUnprocessableEntity,
			},
		},
	}

	// Doing test with table test
	for _, v := range testCase {
		apitest.New().
			Handler(graphQLHander).
			Post("/query").
			GraphQLQuery(v.ReqQuery).
			Expect(t).
			Status(v.Exp.ExpectedCode).
			End()
	}
}

// todo : Success
func TestTodo_Success(t *testing.T) {
	graphQLHander := setUpGraphHandler()

	var testCase = []TestCaseGraph{
		// Query select name
		{
			ReqQuery: `query{
				todo(id: "1"){
					name
				}
			}`,
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusOK,
			},
		},
		// Query select id, name
		{
			ReqQuery: `query{
				todo(id: "1"){
					id
					name
				}
			}`,
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusOK,
			},
		},
		// Query select id, name, priority
		{
			ReqQuery: `query{
				todo(id: "1"){
					id
					name
					priority
				}
			}`,
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusOK,
			},
		},
	}

	// Doing test with table test
	for _, v := range testCase {
		test := apitest.New().
			Handler(graphQLHander).
			Post("/query").
			GraphQLQuery(v.ReqQuery).
			Expect(t).
			Status(v.Exp.ExpectedCode)

		test.End()
	}
}

/**
* Test Mutation
**/
// createTodo : Failed
func TestCreateTodo_Failed(t *testing.T) {
	graphQLHander := setUpGraphHandler()

	var testCase = []TestCaseGraph{
		// Query no selected field
		{
			ReqQuery: `mutation {
				createTodo(input: {name: "a", priority: high}) {
				}
			}`,
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusUnprocessableEntity,
			},
		},
		// Query wrong input
		{
			ReqQuery: `mutation {
				createTodo(input: {name: "a", priority: "high"}) {
					name
				}
			}`,
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusUnprocessableEntity,
			},
		},
		// Query wrong empty name
		{
			ReqQuery: `mutation {
				createTodo(input: {priority: "high"}) {
					name
				}
			}`,
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusUnprocessableEntity,
			},
		},
	}

	// Doing test with table test
	for _, v := range testCase {
		apitest.New().
			Handler(graphQLHander).
			Post("/query").
			GraphQLQuery(v.ReqQuery).
			Expect(t).
			Status(v.Exp.ExpectedCode).
			End()
	}
}

// createTodo : Success
func TestCreateTodo_Success(t *testing.T) {
	graphQLHander := setUpGraphHandler()

	var testCase = []TestCaseGraph{
		// Query create succcess
		{
			ReqQuery: `mutation {
				createTodo(input: {name: "Create todo", priority: low}) {
					name
					priority
				}
			}`,
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusOK,
			},
		},
		// Query create succcess
		{
			ReqQuery: `mutation {
					createTodo(input: {name: "Create todo oke", priority: high}) {
						id
						name
						priority
					}
				}`,
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusOK,
			},
		},
	}

	// Doing test with table test
	for _, v := range testCase {
		test := apitest.New().
			Handler(graphQLHander).
			Post("/query").
			GraphQLQuery(v.ReqQuery).
			Expect(t).
			Status(v.Exp.ExpectedCode)

		test.End()
	}
}

// updateTodo : Failed
// updateTodo : Success
// deleteTodo : Failed
// deleteTodo : Success
// reverseStatusTodo : Failed
// reverseStatusTodo : Success
