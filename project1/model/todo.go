package model

type Todo struct {
	Id       int    `json:"id"`
	Name     string `json:"name" validate:"required"`
	Priority string `json:"priority"`
	Status   int    `json:"status"`
}
