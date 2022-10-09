package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/graph/generated"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input *model.TodoInput) (*model.Todo, error) {
	var todo = r.todoService.Create(*input)
	return &todo, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) (todos []*model.Todo, err error) {
	todos = r.todoService.FindAll()
	return todos, err
}

// Todo is the resolver for the todo field.
func (r *queryResolver) Todo(ctx context.Context, id string) (todo *model.Todo, err error) {
	todo, err = r.todoService.FindById(id)
	return todo, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
