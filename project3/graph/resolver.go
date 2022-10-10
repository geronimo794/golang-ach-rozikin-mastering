package graph

import "github.com/geronimo794/golang-ach-rozikin-mastering/project3/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{ TodoService service.TodoService }
