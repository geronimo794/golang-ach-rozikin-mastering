package model

type RequestParameterTodo struct {
	Keyword string `json:"keyword"`
	Status  int    `json:"status" validate:"oneof=0 1"`
}
