package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/graph/generated"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.TodoInput) (*model.Todo, error) {
	// var todo = r.todoService.Create(*input)
	return &model.Todo{}, nil
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, id string, input model.TodoInput) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: UpdateTodo - updateTodo"))
}

// DeleteTodo is the resolver for the deleteTodo field.
func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: DeleteTodo - deleteTodo"))
}

// ReverseStatusTodo is the resolver for the reverseStatusTodo field.
func (r *mutationResolver) ReverseStatusTodo(ctx context.Context, id string) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: ReverseStatusTodo - reverseStatusTodo"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	// todos = r.todoService.FindAll()
	panic(fmt.Errorf("not implemented: ReverseStatusTodo - reverseStatusTodo"))
}

// Todo is the resolver for the todo field.
func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	// todo, err = r.todoService.FindById(ctx, id)
	panic(fmt.Errorf("not implemented: ReverseStatusTodo - reverseStatusTodo"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
