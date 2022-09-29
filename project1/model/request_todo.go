package model

type RequestTodo struct {
	Name   string `json:"name" validate:"required"`
	Status int    `json:"status"`
}
