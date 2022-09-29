package model

type Todo struct {
	Id     int    `json:"id"`
	Name   string `json:"name" validate:"required"`
	Status int    `json:"status"`
}
