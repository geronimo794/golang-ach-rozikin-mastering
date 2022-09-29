package model

type RequestTodo struct {
	Name   string `json:"name" form:"name" validate:"required"`
	Status int    `json:"status" form:"status"`
}
