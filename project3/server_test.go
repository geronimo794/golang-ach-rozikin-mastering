package main

import (
	"net/http"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/database"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/graph"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/graph/generated"
	"github.com/steinfletcher/apitest"
)

func graphQLHandler() http.Handler {
	database.Connect()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	mux := http.NewServeMux()

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	return mux
}

func TestGetTodos_Success(t *testing.T) {
	var query string = `{
		todos{
			id
			name
		}
	}`

	apitest.New().
		Handler(graphQLHandler()).
		Post("/query").
		GraphQLQuery(query).
		Expect(t).
		Status(http.StatusOK).
		End()
}
func TestCreateProduct_Success(t *testing.T) {

	var query string = `
	mutation{
		createTodo(input:{
		  name: "Mencoba graphql",
		  priority: medium
		}){
		  name
		  priority
		}
	}`

	var result string = `{"data": {"createTodo":{"name":"Mencoba graphql","priority":"medium"}}}`
	apitest.New().
		Handler(graphQLHandler()).
		Post("/query").
		GraphQLQuery(query).
		Expect(t).
		Status(http.StatusOK).
		Body(result).
		End()

}
func TestCreateProduct_Failed(t *testing.T) {

	var query string = `
	mutation{
		createTodo(input:{
		  name: ""
		}){
		  name
		  priority
		}
	}`

	apitest.New().
		Handler(graphQLHandler()).
		Post("/query").
		GraphQLQuery(query).
		Expect(t).
		Status(http.StatusUnprocessableEntity).
		End()

}
