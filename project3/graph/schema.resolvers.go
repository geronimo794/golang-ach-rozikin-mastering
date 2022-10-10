package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/graph/generated"
	gModel "github.com/geronimo794/golang-ach-rozikin-mastering/project3/graph/model"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/helper"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/model"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/model/web"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input gModel.TodoInput) (*gModel.Todo, error) {
	requestData := web.RequestTodo{
		Name:     input.Name,
		Priority: string(input.Priority),
	}
	newData := r.TodoService.Create(ctx, requestData)
	// var todo = r.todoService.Create(*input)
	return helper.ConvertTodoToGraphTodo(&newData), nil
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, id string, input gModel.TodoInput) (*gModel.Todo, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	editData := r.TodoService.Update(ctx, model.Todo{
		Id:       intId,
		Name:     input.Name,
		Priority: string(input.Priority),
	})
	// var todo = r.todoService.Create(*input)
	return helper.ConvertTodoToGraphTodo(&editData), nil
}

// DeleteTodo is the resolver for the deleteTodo field.
func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (*gModel.Todo, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	deleteData := r.TodoService.Delete(ctx, intId)

	return helper.ConvertTodoToGraphTodo(&deleteData), nil
}

// ReverseStatusTodo is the resolver for the reverseStatusTodo field.
func (r *mutationResolver) ReverseStatusTodo(ctx context.Context, id string) (*gModel.Todo, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	deleteData := r.TodoService.ReverseIsDone(ctx, intId)

	return helper.ConvertTodoToGraphTodo(&deleteData), nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*gModel.Todo, error) {
	todo := r.TodoService.FindAll(ctx, web.RequestParameterTodo{})

	return helper.ConvertListTodoToGraphTodo(&todo), nil
}

// Todo is the resolver for the todo field.
func (r *queryResolver) Todo(ctx context.Context, id string) (*gModel.Todo, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	todoData := r.TodoService.FindById(ctx, intId)

	return helper.ConvertTodoToGraphTodo(&todoData), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
